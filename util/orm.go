package util

import "github.com/go-xorm/xorm"

// dbinstance holds an dbinstance of the XORM
var dbinstance *xorm.Engine

// SetDb sets an instance of the XORM in the internal guard
func SetDb(orm *xorm.Engine) {
	dbinstance = orm
}

// GetDb returns the XORM instance
func GetDb() *xorm.Engine {
	return dbinstance
}
