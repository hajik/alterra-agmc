package user

import (
	"context"
	"log"
	"test-github/domain/entity/user"
	_g "test-github/domain/repository/gorm"
	_mdb "test-github/domain/repository/mongodb"
	"test-github/middleware"
	"test-github/util/requests"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserOrm struct {
	db      *gorm.DB
	mongodb *mongo.Database
}

func (u *UserOrm) NewUserOrm() {
	// gorm
	g := _g.Gorm{}
	u.db = g.NewDb()[0]
	// mongodb
	m := _mdb.MongoDB{}
	u.mongodb, _ = m.NewDB()
}

func (u *UserOrm) Login(req *requests.UserLogin) (*user.User, error) {
	var (
		err  error
		user user.User
	)
	if err = u.db.Where("email = ? AND password = ?", req.Email, req.Password).First(&user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middleware.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := u.db.Save(user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserOrm) ListAll(req *requests.UserListAll) ([]user.User, error) {
	csr, err := u.mongodb.Collection("user").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer csr.Close(context.TODO())

	result := make([]user.User, 0)
	for csr.Next(context.TODO()) {
		var row user.User
		err := csr.Decode(&row)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, nil
}

func (u *UserOrm) StoreOne(req *requests.UserStoreOne) error {
	_, err := u.mongodb.Collection("user").InsertOne(context.TODO(), user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserOrm) Update(req *requests.UserUpdate) error {
	var selector = bson.M{"name": req.Name}
	// var changes = user.User{Name: req.Name, Email: req.Email, Password: req.Password}

	update := bson.D{
		{"$set",
			bson.D{
				{"email", req.Email},
			},
		},
	}
	_, err := u.mongodb.Collection("user").UpdateOne(context.TODO(), selector, update)
	// _, err := u.mongodb.Collection("user").UpdateOne(context.TODO(), selector, bson.M{"$set": changes})
	if err != nil {
		return nil
	}
	return nil
}

func (u *UserOrm) Delete(name string) error {
	var selector = bson.M{"name": name}
	_, err := u.mongodb.Collection("user").DeleteOne(context.TODO(), selector)
	if err != nil {
		log.Fatal(err.Error())
	}
	return nil
}
