package transaksi

import "Portofolio/SistemPerpustakaan/model"

type TransaksiUsecase interface {
	Insert(data *model.Transaksi) (*model.Transaksi, error)
	GetAll() (*[]model.Transaksi, error)
	GetByID(id int) (*model.Transaksi, error)
	ConfirmPinjam(id int) error
	ConfirmKembali(id int) error
	Delete(id int) error
}
