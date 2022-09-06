package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Users struct {
	// Id      primitive.ObjectID `bson:"_id"`
	Name    string `bson:"name"`
	Surname string `bson:"surnane"`
	Email   string `bson:"email"`
}

func mongoAdd(users []Users) {
	const uri = "mongodb://root:example@127.0.0.1:27017/"

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.Background())

	col := client.Database("reikianos").Collection("users")

	fmt.Println(users)

	for _, user := range users {
		result, err := col.InsertOne(context.Background(), user)
		if err != nil {
			fmt.Println("Error inserting into mongodb")
		}

		fmt.Println(result.InsertedID)
	}

}

func getUsers() []Users {

	db, err := sql.Open("mysql", "root:example@tcp(127.0.0.1:3307)/teste")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.SetConnMaxIdleTime(time.Minute * 3)

	rows, err := db.Query(`select * from users;`)
	if err != nil {
		panic(err)
	}

	var users []Users

	for rows.Next() {
		var user Users

		if err = rows.Scan(
			&user.Name,
			&user.Surname,
			&user.Email,
		); err != nil {
			fmt.Println("Error creating users slice from db")
		}

		users = append(users, user)
	}

	return users
}

func main() {

	users := getUsers()

	// for _, user := range users {
	// 	fmt.Println("name: ", user.Name)
	// 	fmt.Println("surname: ", user.Surname)
	// 	fmt.Println("email: ", user.Email)
	// }

	// fmt.Println(user)

	mongoAdd(users)

}
