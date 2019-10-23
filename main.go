package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type account struct {
	ID             *primitive.ObjectID `json:"id" bson:"_id"`
	Name           string              `json:"name" bson:"name"`
	Logo           string              `json:"logo" bson:"logo"`
	CaasStatus     bool                `json:"caasStatus" bson:"caasStatus"`
	BillingContact []billingContact    `json:"billingContact" bson:"billingContact"`
}

type billingContact struct {
	FirstName   string `json:"firstName" bson:"firstName"`
	LastName    string `json:"lastName" bson:"lastName"`
	CompanyName string `json:"companyName" bson:"companyName"`
}

func main() {

	clientOptions := options.Client().ApplyURI(os.Getenv("mongodbString"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to MongoDB")

	collection := client.Database("dev").Collection(("accounts"))

	findOptions := options.Find()
	findOptions.SetLimit(5)

	a := []*account{}

	res, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for res.Next((context.TODO())) {
		var r account
		err := res.Decode(&r)
		if err != nil {
			log.Fatal(err)
		}

		a = append(a, &r)
	}

	if err := res.Err(); err != nil {
		log.Fatal(err)
	}

	res.Close(context.TODO())
	result, err := json.Marshal(&a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))
}
