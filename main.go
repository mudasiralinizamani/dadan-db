package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"sync"
)

const Version = "1.0.0"

type (
	Logger interface {
		Fetal(string, ...interface{})
		Error(string, ...interface{})
		Warning(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}
	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir := filepath.Clean(dir)

	opts :=  
}

func (d *Driver) Write() error {

}

func (d *Driver) Read() error {

}

func (d *Driver) ReadAll() {

}

func (d *Driver) Delete() error {

}

func (d *Driver) GetOrCreateMutex() *sync.Mutex {

}

type Address struct {
	City    string
	State   string
	Country string
	Pincode string
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)

	if err != nil {
		log.Panic(err.Error())
	}

	employees := []User{
		{"Mudasir", "24", "24125123123", "Tech Mud", Address{"Hyderabad", "Sindh", "Pakistan", "17000"}},
		{"Irfan", "24", "24125123123", "Tech Mud", Address{"Karachi", "Sindh", "Pakistan", "17000"}},
		{"Hassan", "24", "24125123123", "Tech Mud", Address{"Idk", "Sindh", "Pakistan", "17000"}},
		{"IDK", "24", "24125123123", "Tech Mud", Address{"Hello", "Sindh", "Pakistan", "17000"}},
	}

	for _, employee := range employees {
		db.Write("users", employee.Name, User{
			Name:    employee.Name,
			Age:     employee.Age,
			Contact: employee.Contact,
			Company: employee.Company,
			Address: employee.Address,
		})
	}

	records, err := db.ReadAll("users")

	if err != nil {
		log.Panic(err.Error())
	}

	allUsers := []User{}

	for _, user := range records {
		employeeFound := User{}

		if err := json.Unmarshal([]byte(user), employeeFound); err != nil {
			log.Panic(err.Error())
		}

		allUsers = append(allUsers, employeeFound)
	}
	fmt.Println(allUsers)

	// if err := db.Delete("user", "Irfan"); err != nil {
	// 	log.Panic(err.Error())
	// }

}
