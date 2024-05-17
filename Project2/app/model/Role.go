package model

type Role struct {
	RoleUser      string `gorm:"not null;" json:"roleUser"`
	RoleModerator string `gorm:"not null;" json:"roleModerator"`
	RoleAdmin     string `gorm:"not null;" json:"roleAdmin"`
}

func NewRole() *Role {
	return &Role{
		RoleUser:      "user",
		RoleModerator: "moderator",
		RoleAdmin:     "admin",
	}
}
