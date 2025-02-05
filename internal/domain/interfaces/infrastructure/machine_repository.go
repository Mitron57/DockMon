package infrastructure

import (
    "context"
    "dockMon/internal/domain/models"
)

type MachineRepository interface {
    Put(ctx context.Context, health *models.Machine) error
    All(ctx context.Context) (map[string]*models.Machine, error)
}
