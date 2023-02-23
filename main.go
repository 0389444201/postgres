package main

import (
	"fmt"
	"postgres/driver"
	"postgres/model"
	"postgres/repository/repoimpl"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "thai"
	dbname   = "newEx"
)

func main() {
	db := driver.Connect(host, port, user, password, dbname)
	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect successfully")
	userRepo := repoimpl.NewUserRepo(db.SQL)
	npt := model.User{
		ID:     1,
		Name:   "Noo Phuoc Thinh",
		Gender: "Male",
		Email:  "npt@gmail.com",
	}
	dt := model.User{
		ID:     2,
		Name:   "Dan Truong",
		Gender: "Male",
		Email:  "dt@gmail.com",
	}
	userRepo.Insert(npt)
	userRepo.Insert(dt)
	users, _ := userRepo.Select()
	for i := range users {
		fmt.Println(users[i])
	}
}
