package data

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/rcoll/microservices/mongodb/users/models"
)

type UserRepository struct {
	C *mgo.Collection
}

func (u *UserRepository) GetAll() []models.User {
	var users []models.User
	iter := u.C.Find(nil).Iter()
	result := models.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (u *UserRepository) Create(user *models.User) error {
	obj_id := bson.NewObjectId()
	user.Id = obj_id
	err := u.C.Insert(&user)
	return err
}

func (u *UserRepository) Delete(id string) error {
	err := u.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}