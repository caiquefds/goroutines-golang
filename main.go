package main

import (
	"goroutines-golang/domain"
	"goroutines-golang/dto/request"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func eventWorker(id int, db *gorm.DB, eventChan <-chan domain.Event) {
	for ev := range eventChan {
		if err := db.Create(&ev).Error; err != nil {
			log.Printf("[Worker %d] error: %v\n", id, err)
		} else {
			log.Printf("[Worker %d] saved event: %s\n", id, ev.EventId)
		}
	}
	time.Sleep(5 * time.Second)
}

func main() {

	db, err := gorm.Open(sqlite.Open("events.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to SQLite")
	}

	db.AutoMigrate(&domain.Event{})

	eventChan := make(chan domain.Event)

	for i := 1; i <= 5; i++ {
		go eventWorker(i, db, eventChan)
	}

	router := gin.Default()

	router.POST("/event", func(c *gin.Context) {

		var request request.CreateEventRequest
		c.Bind(&request)

		event := domain.Event{
			EventId: request.Id,
			Value:   request.Value,
		}

		eventChan <- event
		c.Status(200)
	})

	router.Run(":3000")
}
