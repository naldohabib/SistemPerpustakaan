package main

import (
	bookHandler "Portofolio/SistemPerpustakaan/book/handler"
	repoBook "Portofolio/SistemPerpustakaan/book/repo"
	usecaseBook "Portofolio/SistemPerpustakaan/book/usecase"

	categoryHandler "Portofolio/SistemPerpustakaan/category/handler"
	repoCategory "Portofolio/SistemPerpustakaan/category/repo"
	usecaseCategory "Portofolio/SistemPerpustakaan/category/usecase"

	userHandler "Portofolio/SistemPerpustakaan/user/handler"
	repoUser "Portofolio/SistemPerpustakaan/user/repo"
	usecaseUser "Portofolio/SistemPerpustakaan/user/usecase"

	transaksiHandler "Portofolio/SistemPerpustakaan/transaksi/handler"
	repoTransaksi "Portofolio/SistemPerpustakaan/transaksi/repo"
	usecaseTransaksi "Portofolio/SistemPerpustakaan/transaksi/usecase"

	"Portofolio/SistemPerpustakaan/driver"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	db := driver.Connect()
	defer db.Close()
	driver.InitTable(db)

	router := mux.NewRouter().StrictSlash(true)

	bookRepo := repoBook.CreateBookRepoImpl(db)
	bookUsecase := usecaseBook.CreateBookUsecaseImpl(bookRepo)

	categoryRepo := repoCategory.CreateCategoryRepoImpl(db)
	categoryUsecase := usecaseCategory.CreateCategoryUsecaseImpl(categoryRepo)

	userRepo := repoUser.CreateUserRepoImpl(db)
	userUsecase := usecaseUser.CreateUserUsecaseImpl(userRepo)

	transaksiRepo := repoTransaksi.CreateTransaksiRepoImspl(db)
	transaksiUsecase := usecaseTransaksi.CreateTransaksiUsecaseImpl(transaksiRepo, bookRepo, userRepo)

	transaksiHandler.CreateTransaksiHandler(router, transaksiUsecase)
	bookHandler.CreateBookHandler(router, bookUsecase)
	categoryHandler.CreateCategoryHandler(router, categoryUsecase)
	userHandler.CreateUserHandler(router, userUsecase)

	go serverImage()
	fmt.Println("Starting web server at port : ", port)
	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		log.Fatal()
	}
}

func serverImage() {
	fs := http.FileServer(http.Dir("./assets/"))
	http.Handle("/", fs)

	port := "9093"
	fmt.Printf("Starting Image Server at port : " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
