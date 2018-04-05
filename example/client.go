package main

import (
	"context"
	"time"

	cli "MiniProject/git.bluebird.id/mini/produk/endpoint"
	svc "MiniProject/git.bluebird.id/mini/produk/server"
	opt "MiniProject/git.bluebird.id/mini/util/grpc"
	util "MiniProject/git.bluebird.id/mini/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCProdukClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Produk
	client.AddProdukService(context.Background(), svc.Produk{KodeProduk: "A004", NamaProduk: "tesjuga", Keterangan: "AAA", Status: 1, CreateBy: "Karyawan"})

	//Get Produk By Nama
	//cari := "A%"
	//cusNama, _ := client.ReadProdukByNamaService(context.Background(), cari)
	//fmt.Println("produk based on namaproduk:", cusNama)

	//List Produk
	//cuss, _ := client.ReadProdukService(context.Background())
	//fmt.Println("all produks:", cuss)

	//Update Karyawan
	//client.UpdateProdukService(context.Background(), svc.Produk{KodeProduk: "98d", NamaProduk: "tesal", Keterangan: "ZOZ", UpdateBy: "Admin", Status: 1})

}
