package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addProduk          = `insert into produk(kodeproduk, namaproduk,keterangan,status, createby, createon)values (?,?,?,?,?,?)`
	selectProduk       = `select kodeproduk, namaproduk, keterangan, status,createby from produk Where status ='1'`
	updateProduk       = `update produk set namaproduk=?,keterangan=?, status=?, updateby=?, updateon=? where kodeproduk=?`
	selectProdukByNama = `select kodeproduk, namaproduk,keterangan, status, createon from produk where keterangan like ?`
)

//langkah 4
type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddProduk(produk Produk) error {
	fmt.Println("insert")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addProduk, produk.KodeProduk, produk.NamaProduk, produk.Keterangan, OnAdd, produk.CreateBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadProdukByNama(namaproduk string) (Produk, error) {
	fmt.Println("show by nama")
	produk := Produk{NamaProduk: namaproduk}
	err := rw.db.QueryRow(selectProdukByNama, namaproduk).Scan(&produk.KodeProduk, &produk.NamaProduk, &produk.Keterangan, &produk.Status, &produk.CreateBy)

	if err != nil {
		return Produk{}, err
	}

	return produk, nil
}

func (rw *dbReadWriter) ReadProduk() (Produks, error) {
	fmt.Println("show all")
	produk := Produks{}
	rows, _ := rw.db.Query(selectProduk)
	defer rows.Close()
	for rows.Next() {
		var k Produk
		err := rows.Scan(&k.KodeProduk, &k.NamaProduk, &k.Keterangan, &k.Status, &k.CreateBy)
		if err != nil {
			fmt.Println("error query:", err)
			return produk, err
		}
		produk = append(produk, k)
	}
	//fmt.Println("db nya:", mahasiswa)
	return produk, nil
}

func (rw *dbReadWriter) UpdateProduk(kar Produk) error {
	fmt.Println("update successfuly")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateProduk, kar.NamaProduk, kar.Keterangan, kar.Status, kar.UpdateBy, time.Now(), kar.KodeProduk)

	fmt.Println("name:", kar.NamaProduk, kar.Status)
	if err != nil {
		return err
	}

	return tx.Commit()
}
