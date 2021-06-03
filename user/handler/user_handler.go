package handler

import (
	"Portofolio/SistemPerpustakaan/middleware"
	"Portofolio/SistemPerpustakaan/model"
	"Portofolio/SistemPerpustakaan/user"
	"Portofolio/SistemPerpustakaan/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UserHandler struct
type UserHandler struct {
	userUsecase user.UserUsecase
}

// CreateUserHandler sasdadw
func CreateUserHandler(resp *mux.Router, userUsecase user.UserUsecase) {
	userHandler := UserHandler{userUsecase}

	resp.HandleFunc("/login", userHandler.loginUser).Methods(http.MethodPost)
	resp.HandleFunc("/user", userHandler.insert).Methods(http.MethodPost)
	resp.HandleFunc("/user/{id}", middleware.TokenVerifyMiddleware(userHandler.getByID)).Methods(http.MethodGet)
	resp.HandleFunc("/user", middleware.TokenVerifyMiddleware(userHandler.getAll)).Methods(http.MethodGet)
	resp.HandleFunc("/user/{id}", middleware.TokenVerifyMiddleware(userHandler.update)).Methods(http.MethodPut)
	resp.HandleFunc("/user/{id}", middleware.TokenVerifyMiddleware(userHandler.delete)).Methods(http.MethodDelete)
}

func (h UserHandler) insert(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "Oppss, something error")
		fmt.Printf("[UserHandler.InsertData] Error when decoder data from body with error : %v\n", err)
		return
	}

	dataHashed := utils.HashedPassword(writer, user)

	data, err := h.userUsecase.Insert(dataHashed)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, err.Error())
		fmt.Printf("[UserHandler.InsertData] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h UserHandler) getByID(writer http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		utils.HandleError(writer, http.StatusBadRequest, "ID not valid !!")
		fmt.Printf("[UserHandler] Error when convert pathvar with error: %v\n", err)
	}

	data, err := h.userUsecase.GetByID(id)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "Opps, something error")
		fmt.Printf("[UserHandler.getByID]Error when request data with error : %v \n", err)
		return
	}

	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h UserHandler) getAll(writer http.ResponseWriter, request *http.Request) {
	data, err := h.userUsecase.GetAll()
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "Oppss, something error ")
		fmt.Printf("[UserHandler.getAll] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h UserHandler) update(writer http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		utils.HandleError(writer, http.StatusBadRequest, "ID not valid")
		fmt.Printf("[UserHandler.Update]Error when convert pathvar with error : %v\n", err)
	}

	var user model.User
	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "Oopss, something error")
		fmt.Printf("[UserHandler.Update]Error when request data to usecase with error : %v\n", err)
		return
	}

	data, err := h.userUsecase.Update(id, &user)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, err.Error())
		fmt.Printf("[UserHandler.update] Error when request data to usecase with error : %v", err)
		return
	}
	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h UserHandler) delete(writer http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		utils.HandleError(writer, http.StatusBadRequest, "ID not valid")
		fmt.Printf("[UserHandler.delete]Error when convert pathvar with error : %v\n", err)
	}

	err = h.userUsecase.Delete(id)
	if err != nil {
		utils.HandleError(writer, http.StatusNoContent, "Oppss, something error")
		fmt.Printf("[UserHandler.delete]Error when request data to usecase with error : %v\n", err)
		return
	}
	utils.HandleSuccess(writer, http.StatusOK, nil)
}

func (h UserHandler) loginUser(writer http.ResponseWriter, request *http.Request) {
	var data model.User
	var jwt model.Jwt

	json.NewDecoder(request.Body).Decode(&data)

	isValid := utils.ValidateLogin(writer, &data)

	if !isValid {
		return
	}

	checkEmail := utils.CheckValidMail(data.Email)

	if !checkEmail {
		utils.HandleError(writer, http.StatusBadRequest, "Email Address invalid")
		return
	}

	password := data.Password

	userData, msg, err := h.userUsecase.GetUserByEmail(data.Email)

	if msg != "" {
		utils.HandleError(writer, http.StatusInternalServerError, msg)
		return
	}

	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "The user doesnt exists")
		return
	}
	isCorrect := utils.ComparePassword(userData.Password, []byte(password))
	if isCorrect {
		token, err := utils.GenerateToken(userData)
		if err != nil {
			utils.HandleError(writer, http.StatusInternalServerError, "ooppss, something error")
			fmt.Printf("[UserHandler.GenerateToken]Error occured while generate token with error :%v\n", err)
		}
		jwt.Token = token
		utils.HandleSuccess(writer, http.StatusOK, jwt)
	} else {
		utils.HandleError(writer, http.StatusInternalServerError, "Invalid Password")
	}
}
