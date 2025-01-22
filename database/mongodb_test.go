package database

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupDB() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Uri))
	if err != nil {
		return err
	}
	Client = client
	return nil
}

func shutDownDB() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		fmt.Println("Error while disconnecting:", err)
	}
}

func TestEnrollUser_Successfully(t *testing.T) {
	err := setupDB()
	assert.NoError(t, err)

	t.Run("TestAddUser", func(t *testing.T) {
		fullName := "John Doe"
		phoneNumber := "1234567890"
		gender := "Male"
		insertedID, err := AddUserToDatabase(fullName, phoneNumber, gender)
		assert.NoError(t, err)
		assert.NotNil(t, insertedID)
		collection := Client.Database("newdb").Collection("milk-store-hkp")
		var result bson.M
		err = collection.FindOne(context.TODO(), bson.M{"_id": insertedID}).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, fullName, result["fullName"])
		assert.Equal(t, phoneNumber, result["phone_number"])
		assert.Equal(t, gender, result["gender"])
		_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": insertedID})
		assert.NoError(t, err)
	})
	shutDownDB()
}
func TestFindCustomerByPhoneNumber_Successfully(t *testing.T) {
	err := setupDB()
	assert.NoError(t, err)
	t.Run("TestFoundUserByPhoneNumber", func(t *testing.T) {
		phoneNumber := "1234567890"
		_, err := AddUserToDatabase("John Doe", phoneNumber, "Male")
		assert.NoError(t, err)
		user, err := FoundUserByPhoneNumber(phoneNumber)
		assert.NoError(t, err)
		assert.Equal(t, user["phone_number"], phoneNumber)
		_, err = Client.Database("newdb").Collection("milk-store-hkp").DeleteOne(context.TODO(), bson.M{"phone_number": phoneNumber})
		assert.NoError(t, err)
		log.Printf("Full Name: %s, Phone Number: %s, Gender: %s\n", user["fullName"], user["phone_number"], user["gender"])
	})
	shutDownDB()
}
func TestGetAllUsers_Successfully(t *testing.T) {
	err := setupDB()
	assert.NoError(t, err)

	t.Run("TestMultipleUsersFound", func(t *testing.T) {
		collection := Client.Database("newdb").Collection("milk-store-hkp")
		_, err := collection.DeleteMany(context.TODO(), bson.M{})
		assert.NoError(t, err)

		users := []bson.M{
			{
				"fullName":     "John Doe",
				"phone_number": "1234567890",
				"gender":       "Male",
			},
			{
				"fullName":     "Jane Smith",
				"phone_number": "0987654321",
				"gender":       "Female",
			},
			{
				"fullName":     "Bob Johnson",
				"phone_number": "9876543210",
				"gender":       "Male",
			},
		}

		for _, user := range users {
			_, err := AddUserToDatabase(user["fullName"].(string), user["phone_number"].(string), user["gender"].(string))
			assert.NoError(t, err)
		}

		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		GetAllUsers()

		w.Close()
		os.Stdout = old

		var buf bytes.Buffer
		io.Copy(&buf, r)

		for _, user := range users {
			assert.Contains(t, buf.String(), fmt.Sprintf("Full Name: %s, Phone Number: %s, Gender: %s", user["fullName"], user["phone_number"], user["gender"]))
			log.Printf("Full Name: %s, Phone Number: %s, Gender: %s\n", user["fullName"], user["phone_number"], user["gender"])
		}

		_, err = collection.DeleteMany(context.TODO(), bson.M{})
		assert.NoError(t, err)
	})

	shutDownDB()
}
func TestGetAllUsers_NoUsersFound(t *testing.T) {
	err := setupDB()
	assert.NoError(t, err)

	t.Run("TestNoUsersFound", func(t *testing.T) {
		collection := Client.Database("newdb").Collection("milk-store-hkp")
		_, err := collection.DeleteMany(context.TODO(), bson.M{})
		assert.NoError(t, err)

		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		GetAllUsers()

		w.Close()
		os.Stdout = old

		var buf bytes.Buffer
		io.Copy(&buf, r)

		assert.Contains(t, buf.String(), "No users found")
		log.Print("No users found")
	})
	shutDownDB()
}
