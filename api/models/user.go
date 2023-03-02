package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	Username   string    `gorm:"size:255;not null;unique" json:"username"`
	Password   string    `gorm:"size:255;not null;" json:"-"`
	ETHAddress string    `gorm:"size:255;" json:"eth_address"`
}

type UserListing struct {
	ID         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	ETHAddress string    `json:"eth_address"`
}

type UpdateAddressInput struct {
	EthAddress string `json:"eth_address" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return
}
