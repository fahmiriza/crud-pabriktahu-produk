package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "Produk.PabrikTahu.id"
	OnAdd     Status = 1
)

type Produk struct {
	KodeProduk string
	NamaProduk string
	Keterangan string
	Status     int32
	CreateBy   string
	UpdateBy   string
}
type Produks []Produk

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

type ReadWriter interface {
	AddProduk(Produk) error
	ReadProduk() (Produks, error)
	UpdateProduk(Produk) error
	ReadProdukByNama(string) (Produk, error)
}

type ProdukService interface {
	AddProdukService(context.Context, Produk) error
	ReadProdukService(context.Context) (Produks, error)
	UpdateProdukService(context.Context, Produk) error
	ReadProdukByNamaService(context.Context, string) (Produk, error)
}
