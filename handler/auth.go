package handler

import (
	"fmt"
	"jwt/model"
	"time"

	"encoding/json"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Users struct {
	Users model.User `json:"users"`
}

// get all users from json file and send it to client
func GetAllUsers(c *fiber.Ctx) error {
	// file, err := ioutil.ReadFile("database/users.json")
	// if err != nil {
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }
	// var userNodes UserNodes
	// if err := json.Unmarshal(file, &userNodes); err != nil {
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }
	// return c.JSON(userNodes)

	// Open our jsonFile
	// jsonFile, err := os.Open("./users.json")
	// if err != nil {
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }

	// fmt.Println("Successfully Opened users.json")

	// // defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()
	// read file
	data, err := ioutil.ReadFile("./users.json")
	if err != nil {
		fmt.Println("File not found")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	var users model.User
	body := json.Unmarshal(data, &users)
	if body == nil {
		return c.JSON(fiber.Map{"status": "success", "message": "Success get all users", "data": body})
	}
	fmt.Println("data empty2", data)
	return c.JSON(fiber.Map{"status": "success", "message": "Success get all users", "data": users})

	// we iterate through every user within our users array and
	// for i := 0; i < len(users.Users); i++ {
	// 	fmt.Println("User Email: " + users.Users[i].Email)
	// 	fmt.Println("User Password: " + users.Users[i].Password)
	// }

}

// Login get user and password
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	identity := input.Identity
	pass := input.Password
	if identity != "ender" || pass != "ender" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
