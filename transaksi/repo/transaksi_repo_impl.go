package repo

import (
	"Portofolio/SistemPerpustakaan/model"
	"Portofolio/SistemPerpustakaan/transaksi"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// TransaksiRepoImpl ...
type TransaksiRepoImpl struct {
	db *gorm.DB
}

// CreateTransaksiRepoImpl ...
func CreateTransaksiRepoImspl(db *gorm.DB) transaksi.TransaksiRepo {
	return &TransaksiRepoImpl{db}
}

//BeginTrans represent a method to begin Transactional Operation
func (b *TransaksiRepoImpl) BeginTrans() *gorm.DB {
	return b.db.Begin()
}

// Insert ...
func (b *TransaksiRepoImpl) Insert(data *model.Transaksi) (*model.Transaksi, error) {
	var bookDetail model.BookDetail
	var user model.User
	var book model.Book

	satuJamSesudahNya := time.Now().Local().Add(time.Hour*time.Duration(1) + // jam
		time.Minute*time.Duration(0) + // menit
		time.Second*time.Duration(0), // detik
	)

	timeStampDenda := time.Now().Local().AddDate(0, 0, int(data.JumlahHari)) // 3 hari

	b.db.Where("book_id = ? AND status_book = ?", data.BookID, "available").First(&bookDetail)
	b.db.Where("id = ?", data.UserID).First(&user)
	b.db.Where("id = ?", data.BookID).First(&book)

	data.TimeStamp = satuJamSesudahNya.Unix()
	data.DateFineStamp = timeStampDenda.Unix()
	data.BookDetailID = bookDetail.ID
	data.NameBook = book.Title
	data.NameUser = user.Username

	err := b.db.Save(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[TransaksiRepoImpl.Insert] Error occured while inserting Book data to database : %w", err)
	}

	return data, nil
}

// GetAll ...
func (b *TransaksiRepoImpl) GetAll() (*[]model.Transaksi, error) {
	var data []model.Transaksi
	err := b.db.Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[TransaksiRepoImpl.GetAll] Error when query get by id with error: %w", err)
	}

	return &data, nil
}

// GetByID ...
func (b *TransaksiRepoImpl) GetByID(id int) (*model.Transaksi, error) {
	var data model.Transaksi
	waktuSaatIni := time.Now()
	err := b.db.First(&data, id).Error
	if err != nil {
		return nil, fmt.Errorf("[TransaksiRepoImpl.GetByID] Error when query get by id with error: %w", err)
	}
	if data.TimeStamp <= waktuSaatIni.Unix() {
		fmt.Println("waktu anda sudah habis.")
		b.db.Model(&data).Where("id=?", id).Update("status", "expired")
	} else {
		fmt.Println("waktu masih tersisa.")
	}
	return &data, nil
}

// ConfirmPinjam ...
func (b *TransaksiRepoImpl) ConfirmPinjam(id int) error {
	var data model.Transaksi
	var book model.Book
	var bookDetail model.BookDetail

	tx := b.db.Begin()

	tx.Find(&data).Where("id = ?", id)

	err := tx.Model(&data).Where("id=?", id).Update("status", "borrowed").Error

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[TransaksiRepoImpl.ConfirmPinjam] Error when query update status transaksi with error: %w", err)
	}

	tx.Find(&book).Where("id = ?", data.BookID)

	err = tx.Model(&bookDetail).Where("id=?", data.BookDetailID).Update("status_book", "not_available").Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[TransaksiRepoImpl.ConfirmPinjam] Error when query update status book transaksi with error: %w", err)
	}

	count := 0
	err = tx.Model(&bookDetail).Where("book_id = ? and status_book = ?", data.BookID, "available").Count(&count).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[TransaksiRepoImpl.ConfirmPinjam] Error when query check qty available with error: %w", err)
	}

	err = tx.Model(&book).Where("id=?", data.BookID).Update("qty_available", count).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[TransaksiRepoImpl.ConfirmPinjam] Error when query update book with error: %w", err)
	}

	tx.Commit()
	return nil
}

// ConfirmKembali ...
func (b *TransaksiRepoImpl) ConfirmKembali(id int) error {
	var data model.Transaksi
	var book model.Book
	var bookDetail model.BookDetail

	tx := b.db.Begin()

	tx.Find(&data).Where("id = ?", id)

	currentTime := time.Now().Unix() // tgl hari ini

	// fmt.Println(currentTime)
	// fmt.Println(tigaHariSesudahHariIni.Unix())
	result := currentTime - data.DateFineStamp
	totalHari := result / (24 * 60 * 60)
	totalDenda := totalHari * 2000
	fmt.Println(totalDenda)

	err := tx.Model(&data).Where("id=?", id).Updates(map[string]interface{}{"status": "done", "total_denda": totalDenda}).Error

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[TransaksiRepoImpl.ConfirmKembali] Error when query update status transaksi with error: %w", err)
	}

	tx.Find(&book).Where("id = ?", data.BookID)

	err = tx.Model(&bookDetail).Where("id=?", data.BookDetailID).Update("status_book", "available").Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[TransaksiRepoImpl.ConfirmPinjam] Error when query update status book transaksi with error: %w", err)
	}

	count := 0
	err = tx.Model(&bookDetail).Where("book_id = ? and status_book = ?", data.BookID, "available").Count(&count).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[TransaksiRepoImpl.ConfirmPinjam] Error when query check qty available with error: %w", err)
	}

	err = tx.Model(&book).Where("id=?", data.BookID).Update("qty_available", count).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[TransaksiRepoImpl.ConfirmPinjam] Error when query update book with error: %w", err)
	}

	tx.Commit()
	return nil
}

// Delete ...
func (b *TransaksiRepoImpl) Delete(id int) error {
	data := model.Transaksi{}
	err := b.db.Where("id=?", id).Delete(&data).Error
	if err != nil {
		return fmt.Errorf("[TransaksiRepoImpl.Delete] Error when query delete data with error: %w", err)
	}
	return nil
}
