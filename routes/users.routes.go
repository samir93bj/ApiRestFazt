package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samir93bj/go-gorm-restapi/commons"
	"github.com/samir93bj/go-gorm-restapi/db"
	"github.com/samir93bj/go-gorm-restapi/models"
	"gorm.io/gorm"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	result := db.DB.Find(&users)

	if result.Error != nil {
		commons.WriteErrorResponse(w, http.StatusInternalServerError, "An error occurred while fetching users")
		return
	}

	commons.WriteJSONResponse(w, http.StatusOK, users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	result := db.DB.First(&user, id)

	if result.Error == gorm.ErrRecordNotFound {
		commons.WriteErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	if result.Error != nil {
		commons.WriteErrorResponse(w, http.StatusInternalServerError, "An error occurred while fetching user")
		return
	}

	commons.WriteJSONResponse(w, http.StatusOK, user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	result := db.DB.Where("email = ?", user.Email)

	if result.Error == nil {
		commons.WriteErrorResponse(w, http.StatusConflict, "Email already in use")
		return
	} else if result.Error != gorm.ErrRecordNotFound {
		commons.WriteErrorResponse(w, http.StatusInternalServerError, "An error occurred while checking email")
		return
	}

	createdUser := db.DB.Create(&user)

	if createdUser.Error != nil {
		commons.WriteErrorResponse(w, http.StatusInternalServerError, "An error occurred while create user")
		return
	}

	commons.WriteJSONResponse(w, http.StatusCreated, user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var id = mux.Vars(r)["id"]

	var user models.User

	result := db.DB.First(&user, id)

	if result.Error == gorm.ErrRecordNotFound {
		commons.WriteErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	if result.Error != nil {
		commons.WriteErrorResponse(w, http.StatusInternalServerError, "An error occurred while fetching user")
		return
	}

	db.DB.Delete(&user)

	commons.WriteJSONResponse(w, http.StatusOK, user)
}
