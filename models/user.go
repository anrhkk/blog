package models

import (
	"time"
	"github.com/go-macaron/binding"
	"regexp"
	"gopkg.in/macaron.v1"
)

//用户表模型
type User struct {
	Id        int
	Username  string    `form:"username" xorm:"unique;size(20)"`
	Password  string    `form:"password" xorm:"size(32)"`
	Email     string    `form:"email" xorm:"size(32)"`
	LoginNum  int
	LastLogin time.Time `xorm:"datetime updated"`
	LastIp    string    `xorm:"size(32)"`
	Active    int8
	CreatedAt time.Time `xorm:"datetime created"`
	UpdatedAt time.Time `xorm:"datetime updated"`
}

func (self *User) Exist() (bool, error) {
	return orm.Get(self)
}
func (self *User) ExistUsername() (bool, error) {
	return orm.Get(&User{Username: self.Username})
}

func (self *User) ExistEmail() (bool, error) {
	return orm.Get(&User{Email: self.Email})
}

func (self *User) GetUser() (*User, error) {
	user := &User{}
	_, err := orm.Id(self.Id).Get(user)
	return user, err
}

func (self *User) GetUserById(id int) (*User, error) {
	user := &User{Id: id}
	_, err := orm.Get(user)
	return user, err
}

func (self *User) Insert() error {
	self.Active = 1
	_, err := orm.InsertOne(self)
	return err
}

func (self *User) Update() error {
	_, err := orm.Id(self.Id).Update(self)
	return err
}

func (self *User) Delete() error {
	_, err := orm.Delete(self)
	return err
}

type UserSignInForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserSignUpForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
}

type Password struct {
	Id              int    `form:"id" binding:"required"`
	CurrentPassword string `form:"currentPassword" binding:"required"`
	ConfirmPassword string `form:"confirmPassword" binding:"required"`
}

func (user UserSignUpForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	if len(user.Username) < 5 {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"username"},
			Classification: "ComplaintError",
			Message:        "Length of username should be longer than 5.",
		})
	}
	if len(user.Password) < 5 {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"password"},
			Classification: "ComplaintError",
			Message:        "Length of password should be longer than 5.",
		})
	}
	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(user.Email))
	if matched == false {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"email"},
			Classification: "ComplaintError",
			Message:        "Please enter a valid email address.",
		})
	}
	return errs
}

func (password Password) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	if len(password.ConfirmPassword) < 5 {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"confirmPassword"},
			Classification: "ComplaintError",
			Message:        "Length of password should be longer than 5.",
		})
	}
	return errs
}
