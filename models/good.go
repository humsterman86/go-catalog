package models

import "gopkg.in/mgo.v2/bson"

type Good struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  []CoverImages `bson:"cover_image" json:"cover_image"`
	GoodAttribute  []GoodsAttributes `bson:"good_attribute" json:"good_attribute"`
	Description string        `bson:"description" json:"description"`
}

type CoverImages struct {
    Name string
    Link string
}

type GoodsAttributes struct {
    Name string
    Value string
}