package services

import (
    "context"
    "dockMon/internal/domain/models"
)

type Manager interface {
    Save(ctx context.Context, machine *models.Machine) error
    Machines(ctx context.Context) ([]*models.Machine, error)
}
