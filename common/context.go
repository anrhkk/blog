package common

import (
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
)

type Context struct {
	*macaron.Context
	Messages   []string
	Errors     []string
	FormErrors map[string]interface{}
	Response   map[string]interface{}
}

func (self *Context) init() {
	if self.Response == nil {
		self.Response = make(map[string]interface{})
	}
	if self.FormErrors == nil {
		self.FormErrors = make(map[string]interface{})
	}
}

func (self *Context) Get(key string) interface{} {
	return self.Response[key]
}

func (self *Context) Set(key string, val interface{}) {
	self.init()
	self.Response[key] = val
}

func (self *Context) Delete(key string) {
	delete(self.Response, key)
}

func (self *Context) Clear() {
	for key := range self.Response {
		self.Delete(key)
	}
}

func (self *Context) AddMessage(message string) {
	self.Messages = append(self.Messages, message)
}

func (self *Context) HasMessage() bool {
	return len(self.Messages) > 0
}

func (self *Context) ClearMessages() {
	self.Messages = self.Messages[:0]
}

func (self *Context) SetFormError(err binding.Errors) {
	self.init()
	for _, val := range err {
		for _, fieldName := range val.FieldNames {
			if _, exists := self.FormErrors[fieldName]; !exists {
				self.FormErrors[fieldName] = val.Message
			}
		}
	}
}

func (self *Context) AddFormError(field string, err string) {
	self.FormErrors[field] = err
}

func (self *Context) HasFormError() bool {
	return len(self.FormErrors) > 0
}

func (self *Context) AddError(err string) {
	self.Errors = append(self.Errors, err)
}

func (self *Context) HasError() bool {
	return len(self.Errors) > 0
}

func (self *Context) ClearErrors() {
	self.Errors = self.Errors[:0]
	for key := range self.FormErrors {
		if _, exists := self.FormErrors[key]; exists {
			delete(self.FormErrors, key)
		}
	}
}

func InitContext() macaron.Handler {
	return func(c *macaron.Context) {
		ctx := &Context{
			Context: c,
		}
		c.Map(ctx)
	}
}
