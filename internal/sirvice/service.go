package service

import (
	"context"

	"github.com/wilfridterry/audit-log/pkg/domain/audit"
)

type Repository interface {
	Insert(ctx context.Context, logItem *audit.LogItem) error
}

type Audit struct {
	repo Repository
}

func New(repo Repository) *Audit {
	return &Audit{repo}
}

func (s *Audit) Insert(ctx context.Context, req *audit.LogRequest) error {
	logItem := audit.LogItem{
		Entity:    req.GetEntity().String(),
		Action:    req.GetAction().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return s.repo.Insert(ctx, &logItem)
}
