package repo

import "time"

// Repository is parent to almost every other model
type Repository struct {
	ID        string `gorm:"primary_key"` // "ZombieHippie/hello"
	ACL       string // Who has access?
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
