package handler

import (
	"Portofolio/SistemPerpustakaan/book"
	"Portofolio/SistemPerpustakaan/model"
	"Portofolio/SistemPerpustakaan/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// BookHandler ...
type BookHandler struct {
	bookUsecase book.BookUsecase
}

// CreateBookHandler ...
func CreateBookHandler(resp *mux.Router, bookUsecase book.BookUsecase) {
	bookHandler := BookHandler{bookUsecase}

	resp.HandleFunc("/book", bookHandler.insertData).Methods(http.MethodPost)

	resp.HandleFunc("/book", bookHandler.getAll).Methods(http.MethodGet)
	resp.HandleFunc("/book/{id}", bookHandler.getBookByID).Methods(http.MethodGet)
	resp.HandleFunc("/book/{id}", bookHandler.delete).Methods(http.MethodDelete)
	resp.HandleFunc("/book/{id}", bookHandler.updateData).Methods(http.MethodPut)
	resp.HandleFunc("/book/detail/{id}", bookHandler.getAllBookDetail).Methods(http.MethodGet)

}

func (h *BookHandler) insertData(writer http.ResponseWriter, request *http.Request) {
	book, err := h.getFormData(writer, request)
	if err != nil {
		utils.HandleError(writer, http.StatusNotFound, err.Error())
		return
	}

	data, err := h.bookUsecase.Insert(book)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, err.Error())
		fmt.Printf("[BookHandler.Insert]Error when request data to usecase with error : %w\n", err)
		return
	}

	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h *BookHandler) getFormData(writer http.ResponseWriter, request *http.Request) (*model.Book, error) {
	title := request.FormValue("title")
	author := request.FormValue("author")
	publisher := request.FormValue("publisher")
	bookyear, err := strconv.Atoi(request.FormValue("book_year"))
	if err != nil {
		return nil, fmt.Errorf("Book year must be a number")
	}
	synopsis := request.FormValue("synopsis")

	qtybook, err := strconv.Atoi(request.FormValue("qty_book"))
	if err != nil {
		return nil, fmt.Errorf("Quantity book must be a number")
	}
	image, err := utils.UploadImage(request, title, "image", "book")
	if err != nil {
		return nil, err
	}

	book := model.Book{
		Title:        title,
		Author:       author,
		Publisher:    publisher,
		BookYear:     bookyear,
		Synopsis:     synopsis,
		QtyBook:      qtybook,
		QtyAvailable: qtybook,
		Image:        image,
	}

	return &book, nil
}

func (h *BookHandler) getAll(writer http.ResponseWriter, request *http.Request) {
	data, err := h.bookUsecase.GetAll()
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "Ooops something error")

		fmt.Printf("[BookHandler.getAll] Error when request data to usecase with error: %v\n", err)
		return
	}
	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h *BookHandler) getBookByID(writer http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		utils.HandleError(writer, http.StatusBadRequest, "ID NOT VALID !!!")
		fmt.Printf("[BookHandler.getBookByID] Error when convert pathvar with error: %v\n ", err)
	}

	data, err := h.bookUsecase.GetBookByID(id)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "Ooops something error")
		fmt.Printf("[BookHandler.getBookByID] Error when request data to usecase with error %v\n", err)
		return
	}
	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h *BookHandler) getAllBookDetail(writer http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		utils.HandleError(writer, http.StatusBadRequest, "ID NOT VALID !!!")
		fmt.Printf("[BookHandler.getAllByID] Error when convert pathvar with error: %v\n ", err)
	}

	data, err := h.bookUsecase.GetAllBookDetail(id)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "Ooops something error")

		fmt.Printf("[BookHandler.getAllBookDetail] Error when request data to usecase with error: %v\n", err)
		return
	}
	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h *BookHandler) getFormDataForUpdate(writer http.ResponseWriter, request *http.Request) (*model.Book, error) {
	title := request.FormValue("title")
	author := request.FormValue("author")
	publisher := request.FormValue("publisher")
	bookyear, err := strconv.Atoi(request.FormValue("book_year"))
	if err != nil {
		return nil, fmt.Errorf("Book year must be a number")
	}
	synopsis := request.FormValue("synopsis")

	book := model.Book{
		Title:     title,
		Author:    author,
		Publisher: publisher,
		BookYear:  bookyear,
		Synopsis:  synopsis,
	}

	return &book, nil
}

func (h *BookHandler) updateData(writer http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		utils.HandleError(writer, http.StatusBadRequest, "ID NOT VALID !!!")
		fmt.Printf("[BookHandler.getBookByID] Error when convert pathvar with error: %v\n ", err)
	}

	book, err := h.getFormDataForUpdate(writer, request)
	if err != nil {
		utils.HandleError(writer, http.StatusNotFound, err.Error())
		return
	}

	image, _ := utils.UploadImage(request, book.Title, "image", "book")

	fmt.Println(image)
	if image == "" {
		book1, _ := h.bookUsecase.GetBookByID(id)
		book.Image = book1.Image
	} else {
		book.Image = image
	}

	data, err := h.bookUsecase.Update(id, book)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, err.Error())
		fmt.Printf("[BookHandler.Insert]Error when request data to usecase with error : %w\n", err)
		return
	}

	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h *BookHandler) delete(writer http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		utils.HandleError(writer, http.StatusBadRequest, "ID not valid")
		fmt.Printf("[BookHandler.delete]Error when convert pathvar with error : %v\n", err)
	}

	err = h.bookUsecase.Delete(id)
	if err != nil {
		utils.HandleError(writer, http.StatusNoContent, "Oppss, something error")
		fmt.Printf("[BookHandler.delete]Error when request data to usecase with error : %v\n", err)
		return
	}
	utils.HandleSuccess(writer, http.StatusOK, nil)
}
