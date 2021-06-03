package usecase

import (
	"Portofolio/SistemPerpustakaan/book"
	"Portofolio/SistemPerpustakaan/model"
	"Portofolio/SistemPerpustakaan/transaksi"
	"Portofolio/SistemPerpustakaan/user"
	"errors"
)

// TransaksiUsecaseImpl ...
type TransaksiUsecaseImpl struct {
	transaksiRepo transaksi.TransaksiRepo
	bookRepo      book.BookRepo
	userRepo      user.UserRepo
}

// CreateTransaksiUsecaseImpl ...
func CreateTransaksiUsecaseImpl(transaksiRepo transaksi.TransaksiRepo, bookRepo book.BookRepo, userRepo user.UserRepo) transaksi.TransaksiUsecase {
	return &TransaksiUsecaseImpl{transaksiRepo, bookRepo, userRepo}
}

func (b *TransaksiUsecaseImpl) Insert(data *model.Transaksi) (*model.Transaksi, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	_, err := b.bookRepo.GetBookByID(int(data.BookID))
	if err != nil {
		return nil, errors.New("book ID does not exist")
	}

	_, err = b.userRepo.GetByID(int(data.UserID))
	if err != nil {
		return nil, errors.New("User ID does not exist")
	}

	dataok, err := b.transaksiRepo.Insert(data)
	if err != nil {
		return nil, err
	}

	return dataok, nil
}

func (b *TransaksiUsecaseImpl) GetAll() (*[]model.Transaksi, error) {

	data, err := b.transaksiRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *TransaksiUsecaseImpl) GetByID(id int) (*model.Transaksi, error) {
	data, err := b.transaksiRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *TransaksiUsecaseImpl) ConfirmPinjam(id int) error {
	data, err := b.transaksiRepo.GetByID(id)
	if err != nil {
		return errors.New("TransaksiID does not exist")
	}

	if data.Status != "waiting" {
		return errors.New("Transaction status must be waiting")
	}

	err = b.transaksiRepo.ConfirmPinjam(id)
	if err != nil {
		return errors.New("not successful ConfirmBorrow")
	}

	return nil
}

func (b *TransaksiUsecaseImpl) ConfirmKembali(id int) error {
	data, err := b.transaksiRepo.GetByID(id)
	if err != nil {
		return errors.New("TransaksiID does not exist")
	}

	if data.Status != "borrowed" {
		return errors.New("Transaction status must be waiting")
	}

	err = b.transaksiRepo.ConfirmKembali(id)
	if err != nil {
		return errors.New("not successful ConfirmReturn")
	}

	return nil
}

func (b *TransaksiUsecaseImpl) Delete(id int) error {
	_, err := b.transaksiRepo.GetByID(id)
	if err != nil {
		return errors.New("TransaksiID does not exist")
	}

	err = b.transaksiRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
