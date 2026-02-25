package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(30);comment:用户名"`
	Email    string `gorm:"type:varchar(30);comment:邮箱"`
	Age      int64  `gorm:"type:int(3);comment:年龄"`
}

func (u *User) FindUser(db *gorm.DB, username string) error {
	return db.Where("username=?", username).First(&u).Error
}

func (u *User) CreateUser(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u *User) FindUserid(db *gorm.DB, id int64) interface{} {
	return db.Where("id=?", id).First(&u).Error
}

func (u *User) DelUser(db *gorm.DB, id int64) interface{} {
	return db.Delete(&u, id).Error
}

func (u *User) UpdateUser(db *gorm.DB, id int64) interface{} {
	return db.Where("id=?", id).Updates(&u).Error
}
func (u *User) GetUserList(db *gorm.DB, page, size int64) ([]User, error) {
	var users []User
	offset := (page - 1) * size
	err := db.Offset(int(offset)).Limit(int(size)).Find(&users).Error
	return users, err
}
