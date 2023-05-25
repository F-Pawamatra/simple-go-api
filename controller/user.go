package controller

import (
	"fmt"
	"net/http"

	"github.com/f-pawamatra/rest-api/database"
	"github.com/f-pawamatra/rest-api/model"
	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetUsers(c echo.Context) error {
	db := database.ConnectDB()

	users := []model.User{}
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	db := database.ConnectDB()

	id := c.Param("id")
	user := model.User{}
	db.First(&user, id)

	return c.JSON(http.StatusOK, user)
}

func SaveUser(c echo.Context) error {
	db := database.ConnectDB()

	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	db.Save(&user)

	fmt.Println(user)
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	db := database.ConnectDB()

	id := c.Param("id")
	user := model.User{}
	db.First(&user, id)
	c.Bind(&user)
	db.Save(&user)

	fmt.Println(user)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	db := database.ConnectDB()

	id := c.Param("id")
	user := model.User{}
	db.First(&user, id).Delete(&user)

	return c.JSON(http.StatusOK, user)
}
