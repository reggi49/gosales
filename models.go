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

type Penjualan struct {
	Id         int64   `json:"id_penjualan"`
	TglKeluar  string  `json:"tgl_keluar"`
	KodeBon    string  `json:"kode_bon"`
	Nopol      string  `json:"nopol"`
	Merk       string  `json:"merk"`
	Suplier    string  `json:"suplier"`
	HargaModal float64 `json:"harga_modal"`
	HargaJual  float64 `json:"harga_jual"`
}

type User struct {
	Id       int64  `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int64  `json:"role"`
}
