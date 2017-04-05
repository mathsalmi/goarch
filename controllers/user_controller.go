package controllers

import (
	"log"
	"strconv"

	"github.com/mathsalmi/goarch/model"
	"gopkg.in/kataras/iris.v6"
)

// Index shows the index page
func Index(c *iris.Context) {

	orm := Getdb(c)
	results, err := orm.Query(`select * from user`)
	if err != nil {
		log.Fatalln("Error on fetching users: %v", err)
	}

	users := []*model.User{}

	for _, result := range results {
		id, _ := strconv.ParseInt(string(result["id"]), 10, 64)
		active, _ := strconv.ParseBool(string(result["active"]))

		users = append(users, &model.User{
			ID:       id,
			Name:     string(result["username"]),
			Password: string(result["password"]),
			Active:   active,
		})
	}

	c.JSON(iris.StatusOK, users)
}
