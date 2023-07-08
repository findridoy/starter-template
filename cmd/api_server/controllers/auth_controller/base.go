package authcontroller

import (
	"context"
	"os"
	"st/ent/user"
	"st/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Register(c *fiber.Ctx) error {
	req := new(registerRequest)

	err := pkg.BindNValidate(c, req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	hashedPass := pkg.Hash(req.Password)

	err = pkg.EntClient().User.Create().SetUsername(req.Username).SetPassword(hashedPass).Exec(context.Background())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

type registerRequest struct {
	Username        string `json:"username" validate:"required_without_all=Email PhoneNumber"`
	Email           string `json:"email" validate:"omitempty,required_without_all=Username PhoneNumber,email"`
	PhoneNumber     string `json:"phone_number" validate:"omitempty,required_without_all=Email Username,numeric"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func Login(c *fiber.Ctx) error {
	req := new(loginRequest)
	err := pkg.BindNValidate(c, req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	identity := req.Identity
	pass := req.Password

	exists, err := pkg.EntClient().User.Query().Where(user.Username(identity)).Exist(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if !exists {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	user_, err := pkg.EntClient().User.Query().Where(user.Username(identity)).Only(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}


	if pkg.Hash(pass) != user_.Password {
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

	return c.JSON(fiber.Map{"token": t, "env": os.Getenv("APP_ENV")})
}

type loginRequest struct {
	Identity string `json:"identity" validate:"required"`
	Password string `json:"password" validate:"required"`
}
