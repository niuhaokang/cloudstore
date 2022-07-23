package global

import (
	"cloudstore/pkg/utils/database"
	"gorm.io/gorm"
)

var (
	MYSQLCONFIG = database.MYSQLDB{}
	MYSQLDB = &gorm.DB{}
)
