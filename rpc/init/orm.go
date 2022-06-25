package init

import (
	orm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var GormDB *orm.DB

const name = "user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

// OrmInit /**
func OrmInit() {
	db, err := orm.Open("mysql", name)
	if err != nil {
		panic("db init error!")
	}

	GormDB = db
}
