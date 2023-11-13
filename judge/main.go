package main

import (
	"context"
	"fmt"
	"log"

	"github.com/priyanshu360/lab-rank/judge/repository/inmemory"
	psql "github.com/priyanshu360/lab-rank/judge/repository/postgres"
	"github.com/priyanshu360/lab-rank/judge/service/executer"
	"github.com/priyanshu360/lab-rank/judge/service/queue"
	"github.com/priyanshu360/lab-rank/judge/service/watcher"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	InitDB()
	repo := psql.NewSubmissionRepository(db)
	inMemory := inmemory.NewInMemoryQueue()
	executer := executer.NewExecuter()
	queue := queue.NewQueue(repo, inMemory)
	watcher := watcher.NewWatcher(executer, *queue)
	watcher.Run(context.Background())
}

func InitDB() {
	dbURL := "postgres://baeldung:baeldung@localhost:5432/baeldung"
	var err error
	if db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{}); err != nil {
		log.Fatal(err)
	}

	err = db.Exec("SET search_path TO lab_rank").Error
	if err != nil {
		log.Fatal(err)
	}

	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tables in the database:")
	for _, table := range tables {
		fmt.Println(table)
	}

}
