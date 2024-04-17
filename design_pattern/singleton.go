package design_pattern

import (
	"fmt"
)

/*
	Singleton design pattern adalah pola desain, dimana pembuatan instance object hanya dibuat satu kali sepanjang aplikasi berjalan. Jika object tersebut dibutuhkan di berbagai bagian aplikasi maka instance object tersebut tidak dibuat lagi melainkan menggunakan instance yang sudah ada. Ini dapat berguna dalam beberapa kondisi seperti manajemen koneksi dan konfigurasi.

	Dalam go sendiri singleton design pattern dapat diimplementasikan dengan bantuan package sync.Once. Meskipun begitu kita dapat membuat implementasi kita sendiri terhadap design pattern ini
*/

// database connection

/* Berikut adalah contoh sederhana database connection menggunakan singleton design pattern */
var database = Database{}

type Database struct {
	Name string
	DataSource string
}

func DBConnection() Database {
	if database == (Database{}) {
		database = Database{Name: "Postgres", DataSource: "localhost://5432"}
		fmt.Println("Call once"); return database
	}
	return database
}