package repo

import "github.com/jinzhu/gorm"

// Repository is parent to almost every other model
type Repository struct {
	gorm.Model
	ID     string
	ACL    string
	Owner  string
	Name   string
	Active bool
}
