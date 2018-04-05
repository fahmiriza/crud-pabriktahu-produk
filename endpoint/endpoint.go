package endpoint

import (
	"context"

	svc "MiniProject/git.bluebird.id/mini/produk/server"

	kit "github.com/go-kit/kit/endpoint"
)

type ProdukEndpoint struct {
	AddProdukEndpoint        kit.Endpoint
	ReadProdukByNamaEndpoint kit.Endpoint
	ReadProdukEndpoint       kit.Endpoint
	UpdateProdukEndpoint     kit.Endpoint
}

func NewProdukEndpoint(service svc.ProdukService) ProdukEndpoint {
	addProdukEp := makeAddProdukEndpoint(service)
	readProdukByNamaEp := makeReadProdukByNamaEndpoint(service)
	readProdukEp := makeReadProdukEndpoint(service)
	updateProdukEp := makeUpdateProdukEndpoint(service)
	return ProdukEndpoint{AddProdukEndpoint: addProdukEp,
		ReadProdukByNamaEndpoint: readProdukByNamaEp,
		ReadProdukEndpoint:       readProdukEp,
		UpdateProdukEndpoint:     updateProdukEp,
	}
}

func makeAddProdukEndpoint(service svc.ProdukService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Produk)
		err := service.AddProdukService(ctx, req)
		return nil, err
	}
}

func makeReadProdukByNamaEndpoint(service svc.ProdukService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Produk)
		result, err := service.ReadProdukByNamaService(ctx, req.NamaProduk)
		return result, err
	}
}

func makeReadProdukEndpoint(service svc.ProdukService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadProdukService(ctx)
		return result, err
	}
}

func makeUpdateProdukEndpoint(service svc.ProdukService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Produk)
		err := service.UpdateProdukService(ctx, req)
		return nil, err
	}
}
