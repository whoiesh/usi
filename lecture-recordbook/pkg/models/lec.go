package models

import (
	"github.com/akhil/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Lec struct {
	gorm.Model
	Name        string `gorm:"" json:"Name"`
	Author      string `json: "Author"`
	Publication string `json:"Publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Lec{})

}

func (l *Lec) CreateLec() *Lec {
	db.NewRecord(l)
	db.Create(&l)
	return l
}
func GetAllLec() []Lec {
	var Lecs []Lec
	db.Find(&Lecs)
	return Lecs
}
func GetLecById(Id int64) (*Lec, *gorm.DB) {
	var GetLec Lec
	db := db.Where("ID=?", Id).Find(&GetLec)
	return &GetLec, db
}

func DeleteLec(Id int64) Lec {
	var lec Lec
	db.Where("ID=?", Id).Delete(lec)
	return lec
}
