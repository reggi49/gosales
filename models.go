package main

type Persediaan struct {
	Id         int64   `json:"id_barang"`
	Nopol      string  `json:"nopol"`
	TglMasuk   string  `json:"tgl_masuk"`
	Merk       string  `json:"merk"`
	Tahun      string  `json:"tahun"`
	HargaModal float64 `json:"harga_modal"`
	HargaJual  float64 `json:"harga_jual"`
	CreateDate string  `json:"create_date"`
	Suplier    string  `json:"suplier"`
	Telp       string  `json:"telp"`
	Alamat     string  `json:"alamat"`
	Deskripsi  string  `json:"deskripsi"`
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int64  `json:"role"`
}
