package models

import (
	"crypto/sha512"
	"fmt"
	"net/mail"
	"regexp"

	"gorm.io/gorm"
)

type User struct {
	Base
	Name            string `gorm:"uniqueIndex;not null"`
	EmailHash       string `gorm:"uniqueIndex;not null"`
	Email           string `gorm:"-"`
	PasswordSalt    string `gorm:"not null"` // Salt for password hashing
	PasswordHash    string `gorm:"not null"`
	Password        string `gorm:"-"`
	PasswordConfirm string `gorm:"-"`             // Temporary field for password updates, not persisted
	IsAdmin         bool   `gorm:"default:false"` // Indicates if the user is an admin
	IsBanned        bool   `gorm:"default:false"` // Indicates if the user is banned
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if err = u.Validate(); err != nil {
		return err
	}

	// Call Base.BeforeSave if needed
	if err = u.Base.BeforeSave(tx); err != nil {
		return err
	}

	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err = u.Base.BeforeCreate(tx); err != nil {
		return err
	}
	// Additional logic for User.BeforeCreate, if any
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// Call Base.BeforeUpdate if needed, as BeforeUpdate is not defined in Base
	if err = u.Base.BeforeUpdate(tx); err != nil {
		return err
	}
	// Additional logic for User.BeforeUpdate, if any
	return nil
}

func (u *User) Validate() error {
	if len(u.Name) == 0 {
		return fmt.Errorf("username may not be empty")
	}

	if len(u.Name) > 25 {
		return fmt.Errorf("username may not be longer than 25 characters")
	}

	if len(u.Email) == 0 && len(u.EmailHash) == 0 {
		return fmt.Errorf("email may not be empty")
	}

	if len(u.Email) > 0 {
		if matched, _ := regexp.MatchString(`^[a-zA-Z0-9][a-zA-Z0-9_-]*$`, u.Name); !matched {
			return fmt.Errorf("username may only contain alphanumeric characters, underscores, and hyphens")
		}

		if _, err := mail.ParseAddress(u.Email); err != nil {
			return fmt.Errorf("email is not valid")
		}

		// Hash the email using SHA-512
		hasher := sha512.New()
		hasher.Write([]byte(u.Email))
		u.EmailHash = fmt.Sprintf("%x", hasher.Sum(nil))
	}

	if u.Password != "" {
		if u.Password != u.PasswordConfirm {
			return fmt.Errorf("password and password confirmation do not match")
		}

		if len(u.Password) < 8 {
			return fmt.Errorf("password must be at least 8 characters long")
		}

		salt, err := randomHex(512)
		if err != nil {
			return fmt.Errorf("failed to generate password salt: %w", err)
		}
		u.PasswordSalt = salt

		hasher := sha512.New()
		hasher.Write([]byte(u.PasswordSalt + u.Password))
		u.PasswordHash = fmt.Sprintf("%x", hasher.Sum(nil))

		u.Password = ""
		u.PasswordConfirm = ""
	}

	return nil
}
