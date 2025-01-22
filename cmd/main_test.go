package main

import (
	"app/database"
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestMenu(t *testing.T) {
	database.Connect()
	defer database.Disconnect()

	databaseName := "newdb"
	collectionName := "milk-store-hkp"

	ClearCollection(databaseName, collectionName)

	tests := []struct {
		name     string
		choice   int
		expected string
	}{
		{
			name:     "Enroll user - user added successfully",
			choice:   1,
			expected: "User added successfully with ID: ",
		},
		{
			name:     "Find customer by phone number - user found",
			choice:   2,
			expected: "Full Name: John Doe, Phone Number: 9876543210, Gender: Male",
		},
		{
			name:     "Listing all users",
			choice:   3,
			expected: "Listing all users: \nFull Name: John Doe, Phone Number: 9876543210, Gender: Male",
		},
		{
			name:     "Exit",
			choice:   4,
			expected: "Exiting...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.choice == 1 {
				fullName, phoneNumber, gender := "John Doe", "9876543210", "Male"
				result, err := database.AddUserToDatabase(fullName, phoneNumber, gender)
				if err != nil {
					t.Fatalf("Lỗi khi thêm người dùng: %v", err)
				}
				if result == "" {
					t.Errorf("Mã ID người dùng rỗng")
				}
			}

			if tt.choice == 2 {
				phoneNumber := "9876543210"
				user, err := database.FoundUserByPhoneNumber(phoneNumber)
				if err != nil {
					t.Errorf("Lỗi khi tìm người dùng: %v", err)
				}
				expectedUser := "Full Name: John Doe, Phone Number: 9876543210, Gender: Male"
				actual := fmt.Sprintf("Full Name: %s, Phone Number: %s, Gender: %s", user["fullName"], user["phone_number"], user["gender"])
				if actual != expectedUser {
					t.Errorf("Mong muốn: %s, nhưng có: %s", expectedUser, actual)
				}
			}

			if tt.choice == 3 {
				database.GetAllUsers()
			}

			if tt.choice == 4 {
				fmt.Println("Exiting...")
			}
		})
	}

	ClearCollection(databaseName, collectionName)
}

func ClearCollection(dbName, collectionName string) {
	collection := database.Client.Database(dbName).Collection(collectionName)
	_, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("Lỗi khi dọn dẹp collection:", err)
	}
}
