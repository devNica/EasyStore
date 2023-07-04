package configurations

import (
	"fmt"

	"github.com/devnica/EasyStore/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DatabaseConnect(config Config) *gorm.DB {
	username := config.Get("DB_USER")
	password := config.Get("DB_PASSWORD")
	host := config.Get("DB_HOST")
	port := config.Get("DB_PORT")
	database := config.Get("DB_NAME")
	sslMode := config.Get("DB_SSL_MODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, database, sslMode)

	fmt.Println("DSN", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&entities.Rol{})
	err = db.AutoMigrate(&entities.AccountStatus{})
	err = db.AutoMigrate(&entities.UserAccount{})
	err = db.AutoMigrate(&entities.UserHasRole{})

	return db

}
