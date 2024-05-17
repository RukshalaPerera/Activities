package model

type User struct {
	ID   uint   `gorm:"primary_key;unique;not null;auto_increment" json:"id"`
	Name string `gorm:"not null;" json:"name"`
	Nic  string `gorm:"not null;unique" json:"nic"`
}

func NewUser(id uint, name, nic string) *User {
	return &User{
		ID:   id,
		Name: name,
		Nic:  nic,
	}
}
