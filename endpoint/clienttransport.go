package endpoint

import (
	"context"
	"time"

	svc "MiniProject/git.bluebird.id/mini/produk/server"

	pb "MiniProject/git.bluebird.id/mini/produk/grpc"

	util "MiniProject/git.bluebird.id/mini/util/grpc"
	disc "MiniProject/git.bluebird.id/mini/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.ProdukService"
)

func NewGRPCProdukClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.ProdukService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addProdukEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddProdukEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addProdukEp = retry
	}

	var readProdukByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadProdukByNamaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readProdukByNamaEp = retry
	}

	var readProdukEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadProdukEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readProdukEp = retry
	}

	var updateProdukEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateProduk, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateProdukEp = retry
	}
	return ProdukEndpoint{AddProdukEndpoint: addProdukEp, ReadProdukByNamaEndpoint: readProdukByNamaEp,
		ReadProdukEndpoint: readProdukEp, UpdateProdukEndpoint: updateProdukEp}, nil
}

func encodeAddProdukRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Produk)
	return &pb.AddProdukReq{
		KodeProduk: req.KodeProduk,
		NamaProduk: req.NamaProduk,
		Keterangan: req.Keterangan,
		Status:     req.Status,
		CreateBy:   req.CreateBy,
	}, nil
}

func encodeReadProdukByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Produk)
	return &pb.ReadProdukByNamaReq{NamaProduk: req.NamaProduk}, nil
}

func encodeReadProdukRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateProdukRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Produk)
	return &pb.UpdateProdukReq{
		KodeProduk: req.KodeProduk,
		NamaProduk: req.NamaProduk,
		Keterangan: req.Keterangan,
		Status:     req.Status,
		UpdateBy:   req.UpdateBy,
	}, nil
}

func decodeProdukResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadProdukByNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadProdukByNamaResp)
	return svc.Produk{
		KodeProduk: resp.KodeProduk,
		NamaProduk: resp.NamaProduk,
		Keterangan: resp.Keterangan,
		Status:     resp.Status,
		CreateBy:   resp.CreateBy,
	}, nil
}

func decodeReadProdukResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadProdukResp)
	var rsp svc.Produks

	for _, v := range resp.AllProduk {
		itm := svc.Produk{
			KodeProduk: v.KodeProduk,
			NamaProduk: v.NamaProduk,
			Keterangan: v.Keterangan,
			Status:     v.Status,
			CreateBy:   v.CreateBy,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddProdukEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddProduk",
		encodeAddProdukRequest,
		decodeProdukResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddProduk")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddProduk",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadProdukByNamaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadProdukByNama",
		encodeReadProdukByNamaRequest,
		decodeReadProdukByNamaRespones,
		pb.ReadProdukByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadProdukByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadProdukByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadProdukEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadProduk",
		encodeReadProdukRequest,
		decodeReadProdukResponse,
		pb.ReadProdukResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadProduk")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadProduk",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateProduk(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateProduk",
		encodeUpdateProdukRequest,
		decodeProdukResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateProduk")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateProduk",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
