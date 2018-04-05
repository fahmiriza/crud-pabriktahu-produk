package server

import (
	"context"
)

type produk struct {
	writer ReadWriter
}

func NewProduk(writer ReadWriter) ProdukService {
	return &produk{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *produk) AddProdukService(ctx context.Context, produk Produk) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddProduk(produk)
	if err != nil {
		return err
	}

	return nil
}

func (c *produk) ReadProdukService(ctx context.Context) (Produks, error) {
	kar, err := c.writer.ReadProduk()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (c *produk) UpdateProdukService(ctx context.Context, kar Produk) error {
	err := c.writer.UpdateProduk(kar)
	if err != nil {
		return err
	}
	return nil
}

func (c *produk) ReadProdukByNamaService(ctx context.Context, namaproduk string) (Produk, error) {
	kar, err := c.writer.ReadProdukByNama(namaproduk)
	//fmt.Println("mahasiswa:", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}
