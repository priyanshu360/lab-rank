package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/priyanshu360/lab-rank/judge/service/executer"
	"github.com/priyanshu360/lab-rank/judge/service/k8s"
	"github.com/priyanshu360/lab-rank/judge/service/watcher"
	q "github.com/priyanshu360/lab-rank/queue/queue"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"k8s.io/client-go/util/homedir"
)

var db *gorm.DB

func main() {
	fmt.Println("running")
	InitDB()

	// Todo : take config from env
	kubeconfig := flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "path to kubeconfig file")
	namespace := flag.String("namespace", "storage", "Kubernetes namespace")
	flag.Parse()

	k8s, err := k8s.NewKubernetesManager(*kubeconfig, *namespace)
	if err != nil {
		fmt.Printf("Error creating KubernetesManager: %v\n", err)
		os.Exit(1)
	}

	executer := executer.NewExecuter(k8s)
	// todo : make env, handle error
	consumer, err := q.InitRabbitMQConsumer("lab-rank")
	if err != nil {
		log.Fatal(err)
	}
	watcher := watcher.NewWatcher(executer, consumer)
	watcher.Run(context.Background())
}

func InitDB() {
	dbURL := "postgres://new_admin_user:your_password@localhost:5432/postgres"
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
