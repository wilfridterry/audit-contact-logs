package main

import (
	"context"
	"log"
	"time"

	"github.com/wilfridterry/audit-log/internal/config"
	"github.com/wilfridterry/audit-log/internal/repository"
	"github.com/wilfridterry/audit-log/internal/server"
	service "github.com/wilfridterry/audit-log/internal/sirvice"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cf.DB.Username,
		Password: cf.DB.Password,
	})
	opts.ApplyURI(cf.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := dbClient.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := dbClient.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	db := dbClient.Database(cf.DB.Database)

	repo := repository.New(db)
	auditService := service.New(repo)
	auditSrv := server.NewAuditServer(auditService)
	srv := server.New(auditSrv)

	if err := srv.ListenAndServe(cf.Port); err != nil {
		panic(err)
	}
}
