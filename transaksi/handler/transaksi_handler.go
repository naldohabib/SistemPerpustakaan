package handler

import (
	"Portofolio/SistemPerpustakaan/model"
	"Portofolio/SistemPerpustakaan/transaksi"
	"Portofolio/SistemPerpustakaan/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TransaksiHandler ...
type TransaksiHandler struct {
	transaksiUsecase transaksi.TransaksiUsecase
}

// CreateTransaksiHandler ...
func CreateTransaksiHandler(resp *mux.Router, transaksiUsecase transaksi.TransaksiUsecase) {
	tranHandler := TransaksiHandler{transaksiUsecase}

	resp.HandleFunc("/transaksi", tranHandler.insertData).Methods(http.MethodPost)
	resp.HandleFunc("/transaksi", tranHandler.getAll).Methods(http.MethodGet)
	resp.HandleFunc("/transaksi/{id}", tranHandler.getTransaksiByID).Methods(http.MethodGet)
	resp.HandleFunc("/transaksi/confirm/{id}", tranHandler.confirmBorrow).Methods(http.MethodPost)
	resp.HandleFunc("/transaksi/return/{id}", tranHandler.confirmReturn).Methods(http.MethodPost)
}

func (h *TransaksiHandler) insertData(writer http.ResponseWriter, request *http.Request) {
	var transaksi model.Transaksi
	err := json.NewDecoder(request.Body).Decode(&transaksi)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "Oppss, something error")
		fmt.Printf("[transaksiHandler.InsertData] Error when decoder data from body with error : %v\n", err)
		return
	}

	data, err := h.transaksiUsecase.Insert(&transaksi)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, err.Error())
		fmt.Printf("[TransaksiHandler.InsertData] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h *TransaksiHandler) getAll(writer http.ResponseWriter, request *http.Request) {
	data, err := h.transaksiUsecase.GetAll()
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "Ooops something error")

		fmt.Printf("[TransaksiHandler.getAll] Error when request data to usecase with error: %v\n", err)
		return
	}
	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h *TransaksiHandler) getTransaksiByID(writer http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		utils.HandleError(writer, http.StatusBadRequest, "ID NOT VALID !!!")
		fmt.Printf("[BookHandler.getBookByID] Error when convert pathvar with error: %v\n ", err)
	}

	data, err := h.transaksiUsecase.GetByID(id)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "Ooops something error")
		fmt.Printf("[BookHandler.getBookByID] Error when request data to usecase with error %v\n", err)
		return
	}
	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h *TransaksiHandler) confirmBorrow(writer http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		utils.HandleError(writer, http.StatusBadRequest, "ID NOT VALID !!!")
		fmt.Printf("[TransaksiHandler.confirmBorrow] Error when convert pathvar with error: %v\n ", err)
	}

	err = h.transaksiUsecase.ConfirmPinjam(id)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, err.Error())
		fmt.Printf("[TransaksiHandler.confirmBorrow] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.HandleSuccess(writer, http.StatusOK, "Updated successfully")

}

func (h *TransaksiHandler) confirmReturn(writer http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		utils.HandleError(writer, http.StatusBadRequest, "ID NOT VALID !!!")
		fmt.Printf("[TransaksiHandler.confirmReturn] Error when convert pathvar with error: %v\n ", err)
	}

	err = h.transaksiUsecase.ConfirmKembali(id)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, err.Error())
		fmt.Printf("[TransaksiHandler.confirmReturn] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.HandleSuccess(writer, http.StatusOK, "Updated successfully")

}
