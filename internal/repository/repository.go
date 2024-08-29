package repository

import (
	"audit-log/pkg/domain/audit"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Audit struct {
	db *mongo.Database
}

func New(db *mongo.Database) *Audit {
	return &Audit{db}
}

func (r *Audit) Insert(ctx context.Context, logItem *audit.LogItem) error {
	_, err := r.db.Collection("logs").InsertOne(ctx, logItem)

	return err
}