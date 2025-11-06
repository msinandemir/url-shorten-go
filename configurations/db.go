package configurations

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	DB *gorm.DB
}

func ConnectDB() database {
	conectionStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		GetEnv.DB_HOST,
		GetEnv.DB_USER,
		GetEnv.DB_PASSWORD,
		GetEnv.DB_NAME,
		GetEnv.DB_PORT,
	)
	db, err := gorm.Open(postgres.Open(conectionStr), &gorm.Config{})
	if err != nil {
		panic("Veritabanına bağlanılamadı!")
	}
	fmt.Println("Veritabanına bağlanıldı.")
	return database{
		DB: db,
	}
}
