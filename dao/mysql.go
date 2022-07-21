package dao

import (
	"Chess/module"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

type Config struct {
	user   string
	pass   string
	adrr   string
	port   string
	dbName string
	time   string
}

type DbUser struct {
	*module.SUser
}

func ConnDB() *gorm.DB {
	conf := &Config{"root", "root", "localhost", "3306", "ljj", "5s"}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", conf.user, conf.pass, conf.adrr, conf.port, conf.dbName, conf.time)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("connect failed, error=" + err.Error())
	}
	db.Create(&module.SUser{})
	db.AutoMigrate(&module.SUser{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

func CreateNewUser(UserID string, UserName string, Email string, HashedPassword []byte) {
	newUser := module.SUser{
		UserID:   UserID,
		UserName: UserName,
		Email:    Email,
		Password: string(HashedPassword),
	}
	DB.Create(&newUser)
}

func EmailSelect(email string) {
	DB.Where("email = ?", email).First(&module.SUser{})
}

func IDselect(userID string) {
	DB.Where("UserID=?", userID).First(&module.SUser{})
}

func UpdateInfo(userID string, NewUserName string, NewHashedPassword []byte, NewEmail string) {
	user := module.SUser{
		UserID: userID,
	}
	DB.Model(&user).Updates(map[string]interface{}{"UserName": NewUserName, "Password": NewHashedPassword, "Email": NewEmail}).First(&module.SUser{})
}
