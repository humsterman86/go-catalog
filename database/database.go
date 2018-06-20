package database

import (
	. "../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type GoodsDatabase struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "goods"
)

// Connect to MongoDB instance from Config
func (m *GoodsDatabase) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Get all records from database
func (m *GoodsDatabase) FindAll() ([]Good, error) {
	var goods []Good
	err := db.C(COLLECTION).Find(bson.M{}).All(&goods)
	return goods, err
}

// Get record by ID
func (m *GoodsDatabase) FindById(id string) (Good, error) {
	var good Good
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&good)
	return good, err
}

// Add a new record to database
func (m *GoodsDatabase) Insert(good Good) error {
	err := db.C(COLLECTION).Insert(&good)
	return err
}

// Delete record from batabase by ID
func (m *GoodsDatabase) Delete(good Good) error {
	err := db.C(COLLECTION).Remove(&good)
	return err
}

// Update record by ID
func (m *GoodsDatabase) Update(good Good) error {
	err := db.C(COLLECTION).UpdateId(good.ID, &good)
	return err
}
