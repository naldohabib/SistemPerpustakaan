package handler

import (
	"Portofolio/SistemPerpustakaan/category"
	model2 "Portofolio/SistemPerpustakaan/model"
	"Portofolio/SistemPerpustakaan/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	categoryUsecase category.CategoryUsecase
}

func CreateCategoryHandler(resp *mux.Router, categoryUsecase category.CategoryUsecase){
	categoryHandler:= CategoryHandler{categoryUsecase}

	resp.HandleFunc("/category",categoryHandler.insert).Methods(http.MethodPost)
	resp.HandleFunc("/category/{id}", categoryHandler.getByID).Methods(http.MethodGet)
	resp.HandleFunc("/category",categoryHandler.getAll).Methods(http.MethodGet)
	resp.HandleFunc("/category/{id}", categoryHandler.update).Methods(http.MethodPut)
	resp.HandleFunc("/category/{id}",categoryHandler.delete).Methods(http.MethodDelete)
}

func (h CategoryHandler) insert(writer http.ResponseWriter, request *http.Request) {
	var category model2.Category
	err:=json.NewDecoder(request.Body).Decode(&category)
	if err != nil{
		utils.HandleError(writer,http.StatusInternalServerError,"Ooops, something error")
		fmt.Printf("[CategoryHandler.insertData] Error when decoder data from body with error: %v\n", err)
		return
	}

	data, err := h.categoryUsecase.Insert(&category)
	if err != nil {
		utils.HandleError(writer,http.StatusInternalServerError,err.Error())
		fmt.Printf("[CategoryHandler.insertData] Error when request data to usecase with error: %v\n",err)
		return
	}

	utils.HandleSuccess(writer, http.StatusOK,data)
}

func (h CategoryHandler) getByID(writer http.ResponseWriter, request *http.Request) {
	pathVar :=mux.Vars(request)
	id,err := strconv.Atoi(pathVar["id"])
	if err != nil{
		utils.HandleError(writer,http.StatusBadRequest,"ID NOT VALID !!!")
		fmt.Printf("[CategoryHandler.getByID] Error when convert pathvar with error: %v\n ",err)
	}

	data, err:= h.categoryUsecase.GetByID(id)
	if err !=nil{
		utils.HandleError(writer, http.StatusNoContent, "Ooops something error")
		fmt.Printf("[CategoryHandler.getByID] Error when request data to usecase with error %v\n",err)
		return
	}
	utils.HandleSuccess(writer,http.StatusOK,data)
}

func (h CategoryHandler) getAll(writer http.ResponseWriter, request *http.Request) {
	data,err:= h.categoryUsecase.GetAll()
	if err != nil{
		utils.HandleError(writer, http.StatusInternalServerError,"Ooops something error")

		fmt.Printf("[CategoryHandler.getAll] Error when request data to usecase with error: %v\n",err)
		return
	}
	utils.HandleSuccess(writer,http.StatusOK,data)
}

func (h CategoryHandler) update(writer http.ResponseWriter, request *http.Request) {
	pathVar :=mux.Vars(request)
	id,err := strconv.Atoi(pathVar["id"])
	if err != nil{
		utils.HandleError(writer,http.StatusBadRequest,"ID NOT VALID !!!")
		fmt.Printf("[CategoryHandler.update] Error when convert pathvar with error: %v\n ",err)
	}

	var category model2.Category
	err=json.NewDecoder(request.Body).Decode(&category)
	if err != nil{
		utils.HandleError(writer,http.StatusInternalServerError,"Ooops, something error")
		fmt.Printf("[CategoryHandler.update] Error when decoder data from body with error: %v\n", err)
		return
	}

	data, err := h.categoryUsecase.Update(id,&category)
	if err != nil {
		utils.HandleError(writer,http.StatusInternalServerError,err.Error())
		fmt.Printf("[CategoryHandler.update] Error when request data to usecase with error: %v\n",err)
		return
	}

	utils.HandleSuccess(writer, http.StatusOK,data)
}

func (h CategoryHandler) delete(writer http.ResponseWriter, request *http.Request) {
	pathVar :=mux.Vars(request)
	id,err := strconv.Atoi(pathVar["id"])
	if err != nil{
		utils.HandleError(writer,http.StatusBadRequest,"ID NOT VALID !!!")
		fmt.Printf("[CategoryHandler.delete] Error when convert pathvar with error: %v\n ",err)
	}

	err= h.categoryUsecase.Delete(id)
	if err !=nil{
		utils.HandleError(writer, http.StatusNoContent, "Ooops something error")
		fmt.Printf("[CategoryHandler.delete] Error when request data to usecase with error %v\n",err)
		return
	}
	utils.HandleSuccess(writer,http.StatusOK,nil)
}

