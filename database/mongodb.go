package database

import (
	"app/utils"
	"app/validate"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Uri = "mongodb://khoauser:Sycomore22@sycomore.vn:27014/?authSource=newdb"
)

var Client *mongo.Client

func Connect() {
	ctx := context.Background()
	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(Uri))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Initializing...")
	utils.Loading(2 * time.Second)
	fmt.Println("Connected to MongoDB!")
	time.Sleep(1 * time.Second)
	utils.ClearCmd()
}
func Disconnect() error {
	if err := Client.Disconnect(context.TODO()); err != nil {
		return err
	}
	fmt.Println("Disconnected from MongoDB!")
	return nil
}
func GetUserInformation() (string, string, string, error) {
	var fullName, phoneNumber, gender string

	fmt.Print("Enter user information:\n")
	fmt.Print("Full Name: ")
	fullName = utils.GetInputFromKeyboard().(string)
	err := validate.ValidateFullName(fullName)
	if err != nil {
		return "", "", "", err
	}

	fmt.Print("Phone Number: ")
	phoneNumber = utils.GetInputFromKeyboard().(string)
	err = validate.ValidatePhoneNumber(phoneNumber)
	if err != nil {
		return "", "", "", err
	}

	fmt.Print("Gender (Male/Female): ")
	gender = utils.GetInputFromKeyboard().(string)
	err = validate.ValidateGender(gender)
	if err != nil {
		return "", "", "", err
	}

	return fullName, phoneNumber, gender, nil
}
func UserExists(phoneNumber string) bool {
	collection := Client.Database("newdb").Collection("milk-store-hkp")
	var existingUser bson.M
	err := collection.FindOne(context.TODO(), bson.M{"phone_number": phoneNumber}).Decode(&existingUser)
	if err == nil {
		return true
	}
	return err != mongo.ErrNoDocuments
}
func AddUserToDatabase(fullName, phoneNumber, gender string) (interface{}, error) {
	collection := Client.Database("newdb").Collection("milk-store-hkp")
	data := bson.M{
		"fullName":     fullName,
		"phone_number": phoneNumber,
		"gender":       gender,
		"create_time":  time.Now(),
	}

	result, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}
func FoundUserByPhoneNumber(phoneNumber string) (bson.M, error) {
	collection := Client.Database("newdb").Collection("milk-store-hkp")
	var user bson.M
	err := collection.FindOne(context.TODO(), bson.M{"phone_number": phoneNumber}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("User not found")
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
func GetPhoneNumber() (string, error) {
	fmt.Print("Enter phone number of user to find: ")
	phoneNumber := utils.GetInputFromKeyboard().(string)
	err := validate.ValidatePhoneNumber(phoneNumber)
	if err != nil {
		return "", err
	}
	return phoneNumber, nil
}
func GetAllUsers() {
	collection := Client.Database("newdb").Collection("milk-store-hkp")
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Print(err)
		return
	}
	var users []bson.M
	for cursor.Next(context.TODO()) {
		var user bson.M
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Print("Error decoding user: ", err)
			return
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println("Error iterating through results:", err)
		return
	}
	if len(users) == 0 {
		fmt.Println("No users found")
		return
	}
	fmt.Printf("All users:\n")
	for _, user := range users {
		fmt.Printf("Full Name: %s, Phone Number: %s, Gender: %s\n", user["fullName"], user["phone_number"], user["gender"])
	}
}
