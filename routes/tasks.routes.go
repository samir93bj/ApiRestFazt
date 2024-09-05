package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samir93bj/go-gorm-restapi/commons"
	"github.com/samir93bj/go-gorm-restapi/db"
	"github.com/samir93bj/go-gorm-restapi/models"
	"gorm.io/gorm"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	result := db.DB.Find(&tasks)

	if result.Error != nil {
		commons.WriteErrorResponse(w, http.StatusInternalServerError, "An error occurred while fetching tasks")
		return
	}

	commons.WriteJSONResponse(w, http.StatusOK, tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var task models.Task
	result := db.DB.First(&task, id)

	if result.Error == gorm.ErrRecordNotFound {
		commons.WriteErrorResponse(w, http.StatusNotFound, fmt.Sprintf("Task with id: %s not found", id))
		return
	}

	if result.Error != nil {
		commons.WriteErrorResponse(w, http.StatusInternalServerError, "An error occurred while fetching task")
		return
	}

	commons.WriteJSONResponse(w, http.StatusOK, task)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	var user models.User

	json.NewDecoder(r.Body).Decode(&task)

	searchUserResult := db.DB.First(&user, task.UserID)

	if searchUserResult.Error == gorm.ErrRecordNotFound {
		commons.WriteErrorResponse(w, http.StatusConflict, fmt.Sprintf("User with id: %d not found", task.UserID))
		return
	} else if searchUserResult.Error != nil {
		commons.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("An error occurred while checking user: %v", searchUserResult.Error))
		return
	}

	result := db.DB.Create(&task)

	if result.Error != nil {
		commons.WriteErrorResponse(w, http.StatusInternalServerError, "An error occurred while creating task")
		return
	}

	commons.WriteJSONResponse(w, http.StatusCreated, task)
}
