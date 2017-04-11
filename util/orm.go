package util

import "github.com/go-xorm/xorm"

// instance holds an instance of the XORM
var instance *xorm.Engine

// SetDb sets an instance of the XORM in the internal guard
func SetDb(orm *xorm.Engine) {
	instance = orm
}

// GetDb returns the XORM instance
func GetDb() *xorm.Engine {
	return instance
}
