package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "frontend/register.html")
		return
	}
	// grab user info
	username := r.FormValue("username")
	password := r.FormValue("password")
	role := r.FormValue("role")
	// Check existence of user
	var user User
	err := db.QueryRow("SELECT username, password, role FROM users WHERE username=?",
		username).Scan(&user.Username, &user.Password, &user.Role)
	switch {
	// user is available
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		checkInternalServerError(err, w)
		// insert to database
		_, err = db.Exec(`INSERT INTO users(username, password, role) VALUES(?, ?, ?)`,
			username, hashedPassword, role)
		fmt.Println("Created user: ", username)
		checkInternalServerError(err, w)
	case err != nil:
		http.Error(w, "loi: "+err.Error(), http.StatusBadRequest)
		return
	default:
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "frontend/login.html")
		return
	}
	// grab user info from the submitted form
	username := r.FormValue("usrname")
	password := r.FormValue("psw")
	// query database to get match username
	var user User
	err := db.QueryRow("SELECT username, password FROM users WHERE username=?",
		username).Scan(&user.Username, &user.Password)
	checkInternalServerError(err, w)
	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/login", 301)
	}
	authenticated = true
	http.Redirect(w, r, "/list", 301)

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	authenticated = false
	isAuthenticated(w, r)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
	rows, err := db.Query("SELECT * FROM persediaan")
	checkInternalServerError(err, w)
	var funcMap = template.FuncMap{
		"multiplication": func(n float64, f float64) float64 {
			return n * f
		},
		"addOne": func(n int) int {
			return n + 1
		},
	}
	var persediaans []Persediaan
	var persediaan Persediaan
	for rows.Next() {
		err = rows.Scan(&persediaan.Id, &persediaan.Nopol, &persediaan.TglMasuk,
			&persediaan.Merk, &persediaan.Tahun, &persediaan.HargaModal, &persediaan.HargaJual, &persediaan.CreateDate, &persediaan.Suplier, &persediaan.Telp, &persediaan.Alamat, &persediaan.Deskripsi)
		checkInternalServerError(err, w)
		persediaans = append(persediaans, persediaan)
	}
	t, err := template.New("list.html").Funcs(funcMap).ParseFiles("frontend/list.html")
	checkInternalServerError(err, w)
	err = t.Execute(w, persediaans)
	checkInternalServerError(err, w)

}

func addPersediaanHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	if r.Method != "POST" {
		http.ServeFile(w, r, "frontend/addpersediaan.html")
		return
	}
}

func createPersediaanHandler(w http.ResponseWriter, r *http.Request) {
	//isAuthenticated(w, r)
	if r.Method != "POST" {
		//http.Redirect(w, r, "/", 301)
		http.ServeFile(w, r, "frontend/createpersediaan.html")
		return
	}
	var persediaan Persediaan
	persediaan.Nopol = r.FormValue("Nopol")
	persediaan.TglMasuk = r.FormValue("TglMasuk")
	persediaan.Merk = r.FormValue("Merk")
	persediaan.Tahun = r.FormValue("Tahun")
	persediaan.HargaModal, _ = strconv.ParseFloat(r.FormValue("HargaModal"), 64)
	persediaan.HargaJual, _ = strconv.ParseFloat(r.FormValue("HargaJual"), 64)
	persediaan.CreateDate = r.FormValue("CreateDate")
	persediaan.Suplier = r.FormValue("Suplier")
	persediaan.Telp = r.FormValue("Telp")
	persediaan.Alamat = r.FormValue("Alamat")
	persediaan.Deskripsi = r.FormValue("Deskripsi")
	fmt.Println(persediaan)

	// Save to database
	stmt, err := db.Prepare(`
	INSERT INTO persediaan(id_barang,nopol,tgl_masuk,merk,tahun,harga_modal,harga_jual,create_date,suplier,telp,alamat,
deskripsi)
		VALUES(?,?,?,?,?,?,?,?,?,?,?,?)
	`)
	if err != nil {
		fmt.Println("Prepare query error")
		panic(err)
	}
	_, err = stmt.Exec(persediaan.Id, persediaan.Nopol, persediaan.TglMasuk,
		persediaan.Merk, persediaan.Tahun, persediaan.HargaModal, persediaan.HargaJual, persediaan.CreateDate, persediaan.Suplier, persediaan.Telp, persediaan.Alamat, persediaan.Deskripsi)
	if err != nil {
		fmt.Println("Execute query error")
		panic(err)
	}
	http.Redirect(w, r, "/", 301)
}

func editpsdaanHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	var funcMap = template.FuncMap{
		"multiplication": func(n float64, f float64) float64 {
			return n * f
		},
		"addOne": func(n int) int {
			return n + 1
		},
	}
	id := r.FormValue("id")
	rows, err := db.Query(
		`SELECT *
		FROM persediaan
		WHERE id_barang = ` + id + `;`)
	if err != nil {
		log.Fatalln(err)
	}

	var persediaans []Persediaan
	var persediaan Persediaan
	for rows.Next() {
		err = rows.Scan(&persediaan.Id, &persediaan.Nopol, &persediaan.TglMasuk,
			&persediaan.Merk, &persediaan.Tahun, &persediaan.HargaModal, &persediaan.HargaJual, &persediaan.CreateDate, &persediaan.Suplier, &persediaan.Telp, &persediaan.Alamat, &persediaan.Deskripsi)
		checkInternalServerError(err, w)
		persediaans = append(persediaans, persediaan)
	}
	t, err := template.New("editpsdaan.html").Funcs(funcMap).ParseFiles("frontend/editpsdaan.html")
	checkInternalServerError(err, w)
	err = t.Execute(w, persediaans)
	checkInternalServerError(err, w)
}

func updatepsdaanHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
	}
	_, err := db.Exec(
		"UPDATE persediaan SET nopol=?, tgl_masuk=?, merk=?, tahun=?, harga_modal=?, harga_jual=?, create_date=?, suplier=?, telp=?, alamat=?, deskripsi=? WHERE id_barang=? ",
		r.FormValue("Nopol"),
		r.FormValue("TglMasuk"),
		r.FormValue("Merk"),
		r.FormValue("Tahun"),
		r.FormValue("HargaModal"),
		r.FormValue("HargaJual"),
		r.FormValue("CreateDate"),
		r.FormValue("Suplier"),
		r.FormValue("Telp"),
		r.FormValue("Alamat"),
		r.FormValue("Deskripsi"),
		r.FormValue("Id"),
	)
	fmt.Println(r)

	// checkErr(er)
	// var persediaan Persediaan
	// persediaan.Id, _ = strconv.ParseInt(r.FormValue("Id"), 10, 64)
	// persediaan.Nopol = r.FormValue("Nopol")
	// persediaan.TglMasuk = r.FormValue("TglMasuk")
	// persediaan.Merk = r.FormValue("Merk")
	// persediaan.Tahun = r.FormValue("Tahun")
	// persediaan.HargaModal, _ = strconv.ParseFloat(r.FormValue("HargaModal"), 64)
	// persediaan.HargaJual, _ = strconv.ParseFloat(r.FormValue("HargaJual"), 64)
	// persediaan.CreateDate = r.FormValue("CreateDate")
	// persediaan.Suplier = r.FormValue("Suplier")
	// persediaan.Telp = r.FormValue("Telp")
	// persediaan.Alamat = r.FormValue("Alamat")
	// persediaan.Deskripsi = r.FormValue("Deskripsi")
	// fmt.Println(persediaan)
	// stmt, err := db.Prepare(`
	// 	UPDATE persediaan SET nopol=?, tgl_masuk=?, merk=?, tahun=?, harga_modal=?, harga_jual=?, create_date=?, suplier=?, telp=?, telp=?, alamat=?, deskripsi=?
	// 	WHERE id_barang=?
	// `)
	// checkInternalServerError(err, w)
	// res, err := stmt.Exec(persediaan.Nopol, persediaan.TglMasuk, persediaan.Merk, persediaan.Tahun, persediaan.HargaModal, persediaan.HargaJual, persediaan.CreateDate, persediaan.Suplier, persediaan.Telp, persediaan.Alamat, persediaan.Deskripsi, persediaan.Id)
	// checkInternalServerError(err, w)
	// _, err = res.RowsAffected()
	checkInternalServerError(err, w)
	http.Redirect(w, r, "/", 301)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
	}
	var persediaanId, _ = strconv.ParseInt(r.FormValue("id"), 10, 64)
	fmt.Println(persediaanId)
	stmt, err := db.Prepare("DELETE FROM persediaan WHERE id_barang=?")
	checkInternalServerError(err, w)
	res, err := stmt.Exec(persediaanId)
	checkInternalServerError(err, w)
	_, err = res.RowsAffected()
	checkInternalServerError(err, w)
	http.Redirect(w, r, "/", 301)

}

// Penjualan Handler
func lpenjHandler(w http.ResponseWriter, r *http.Request) {
	// isAuthenticated(w, r)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
	rows, err := db.Query("SELECT * FROM penjualan")
	checkInternalServerError(err, w)
	var funcMap = template.FuncMap{
		"multiplication": func(n float64, f float64) float64 {
			return n * f
		},
		"addOne": func(n int) int {
			return n + 1
		},
	}
	var penjualans []Penjualan
	var penjualan Penjualan
	for rows.Next() {
		err = rows.Scan(&penjualan.Id, &penjualan.TglKeluar, &penjualan.KodeBon, &penjualan.Nopol, &penjualan.Merk, &penjualan.Suplier, &penjualan.HargaModal, &penjualan.HargaJual)
		checkInternalServerError(err, w)
		penjualans = append(penjualans, penjualan)
	}
	t, err := template.New("lpenj.html").Funcs(funcMap).ParseFiles("frontend/lpenj.html")
	checkInternalServerError(err, w)
	err = t.Execute(w, penjualans)
	checkInternalServerError(err, w)

}

func createPenjHandler(w http.ResponseWriter, r *http.Request) {
	//isAuthenticated(w, r)
	if r.Method != "POST" {
		//http.Redirect(w, r, "/", 301)
		http.ServeFile(w, r, "frontend/createpenjualan.html")
		return
	}
	var penjualan Penjualan
	penjualan.TglKeluar = r.FormValue("TglKeluar")
	penjualan.KodeBon = r.FormValue("KodeBon")
	penjualan.Nopol = r.FormValue("Nopol")
	penjualan.Merk = r.FormValue("Merk")
	penjualan.Suplier = r.FormValue("Suplier")
	penjualan.HargaModal, _ = strconv.ParseFloat(r.FormValue("HargaModal"), 64)
	penjualan.HargaJual, _ = strconv.ParseFloat(r.FormValue("HargaJual"), 64)
	fmt.Println(penjualan)

	// Save to database
	stmt, err := db.Prepare(`
	INSERT INTO penjualan(id_penjualan,tgl_keluar,kode_bon,nopol,merk,suplier,harga_modal,harga_jual)
		VALUES(?,?,?,?,?,?,?,?)
	`)
	if err != nil {
		fmt.Println("Prepare query error")
		panic(err)
	}
	_, err = stmt.Exec(penjualan.TglKeluar, penjualan.KodeBon, penjualan.Nopol,
		penjualan.Merk, penjualan.Suplier, penjualan.HargaModal, penjualan.HargaJual)
	if err != nil {
		fmt.Println("Execute query error")
		panic(err)
	}
	http.Redirect(w, r, "/", 301)
}

func editpenjHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	var funcMap = template.FuncMap{
		"multiplication": func(n float64, f float64) float64 {
			return n * f
		},
		"addOne": func(n int) int {
			return n + 1
		},
	}
	id := r.FormValue("id")
	rows, err := db.Query(
		`SELECT *
		FROM penjualan
		WHERE id_penjualan = ` + id + `;`)
	if err != nil {
		log.Fatalln(err)
	}

	var penjualans []Penjualan
	var penjualan Penjualan
	for rows.Next() {
		err = rows.Scan(&penjualan.Id, &penjualan.TglKeluar, &penjualan.KodeBon, &penjualan.Nopol,
			&penjualan.Merk, &penjualan.Suplier, &penjualan.HargaModal, &penjualan.HargaJual)
		checkInternalServerError(err, w)
		penjualans = append(penjualans, penjualan)
	}
	t, err := template.New("editpenj.html").Funcs(funcMap).ParseFiles("frontend/editpenj.html")
	checkInternalServerError(err, w)
	err = t.Execute(w, penjualans)
	checkInternalServerError(err, w)
}

func updatepenjHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
	}
	_, err := db.Exec(
		"UPDATE penjualan SET tgl_keluar=?, kode_bon=?, nopol=?, merk=?, suplier=?, harga_modal=?, harga_jual=? WHERE id_penjualan=? ",
		r.FormValue("TglKeluar"),
		r.FormValue("KodeBon"),
		r.FormValue("Nopol"),
		r.FormValue("Merk"),
		r.FormValue("Suplier"),
		r.FormValue("HargaModal"),
		r.FormValue("HargaJual"),
		r.FormValue("Id"),
	)
	fmt.Println(r)
	checkInternalServerError(err, w)
	http.Redirect(w, r, "/", 301)
}

// pengguna Handler
func lpengHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
	rows, err := db.Query("SELECT * FROM users")
	checkInternalServerError(err, w)
	var funcMap = template.FuncMap{
		"multiplication": func(n float64, f float64) float64 {
			return n * f
		},
		"addOne": func(n int) int {
			return n + 1
		},
	}
	var users []User
	var user User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Nama, &user.Username, &user.Email, &user.Password, &user.Role)
		checkInternalServerError(err, w)
		users = append(users, user)
	}
	t, err := template.New("lpeng.html").Funcs(funcMap).ParseFiles("frontend/lpeng.html")
	checkInternalServerError(err, w)
	err = t.Execute(w, users)
	checkInternalServerError(err, w)

}

func createPengHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "frontend/createpengguna.html")
		return
	}
	// grab user info
	nama := r.FormValue("Nama")
	username := r.FormValue("Username")
	email := r.FormValue("Email")
	password := r.FormValue("Password")
	role := r.FormValue("Role")
	// Check existence of user
	var user User
	err := db.QueryRow("SELECT nama, username, email, password, role FROM users WHERE username=?",
		username).Scan(&user.Nama, &user.Username, &user.Email, &user.Password, &user.Role)
	switch {
	// user is available
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		checkInternalServerError(err, w)
		// insert to database
		_, err = db.Exec(`INSERT INTO users(nama, username, email ,password, role) VALUES(?, ?, ?, ?, ?)`,
			nama, username, email, hashedPassword, role)
		fmt.Println("Created user: ", username)
		checkInternalServerError(err, w)
	case err != nil:
		http.Error(w, "loi: "+err.Error(), http.StatusBadRequest)
		return
	default:
		http.Redirect(w, r, "/", 301)
	}
}

func editpengHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	var funcMap = template.FuncMap{
		"multiplication": func(n float64, f float64) float64 {
			return n * f
		},
		"addOne": func(n int) int {
			return n + 1
		},
	}
	id := r.FormValue("id")
	rows, err := db.Query(
		`SELECT *
		FROM users
		WHERE id = ` + id + `;`)
	if err != nil {
		log.Fatalln(err)
	}

	var users []User
	var user User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Nama, &user.Username, &user.Email, &user.Password, &user.Role)
		checkInternalServerError(err, w)
		users = append(users, user)
	}
	t, err := template.New("editpeng.html").Funcs(funcMap).ParseFiles("frontend/editpeng.html")
	checkInternalServerError(err, w)
	err = t.Execute(w, users)
	checkInternalServerError(err, w)
}

func updatepengHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("Password")), bcrypt.DefaultCost)

	_, err = db.Exec(
		"UPDATE users SET nama=?, username=?, email=?, password=?, role=? WHERE id=? ",
		r.FormValue("Nama"),
		r.FormValue("Username"),
		r.FormValue("Email"),
		hashedPassword,
		r.FormValue("Role"),
		r.FormValue("Id"),
	)
	fmt.Println(r)
	checkInternalServerError(err, w)
	http.Redirect(w, r, "/", 301)
}

func laporanHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("Password")), bcrypt.DefaultCost)

	_, err = db.Exec(
		"UPDATE users SET nama=?, username=?, email=?, password=?, role=? WHERE id=? ",
		r.FormValue("Nama"),
		r.FormValue("Username"),
		r.FormValue("Email"),
		hashedPassword,
		r.FormValue("Role"),
		r.FormValue("Id"),
	)
	fmt.Println(r)
	checkInternalServerError(err, w)
	http.Redirect(w, r, "/", 301)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)
	http.Redirect(w, r, "/list", 301)
}
