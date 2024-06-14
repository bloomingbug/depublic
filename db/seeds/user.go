package seeds

import (
	"fmt"
	"time"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUserSeeds(db *gorm.DB) {
	data := map[string]entity.User{
		"admin": {
			ID:       uuid.New(),
			Name:     "Gue Admin",
			Email:    "admin@mail.com",
			Password: "password",
			Role:     entity.Admin,
			Phone:    faker.Phonenumber(),
			Address:  faker.DomainName(),
			Avatar:   faker.Sentence(),
			Birthdate: func() *time.Time {
				date, _ := time.Parse("2006-01-02", faker.Date())
				return &date
			}(),
			Gender: entity.Male,
		},
		"buyer": {
			ID:       uuid.New(),
			Name:     "Gue Buyer",
			Email:    "buyer@mail.com",
			Password: "password",
			Role:     entity.Buyer,
			Phone:    faker.Phonenumber(),
			Address:  faker.DomainName(),
			Avatar:   faker.Sentence(),
			Birthdate: func() *time.Time {
				date, _ := time.Parse("2006-01-02", faker.Date())
				return &date
			}(),
			Gender: entity.Female,
		},
	}

	for _, data := range data {
		pw, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("Error when create user %s: %s\n", data.Name, err)
			return
		}
		data := entity.NewUser(data.Name,
			data.Email,
			string(pw),
			data.Phone,
			data.Address,
			data.Avatar,
			data.Birthdate,
			data.Gender,
			data.Role,
		)

		if err := db.Create(&data).Error; err != nil {
			fmt.Printf("Error when create user %s: %s\n", data.Name, err)
			return
		}
	}
}
