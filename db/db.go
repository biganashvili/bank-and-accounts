package db

import (
	"context"
	"log"
	"time"

	"github.com/biganashvili/bank-and-accounts/proto_files"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb+srv://testmongo:9z8A1py2jtQiguJK@cluster0.yz59zqr.mongodb.net/?retryWrites=true&w=majority"

// DBNAME Database name
const DBNAME = "bank-and-accounts"

// COLLNAME Collection name
const COLLNAME = "accounts"

var db *mongo.Database

// Connect establish a connection to database
func init() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(CONNECTIONSTRING).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(DBNAME)

}

func InsertAccount(account *proto_files.Account) (interface{}, error) {
	res, err := db.Collection(COLLNAME).InsertOne(
		context.Background(),
		bson.M{
			"id":        account.Id,
			"wallet_id": account.WalletID,
			"balance":   account.Balance,
		},
	)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

// UpdateAccount updates an existing account
func UpdateAccount(account *proto_files.Account) error {
	_, err := db.Collection(COLLNAME).UpdateOne(
		context.Background(),
		bson.M{"id": account.Id},
		bson.M{"$set": bson.M{"balance": account.Balance}},
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetAllAccounts returns all people from DB
func GetAllAccounts() ([]*proto_files.Account, error) {

	elements := []*proto_files.Account{}
	elem := proto_files.Account{}
	cur, err := db.Collection(COLLNAME).Find(context.Background(), bson.M{}, nil)
	if err != nil {
		return elements, err
	}

	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			return elements, err
		}
		elements = append(elements, &elem)
	}
	if err := cur.Err(); err != nil {
		return elements, err
	}
	cur.Close(context.Background())
	return elements, nil
}

// GetAllAccounts returns all people from DB
func GetAccount(id string) (*proto_files.Account, error) {

	account := &proto_files.Account{}
	res := db.Collection(COLLNAME).FindOne(context.Background(), bson.M{"id": id}, nil)
	err := res.Decode(&account)
	if err != nil {
		return account, err
	}

	return account, nil
}
