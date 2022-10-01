package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/akhil/go-bookstore/pkg/models"
	"github.com/akhil/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewLec models.Lec

func GetLec(w http.ResponseWriter, r *http.Request) {
	newLecs := models.GetAllLec()
	res, _ := json.Marshal(newLecs)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetLecById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lecId := vars["lecId"]
	ID, err := strconv.ParseInt(lecId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	lecDetails, _ := models.GetLecById(ID)
	res, _ := json.Marshal(lecDetails)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateLec(w http.ResponseWriter, r *http.Request) {
	CreateLec := &models.Lec{}
	utils.ParseBody(r, CreateLec)
	b := CreateLec.CreateLec()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteLec(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lecId := vars["lecId"]
	ID, err := strconv.ParseInt(lecId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	lec := models.DeleteLec(ID)
	res, _ := json.Marshal(lec)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateLec(w http.ResponseWriter, r *http.Request) {
	var UpdateLec = &models.Lec{}
	utils.ParseBody(r, UpdateLec)
	vars := mux.Vars(r)
	lecId := vars["lecId"]
	ID, err := strconv.ParseInt(lecId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	lecDetails, db := models.GetLecById(ID)
	if UpdateLec.Name != "" {
		lecDetails.Name = UpdateLec.Name
	}
	if UpdateLec.Author != "" {
		lecDetails.Author = UpdateLec.Author

	}
	if UpdateLec.Publication != "" {
		lecDetails.Publication = UpdateLec.Publication
	}
	db.Save(&lecDetails)
	res, _ := json.Marshal(lecDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
