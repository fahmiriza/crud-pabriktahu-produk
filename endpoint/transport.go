package endpoint

import (
	"context"

	scv "MiniProject/git.bluebird.id/mini/produk/server"

	pb "MiniProject/git.bluebird.id/mini/produk/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcProdukServer struct {
	addProduk        grpctransport.Handler
	readProdukByNama grpctransport.Handler
	readProduk       grpctransport.Handler
	updateProduk     grpctransport.Handler
}

func NewGRPCProdukServer(endpoints ProdukEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.ProdukServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcProdukServer{
		addProduk: grpctransport.NewServer(endpoints.AddProdukEndpoint,
			decodeAddProdukRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddProduk", logger)))...),
		readProdukByNama: grpctransport.NewServer(endpoints.ReadProdukByNamaEndpoint,
			decodeReadProdukByNamaRequest,
			encodeReadProdukByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadProdukByNama", logger)))...),
		readProduk: grpctransport.NewServer(endpoints.ReadProdukEndpoint,
			decodeReadProdukRequest,
			encodeReadProdukResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadProduk", logger)))...),
		updateProduk: grpctransport.NewServer(endpoints.UpdateProdukEndpoint,
			decodeUpdateProdukRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateProduk", logger)))...),
	}
}

func decodeAddProdukRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddProdukReq)
	return scv.Produk{KodeProduk: req.GetKodeProduk(), NamaProduk: req.GetNamaProduk(), Keterangan: req.GetKeterangan(), Status: req.GetStatus(), CreateBy: req.GetCreateBy()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcProdukServer) AddProduk(ctx oldcontext.Context, produk *pb.AddProdukReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addProduk.ServeGRPC(ctx, produk)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadProdukByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadProdukByNamaReq)
	return scv.Produk{NamaProduk: req.NamaProduk}, nil
}

func decodeReadProdukRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadProdukByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Produk)
	return &pb.ReadProdukByNamaResp{
		KodeProduk: resp.KodeProduk,
		NamaProduk: resp.NamaProduk,
		Keterangan: resp.Keterangan,
		Status:     resp.Status,
		CreateBy:   resp.CreateBy,
	}, nil
}

func encodeReadProdukResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Produks)

	rsp := &pb.ReadProdukResp{}

	for _, v := range resp {
		itm := &pb.ReadProdukByNamaResp{
			KodeProduk: v.KodeProduk,
			NamaProduk: v.NamaProduk,
			Keterangan: v.Keterangan,
			Status:     v.Status,
			CreateBy:   v.CreateBy,
		}
		rsp.AllProduk = append(rsp.AllProduk, itm)
	}
	return rsp, nil
}

func (s *grpcProdukServer) ReadProdukByNama(ctx oldcontext.Context, namaproduk *pb.ReadProdukByNamaReq) (*pb.ReadProdukByNamaResp, error) {
	_, resp, err := s.readProdukByNama.ServeGRPC(ctx, namaproduk)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadProdukByNamaResp), nil
}

func (s *grpcProdukServer) ReadProduk(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadProdukResp, error) {
	_, resp, err := s.readProduk.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadProdukResp), nil
}

func decodeUpdateProdukRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateProdukReq)
	return scv.Produk{KodeProduk: req.KodeProduk,
		NamaProduk: req.NamaProduk,
		Keterangan: req.Keterangan,
		Status:     req.Status,
		UpdateBy:   req.UpdateBy,
	}, nil
}

func (s *grpcProdukServer) UpdateProduk(ctx oldcontext.Context, cus *pb.UpdateProdukReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateProduk.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
