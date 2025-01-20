package main

import (
	"app/database"
	"app/utils"
	"fmt"
	"os"
	"time"
)

var pin = "310303"

func menu() {
	for {
		utils.ClearCmd()
		fmt.Println("Welcome to the Milk-Store")
		fmt.Println("1. Enroll user")
		fmt.Println("2. Find customer by phone number")
		fmt.Println("3. Listing all users")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice(1-4): ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter your choice (1-4).")
			fmt.Println("Press Enter to continue...")
			fmt.Scanln()
			continue
		}
		switch choice {
		case 1:
			fullName, phoneNumber, gender, err := database.GetUserInformation()
			if err != nil {
				fmt.Println(err)
				fmt.Println("Press Enter to continue...")
				fmt.Scanln()
				continue
			}
			// if !database.ConfirmUserDetails(fullName, phoneNumber, gender) {
			// 	fmt.Println("Enrollment cancelled.")
			// 	fmt.Println("Press Enter to continue...")
			// 	fmt.Scanln()
			// 	continue
			// }

			if database.UserExists(phoneNumber) {
				fmt.Println("User already exists.")
				fmt.Println("Press Enter to continue...")
				fmt.Scanln()
				continue
			}

			result, err := database.AddUserToDatabase(fullName, phoneNumber, gender)
			if err != nil {
				fmt.Println("Error adding user:", err)
				fmt.Println("Press Enter to continue...")
				fmt.Scanln()
				continue
			}
			fmt.Printf("User added successfully with ID: %s\n", result)
		case 2:
			phoneNumber, err := database.GetPhoneNumber()
			if err != nil {
				fmt.Println(err)
				fmt.Println("Press Enter to continue...")
				fmt.Scanln()
				continue
			}
			user, err := database.FoundUserByPhoneNumber(phoneNumber)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Press Enter to continue...")
				fmt.Scanln()
				continue
			}
			fmt.Printf("Full Name: %s, Phone Number: %s, Gender: %s\n", user["fullName"], user["phone_number"], user["gender"])
		case 3:
			database.GetAllUsers()
		case 4:
			fmt.Println("Exiting...")
			utils.Loading(1 * time.Second)
			utils.ClearCmd()
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
	}
}
func main() {
	fmt.Print("\\ \\        /         __ \\ _)      _) |         | \n" +
		" \\ \\  \\   / _ \\  _ \\ |   | |  _` | | __|  _` | | \n" +
		"  \\ \\  \\ /  __/  __/ |   | | (   | | |   (   | | \n" +
		"   \\_/\\_/ \\___|\\___|____/ _|\\__, |_|\\__|\\__,_|_| \n" +
		"                            |___/                \n")

	database.Connect()
	time.Sleep(1 * time.Second)
	defer database.Disconnect()
	fmt.Println("Connected to MongoDB!")
	fmt.Print("Enter your pin code to log in: ")
	var enteredPin string
	fmt.Scanln(&enteredPin)
	if enteredPin != pin {
		fmt.Println("Invalid pin code. Exiting...")
		os.Exit(1)
	}
	menu()
}
