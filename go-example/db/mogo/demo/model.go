package main

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Bookmark struct {
	ID                          bson.ObjectId `bson:"_id,omitempty"`
	Name, Description, Location string
	Priority                    int
	CreatedOn                   time.Time
	Tags                        []string
}

type BookmarkModel struct {
	C *mgo.Collection
}

func (model BookmarkModel) Create(b *Bookmark) error {
	b.ID = bson.NewObjectId()
	err := model.C.Insert(b)
	return err
}

func (model BookmarkModel) Update(b *Bookmark) error {
	err := model.C.Update(bson.M{"_id": b.ID}, bson.M{"$set": bson.M{
		"name":        b.Name,
		"description": b.Description,
		"location":    b.Location,
		"priority":    b.Priority,
		"tags":        b.Tags,
	}})
	return err
}

func (model BookmarkModel) Delete(id string) error {
	err := model.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (model BookmarkModel) GetAll() []Bookmark {
	var b []Bookmark
	iter := model.C.Find(nil).Limit(100).Skip(9999).Iter()
	result := Bookmark{}
	for iter.Next(&result) {
		b = append(b, result)
	}
	return b
}

func (model BookmarkModel) GetByOID(id string) (Bookmark, error) {
	var b Bookmark
	err := model.C.FindId(bson.ObjectIdHex(id)).One(&b)
	return b, err
}

func (model BookmarkModel) GetByTag(tags []string) []Bookmark {
	var b []Bookmark
	iter := model.C.Find(bson.M{"tags": bson.M{"$in": tags}}).Iter() //.Sort("priority", "-createdon")
	result := Bookmark{}
	for iter.Next(&result) {
		b = append(b, result)
	}
	return b
}
