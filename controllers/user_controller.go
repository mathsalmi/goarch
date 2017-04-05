package controllers

import (
	"github.com/mathsalmi/goarch/model"
	"gopkg.in/kataras/iris.v6"
)

// ListUser shows users
func ListUser(c *iris.Context) {
	orm := Getdb(c)

	var users []model.User

	err := orm.Find(&users)
	if err != nil {
		c.JSON(iris.StatusInternalServerError, nil)
		return
	}
	c.JSON(iris.StatusOK, users)
}

// CreateUser creates user
func CreateUser(c *iris.Context) {
	u := &model.User{}
	err := c.ReadJSON(u)
	if err != nil || u.IsValid() != nil {
		c.JSON(iris.StatusNotAcceptable, nil)
		return
	}

	orm := Getdb(c)
	_, err = orm.Insert(u)
	if err != nil {
		c.JSON(iris.StatusNotFound, nil)
		return
	}
	c.JSON(iris.StatusOK, nil)
}

// EditUser edits user
func EditUser(c *iris.Context) {
	id, err := c.ParamInt("id")
	if err != nil {
		c.JSON(iris.StatusBadRequest, nil)
		return
	}

	u := &model.User{}
	err = c.ReadJSON(u)
	if err != nil || u.IsValid() != nil {
		c.JSON(iris.StatusNotAcceptable, nil)
		return
	}

	orm := Getdb(c)
	_, err = orm.Where("id = ?", id).Update(*u)
	if err != nil {
		c.JSON(iris.StatusNotFound, nil)
		return
	}
	c.JSON(iris.StatusOK, nil)
}

// DeleteUser delete user
func DeleteUser(c *iris.Context) {
	id, err := c.ParamInt("id")
	if err != nil {
		c.JSON(iris.StatusNotFound, nil)
		return
	}

	orm := Getdb(c)
	affected, err := orm.Id(id).Delete(&model.User{})
	if affected == 0 || err != nil {
		c.JSON(iris.StatusNotFound, nil)
		return
	}
	c.JSON(iris.StatusOK, nil)
}
