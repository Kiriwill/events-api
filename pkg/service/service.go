package service

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model

		Name      string `json:"name" validate:"required"`
		Birthdate string `json:"birthdate" validate:"required,datetime=2006-01-02"`
		Email     string `json:"email" validate:"required,email" gorm:"unique"`
		Password  string `json:"password" validate:"required,omitempty"`
	}

	Event struct {
		gorm.Model

		Category    string `json:"category" validate:"required"`
		Description string `json:"description" validate:"required"`
		Country     string `json:"country" validate:"required" gorm:"type:enum('argentina', 'brasil', 'colômbia')"`
	}

	/*
		Service interface {
			// User
			CreateUserAddress(ctx context.Context, user *User) error
			UpdateUser(tx *sql.Tx, user *User, user_id string) error
			CreateUser(tx *sql.Tx, user *User) (sql.Result, error)
			LookupUser(ctx context.Context, query string) (*User, error)
			LookupUserById(ctx context.Context, query string) (*User, error)
			DeleteUser(ctx context.Context, user_id string) (int64, error)

			// Address
			UpdateAddress(ctx context.Context, address *Address, user_id string) (int64, error)
			UpdateUserAddress(tx *sql.Tx, user_id int, address *Address) error
			UpdateUserAddresses(ctx context.Context, user *User, user_id string) error
			CreateAddress(tx *sql.Tx, user_id int64, address *Address) error
			LookupAddressById(ctx context.Context, query string) (*Address, error)
			DeleteAddress(ctx context.Context, user_id string) (int64, error)
		}
	*/
)
