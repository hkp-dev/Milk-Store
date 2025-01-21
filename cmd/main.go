package main

import (
	"app/database"
	"app/utils"
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

const pin = "310303"

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
	if enterPin() {
		menu()
	} else {
		fmt.Println("Exiting...")
		utils.Loading(1 * time.Second)
		utils.ClearCmd()
		os.Exit(0)
	}
}
func menu() {
	for {
		utils.ClearCmd()
		fmt.Println("Welcome to the Milk-Store")
		fmt.Println("1. Enroll user")
		fmt.Println("2. Find customer by phone number")
		fmt.Println("3. Listing all users")
		fmt.Println("4. Exit")
		fmt.Println("5. Help")
		fmt.Print("Enter your choice(1-5): ")

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
		case 5:
			utils.ClearCmd()
			displayHelp()
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
	}
}
func displayHelp() {
	fmt.Println("Help - User Guide")
	fmt.Println("1. Enroll user:")
	fmt.Println("   - To enroll a new user, choose option 1 from the menu.")
	fmt.Println("   - Enter the user's full name, phone number, and gender.")
	fmt.Println("   - The system will check if the user already exists. If not, the user will be added.")
	fmt.Println()
	fmt.Println("2. Find customer by phone number:")
	fmt.Println("   - Choose option 2 from the menu to find a user by phone number.")
	fmt.Println("   - Enter the phone number, and the system will display the user's details.")
	fmt.Println()
	fmt.Println("3. Listing all users:")
	fmt.Println("   - Option 3 allows you to see a list of all users in the system.")
	fmt.Println("   - This will display the full name, phone number, and gender of each user.")
	fmt.Println()
	fmt.Println("4. Exit:")
	fmt.Println("   - Option 4 will exit the program.")
	fmt.Println()
	fmt.Println("5. Help:")
	fmt.Println("   - Option 5 shows this guide to help you navigate the system.")
}

func enterPin() bool {
	fmt.Print("Enter your 6-digit PIN code to log in: ")
	enterCount := 3

	for enterCount > 0 {
		bytePin, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			fmt.Println("Error reading pin code: ", err)
			return false
		}
		enteredPin := string(bytePin)

		if enteredPin == pin {
			return true
		}

		enterCount--
		if enterCount > 0 {
			fmt.Printf("Invalid pin code. Please try again (%d attempts left).\n", enterCount)
			fmt.Print("Enter your 6-digit PIN code to log in: ")
		}
	}

	utils.ClearCmd()
	fmt.Println("Exceeded maximum number of attempts. Exiting...")
	os.Exit(1)

	return false
}
