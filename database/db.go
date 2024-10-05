package database

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type CalendarEvent struct {
	gorm.Model
	Title       string
	Description string
	Type        string
	Date        time.Time
}

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	DB.AutoMigrate(&CalendarEvent{})
}

func GetAllEvents() []CalendarEvent {
	var events []CalendarEvent
	DB.Find(&events)
	return events
}
