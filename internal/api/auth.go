package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/dmitryt/tiny-url-service-backend/internal/models"
	"github.com/dmitryt/tiny-url-service-backend/internal/session"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var usersStoragePath = filepath.Join("fixtures", "users.json")

var ErrCheckPassword = errors.New("error occurred during checking the password")

func getSaltedPassword(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
}

func comparePassword(hashedPassword, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pwd))
}

func readUsersFromFile() (result []models.User, err error) {
	content, err := ioutil.ReadFile(usersStoragePath)
	if err != nil {
		return
	}
	if len(content) == 0 {
		content = []byte("[]")
	}
	err = json.Unmarshal(content, &result)

	return result, logErrorIfExists(ErrContentUnmarshal, err)
}

func writeUsersToFile(data []models.User) (err error) {
	content, err := json.Marshal(data)
	if err != nil {
		return logErrorIfExists(ErrContentMarshal, err)
	}
	err = ioutil.WriteFile(usersStoragePath, content, 0o600)

	return logErrorIfExists(ErrFileWrite, err)
}

func loginHandler(c *fiber.Ctx) error {
	payload := new(models.Session)
	if err := c.BodyParser(payload); err != nil {
		return logErrorIfExists(ErrParsingBody, err)
	}
	store, err := session.Store.Get(c)
	if err != nil {
		panic(err)
	}

	users, err := readUsersFromFile()
	if err != nil {
		return err
	}

	foundUser := models.User{}
	for _, user := range users {
		if user.Username == payload.Username && comparePassword(user.Password, payload.Password) == nil {
			foundUser = user
		}
	}

	if foundUser.ID == "" {
		c.Status(http.StatusUnauthorized)

		return c.JSON(DummyResponse{OK: false, Err: "UNAUTHORIZED"})
	}

	store.Set("uid", foundUser.ID)
	err = store.Save()
	if err != nil {
		panic(err)
	}

	foundUser.Password = ""

	return c.JSON(foundUser)
}

func logoutHandler(c *fiber.Ctx) error {
	store, err := session.Store.Get(c)
	if err != nil {
		panic(err)
	}

	err = store.Destroy()
	if err != nil {
		panic(err)
	}

	return c.JSON(DummyResponse{OK: true})
}

func registerHandler(c *fiber.Ctx) error {
	payload := new(models.User)
	if err := c.BodyParser(payload); err != nil {
		return logErrorIfExists(ErrParsingBody, err)
	}
	if payload.Username == "" {
		return logErrorIfExists(ErrValidationAddItem, errors.New("invalid data"))
	}
	userPass, _ := getSaltedPassword(payload.Password)
	newUser := models.User{ID: primitive.NewObjectID().Hex(), FirstName: payload.FirstName, LastName: payload.LastName, Username: payload.Username, Password: string(userPass)}
	users, err := readUsersFromFile()
	if err != nil {
		return err
	}
	err = writeUsersToFile(append(users, newUser))
	if err != nil {
		return err
	}

	newUser.Password = ""

	return c.JSON(newUser)
}

func (p *API) HandleAuth(r fiber.Router) {
	r.Post("/login", loginHandler)
	r.Post("/logout", logoutHandler)
	r.Post("/register", registerHandler)
}
