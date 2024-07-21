package seeds

import (
	"log"

	"gorm.io/gorm"

	"news-topic-api/internal/entities"
)

func SeedNewsStatus(db *gorm.DB) {
	statuses := []entities.Statuses{
		{
			Title: "Published",
			Value: "published",
		},
		{
			Title: "Draft",
			Value: "draft",
		},
		{
			Title: "Deleted",
			Value: "deleted",
		},
	}

	for _, role := range statuses {
		if err := db.Create(&role).Error; err != nil {
			log.Fatalf("Could not seed news status: %v", err)
		}
	}

	log.Println("Seeding completed")
}
