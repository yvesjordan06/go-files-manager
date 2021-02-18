package application

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	/// DB application Gorm Database
	DB  = new(gorm.DB)
	err error
)

func init() {
	/*newLogger := logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold: 1*time.Nanosecond,   // Slow SQL threshold
		LogLevel:      logger.Silent, // Log level
		Colorful:      true,         // Disable color
	},
	)*/

	log.Println("Initializing application/database")
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		//Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}
	DB.Logger.LogMode(logger.Info)
	DB = DB.Debug()

}
