package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI = "mongodb://khoauser:Sycomore22@sycomore.vn:27014/?authSource=newdb"
const dbName = "newdb"
const collectionName = "customers"

type Customer struct {
	FullName string `bson:"full_name"`
	Phone    string `bson:"phone"`
	Gender   string `bson:"gender"`
}

func main() {
	pin := generatePIN()
	showQRCode(pin)

	if !verifyPIN(pin) {
		fmt.Println("Sai mã PIN. Thoát chương trình.")
		return
	}

	client, err := connectDB()
	if err != nil {
		log.Fatal("Không thể kết nối MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database(dbName).Collection(collectionName)

	runApp(collection)
}

func generatePIN() string {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("Lỗi tạo mã PIN:", err)
	}
	return base64.StdEncoding.EncodeToString(b)[:6]
}

func showQRCode(pin string) {
	fmt.Println("Quét QR này để lấy mã PIN:")
	qrFile := "pin_qr.png"
	err := qrcode.WriteFile(pin, qrcode.Highest, 256, qrFile)
	if err != nil {
		log.Fatal("Lỗi tạo QR code:", err)
	}
	fmt.Println("Mã PIN đã được lưu vào:", qrFile)
}

func verifyPIN(correctPin string) bool {
	fmt.Print("Nhập mã PIN để tiếp tục: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input) == correctPin
}

func connectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("✅ Kết nối MongoDB thành công")
	return client, nil
}

func runApp(collection *mongo.Collection) {
	for {
		fmt.Println("\nChọn chức năng:")
		fmt.Println("1. Thêm khách hàng")
		fmt.Println("2. Tìm khách hàng theo số điện thoại")
		fmt.Println("3. Hiển thị tất cả khách hàng")
		fmt.Println("4. Thoát")
		fmt.Print("Nhập lựa chọn: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addCustomer(collection)
		case 2:
			findCustomer(collection)
		case 3:
			listCustomers(collection)
		case 4:
			fmt.Println("Thoát ứng dụng.")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ. Vui lòng nhập lại.")
		}
	}
}

func addCustomer(collection *mongo.Collection) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập tên khách hàng: ")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	fmt.Print("Nhập số điện thoại: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	fmt.Print("Nhập giới tính (male/female): ")
	gender, _ := reader.ReadString('\n')
	gender = strings.TrimSpace(gender)

	customer := Customer{FullName: fullName, Phone: phone, Gender: gender}
	_, err := collection.InsertOne(context.TODO(), customer)
	if err != nil {
		log.Println("Lỗi khi thêm khách hàng:", err)
	} else {
		fmt.Println("Khách hàng đã được thêm thành công.")
	}
}

func findCustomer(collection *mongo.Collection) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập số điện thoại cần tìm: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	var result Customer
	err := collection.FindOne(context.TODO(), bson.M{"phone": phone}).Decode(&result)
	if err != nil {
		fmt.Println("Không tìm thấy khách hàng.")
	} else {
		fmt.Printf("Tên: %s, Số điện thoại: %s, Giới tính: %s\n", result.FullName, result.Phone, result.Gender)
	}
}

func listCustomers(collection *mongo.Collection) {
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Lỗi khi lấy danh sách khách hàng:", err)
		return
	}
	defer cursor.Close(context.TODO())

		fmt.Println("Danh sách khách hàng:")
		for cursor.Next(context.TODO()) {
			var customer Customer
			cursor.Decode(&customer)
			fmt.Printf("Tên: %s, SĐT: %s, Giới tính: %s\n", customer.FullName, customer.Phone, customer.Gender)
		}
	}

