package controllers
import (
	"encoding/json"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"harukaedu-users/models"
	"harukaedu-main/database"
	"harukaedu-main/errors"
	"github.com/gorilla/mux"
	"harukaedu-main/logs"
)

var db *gorm.DB
var err error

func UserIndex (response http.ResponseWriter, request *http.Request) {

	db := database.Connect()
	defer db.Close()

	var users [] models.User
	db.Find(&users)
	json.NewEncoder(response).Encode(users)
	logs.Logging("INFO", "Hit UserIndex - show list users")
}

func UserStore (response http.ResponseWriter, request *http.Request) {
	
	db := database.Connect()
	defer db.Close()

	name := request.FormValue("name")
	email := request.FormValue("email")

	if name == "" {
		errors.LogNPrint(response, "Hit UserStore - name is empty", "ERROR")
		return
	}
	if email == "" {
		errors.LogNPrint(response, "Hit UserStore - email is empty", "ERROR")
		return
	}

	db.Create(&models.User{Name: name, Email: email}) 
	errors.LogNPrint(response, "User successfully Created", "INFO")
}

func UserUpdate (response http.ResponseWriter, request *http.Request) {
	
	db := database.Connect()
	defer db.Close()

	vars := mux.Vars(request)
	id := vars["id"]
	name := request.FormValue("name")
	email := request.FormValue("email")

	var user models.User
	db.Where("ID = ?", id)

	user.Name = name
	user.Email = email

	db.Save(&user)
	errors.LogNPrint(response, "User successfully Updated", "INFO")

}

func UserDelete (response http.ResponseWriter, request *http.Request) {
	
	db := database.Connect()
	defer db.Close()

	vars := mux.Vars(request)
	id := vars["id"]

	var user models.User
	db.Where("ID = ?", id).Find(&user)
	db.Delete(&user)
	errors.LogNPrint(response, "User successfully Deleted", "INFO")
}
