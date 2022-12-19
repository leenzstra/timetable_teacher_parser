package db

import (
	"log"
	"os"
	"time"

	"github.com/leenzstra/teacher_parser/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
		Colorful:                  true,        // Disable color
	})})

	if err != nil {
		log.Fatalln(err)
	}

	// db.Migrator().DropTable(&models.Teacher{}, &models.TeacherEvaluation{})
	db.AutoMigrate(&models.Teacher{})

	return db
}
