package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db            *sql.DB
	err           error
	authenticated = false
)

func main() {
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/gosales")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// test connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	os.Setenv("PORT", "8898")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	// route
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/register", registerHandler)
	//persediaan
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/createpsdaan", createPersediaanHandler)
	http.HandleFunc("/editpsdaan", editpsdaanHandler)
	http.HandleFunc("/updatepsdaan", updatepsdaanHandler)
	http.HandleFunc("/delete", deleteHandler)
	//Penjualan
	http.HandleFunc("/lpenj", lpenjHandler)
	http.HandleFunc("/cpenj", createPenjHandler)
	http.HandleFunc("/epenj", editpenjHandler)
	http.HandleFunc("/upenj", updatepenjHandler)
	// http.HandleFunc("/delete", deleteHandler)
	//Pengguna
	http.HandleFunc("/lpeng", lpengHandler)
	http.HandleFunc("/cpeng", createPengHandler)
	http.HandleFunc("/epeng", editpengHandler)
	http.HandleFunc("/upeng", updatepengHandler)
	// http.HandleFunc("/delete", deleteHandler)

	http.HandleFunc("/laporan", editpengHandler)

	// //static files
	http.Handle("/statics/",
		http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))),
	)
	http.ListenAndServe(":"+port, nil)
}
