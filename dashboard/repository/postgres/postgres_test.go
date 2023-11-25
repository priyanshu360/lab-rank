package postgres

import (
	"context"
	"log"
	"log/slog"

	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/priyanshu360/lab-rank/dashboard/config"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetQueueData(t *testing.T) {
	if err := godotenv.Load("../../local.env"); err != nil {
		slog.Warn("Error in loading env file, Generate .env file")
	}

	dbConf := config.NewDBConfig()
	var db *gorm.DB
	var err error
	if db, err = gorm.Open(postgres.Open(dbConf.GetURL()), &gorm.Config{}); err != nil {
		log.Fatal(err)
	}

	err = db.Exec("SET search_path TO lab_rank").Error
	if err != nil {
		log.Fatal(err)
	}
	r := NewSubmissionPostgresRepo(db)
	id, err := uuid.Parse("2251d214-4546-40e5-b12f-adcaf61c4d6e")
	if err != nil {
		log.Fatal(err)
	}
	submission := models.Submission{
		ProblemID: id,
		Lang:      models.Python,
	}
	val, err := r.GetQueueData(context.TODO(), submission)

	log.Println(val)
	log.Println(err)
}
