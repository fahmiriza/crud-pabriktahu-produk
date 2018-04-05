package endpoint

import (
	"context"
	"fmt"

	sv "MiniProject/git.bluebird.id/mini/produk/server"
)

func (ke ProdukEndpoint) AddProdukService(ctx context.Context, produk sv.Produk) error {
	_, err := ke.AddProdukEndpoint(ctx, produk)
	return err
}

func (ke ProdukEndpoint) ReadProdukByNamaService(ctx context.Context, namaproduk string) (sv.Produk, error) {
	req := sv.Produk{NamaProduk: namaproduk}
	fmt.Println(req)
	resp, err := ke.ReadProdukByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Produk)
	return cus, err
}

func (ke ProdukEndpoint) ReadProdukService(ctx context.Context) (sv.Produks, error) {
	resp, err := ke.ReadProdukEndpoint(ctx, nil)
	fmt.Println("ke resp y", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Produks), err
}

func (ke ProdukEndpoint) UpdateProdukService(ctx context.Context, kar sv.Produk) error {
	_, err := ke.UpdateProdukEndpoint(ctx, kar)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
