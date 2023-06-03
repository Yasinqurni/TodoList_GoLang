package pkg

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DefaultDao struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"createdBy"`
}

func GormInit() *gorm.DB {

	godotenv.Load()

	host := os.Getenv("GORMDB_HOST")
	user := os.Getenv("GORMDB_USER")
	pass := os.Getenv("GORMDB_PASS")
	name := os.Getenv("GORMDB_NAME")
	port := os.Getenv("GORMDB_PORT")
	sslMode := os.Getenv("GORMDB_SSL")
	tz := os.Getenv("GORMDB_TZ")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, pass, name, port, sslMode, tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		LoggerErr(err.Error())
		panic(err.Error())
	}
	Logger(fmt.Sprintf("DB %s connection open", name))
	return db
}
