package utils

import (
	"Portofolio/SistemPerpustakaan/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HandleSuccess ...
func HandleSuccess(resp http.ResponseWriter, status int, data interface{}) {
	responses := model.ResponseWrapper{
		Success: true,
		Message: "Success",
		Data:    data,
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)

	err := json.NewEncoder(resp).Encode(responses)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Oopss, something error"))
		fmt.Printf("[HandleSuccess] error when encode data with error : %v \n", err)
	}
}

//HandleError ...
func HandleError(resp http.ResponseWriter, status int, msg string) {
	responses := model.ResponseWrapper{
		Success: false,
		Message: msg,
		Data:    nil,
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)

	err := json.NewEncoder(resp).Encode(responses)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("ooppss, something error"))
		fmt.Printf("[HandleError] error when encode data with error : %v \n", err)
	}
}

// HashedPassword ...
func HashedPassword(resp http.ResponseWriter, user model.User) *model.User {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		HandleError(resp, http.StatusBadRequest, "ooppss, something when wrong")
		fmt.Println("[HashedPassword]Error occured while hashing password ", err.Error())
		return nil
	}

	user.Password = string(hashedPassword)

	return &user
}

// ComparePassword ...
func ComparePassword(HashedPassword string, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), password)
	if err != nil {
		return false
	}
	return true
}

// GenerateToken ...
func GenerateToken(user *model.User) (string, error) {

	secret := os.Getenv("SECRET")
	fmt.Println(user.Username)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatalln(err)
	}

	return tokenString, nil
}

func CheckValidMail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}

// ValidateLogin ...
func ValidateLogin(resp http.ResponseWriter, user *model.User) bool {
	if user.Email == "" {
		HandleError(resp, http.StatusBadRequest, "Email cannot be empty")
		fmt.Println("[ValidateLogin] Error when validate email")
		return false
	}
	if user.Password == "" {
		HandleError(resp, http.StatusBadRequest, "Password cannot be empty")
		fmt.Println("[ValidateLogin]")
		return false
	}
	return true
}

//UploadImage represent a method to upload image to local
func UploadImage(req *http.Request, name string, key string, dest string) (string, error) {
	fileServer := os.Getenv("IMAGE_SERVER_URL")
	file, fileHandler, err := req.FormFile(key)
	if err != nil {
		return "", fmt.Errorf("Unable to open" + strings.ToUpper(key) + " file")
	}
	defer file.Close()
	splitedName := strings.Split(fileHandler.Filename, ".")
	extension := "." + splitedName[len(splitedName)-1]
	newName := strings.ReplaceAll(name, " ", "")
	fileName := "/" + newName + "-" + strings.ToUpper(key) + extension

	tempFile := filepath.Join(dest + fileName)

	isValid, msg := ValidateImageExtension(extension, int(fileHandler.Size))
	if !isValid {
		return fileServer + tempFile, fmt.Errorf(strings.ToUpper(key) + msg)
	}

	targetFile, err := os.OpenFile("assets/"+tempFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fileServer + tempFile, fmt.Errorf("Unable to copy " + strings.ToUpper(key) + " file")
	}

	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		return tempFile, fmt.Errorf("Unable to save " + strings.ToUpper(key) + " file")
	}
	return fileServer + tempFile, nil
}

//UploadImageForUpdate represent a method to upload image to local
func UploadImageForUpdate(req *http.Request, name string, key string, dest string) (string, error) {
	fileServer := os.Getenv("IMAGE_SERVER_URL")
	file, fileHandler, err := req.FormFile(key)
	if err != nil {
		return "", fmt.Errorf("Unable to open" + strings.ToUpper(key) + " file")
	}
	defer file.Close()
	splitedName := strings.Split(fileHandler.Filename, ".")
	extension := "." + splitedName[len(splitedName)-1]
	newName := strings.ReplaceAll(name, " ", "")
	fileName := "/" + newName + "-" + strings.ToUpper(key) + extension

	tempFile := filepath.Join(dest + fileName)

	isValid, msg := ValidateImageExtension(extension, int(fileHandler.Size))
	if !isValid {
		return fileServer + tempFile, fmt.Errorf(strings.ToUpper(key) + msg)
	}

	targetFile, err := os.OpenFile("assets/"+tempFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fileServer + tempFile, fmt.Errorf("Unable to copy " + strings.ToUpper(key) + " file")
	}

	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		return tempFile, fmt.Errorf("Unable to save " + strings.ToUpper(key) + " file")
	}
	return fileServer + tempFile, nil
}

func ValidateImageExtension(image string, size int) (bool, string) {
	if image != ".jpg" && image != ".png" && image != ".jpeg" {
		return false, " file must be .jpg/.jpeg/.png"
	}
	if size >= 1000024 {
		return false, " file max size is 1MB"
	}
	return true, ""
}
