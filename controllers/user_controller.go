package controllers

import (
	"github.com/mathsalmi/goarch/model"
	"github.com/mathsalmi/goarch/util"
	"gopkg.in/kataras/iris.v6"
)

// ListUser shows users
func ListUser(c *iris.Context) {
	orm := util.Getdb(c)

	var users []model.User

	err := orm.Find(&users)
	if err != nil {
		c.JSON(iris.StatusInternalServerError, util.NewError(err))
		return
	}
	c.JSON(iris.StatusOK, users)
}

// CreateUser creates user
func CreateUser(c *iris.Context) {
	u := &model.User{}
	err := c.ReadJSON(u)
	if err != nil {
		c.JSON(iris.StatusNotAcceptable, util.NewError(err))
		return
	}
	if valid := u.IsValid(); valid != nil {
		c.JSON(iris.StatusNotAcceptable, valid)
		return
	}

	orm := util.Getdb(c)
	_, err = orm.Insert(u)
	if err != nil {
		c.JSON(iris.StatusNotFound, util.NewError(err))
		return
	}
	c.JSON(iris.StatusOK, nil)
}

// EditUser edits user
func EditUser(c *iris.Context) {
	id, err := c.ParamInt("id")
	if err != nil {
		c.JSON(iris.StatusBadRequest, util.NewError(err))
		return
	}

	u := &model.User{}
	err = c.ReadJSON(u)
	if err != nil {
		c.JSON(iris.StatusNotAcceptable, util.NewError(err))
		return
	}
	if valid := u.IsValid(); valid != nil {
		c.JSON(iris.StatusNotAcceptable, valid)
		return
	}

	orm := util.Getdb(c)
	_, err = orm.Where("id = ?", id).Update(*u)
	if err != nil {
		c.JSON(iris.StatusNotFound, util.NewError(err))
		return
	}
	c.JSON(iris.StatusOK, nil)
}

// DeleteUser delete user
func DeleteUser(c *iris.Context) {
	id, err := c.ParamInt("id")
	if err != nil {
		c.JSON(iris.StatusNotFound, util.NewError(err))
		return
	}

	orm := util.Getdb(c)
	affected, err := orm.Id(id).Delete(&model.User{})
	if err != nil {
		c.JSON(iris.StatusNotFound, util.NewError(err))
		return
	}
	if affected == 0 {
		c.JSON(iris.StatusNotFound, util.NewError(model.ErrUserNotFound))
		return
	}

	c.JSON(iris.StatusOK, nil)
}
