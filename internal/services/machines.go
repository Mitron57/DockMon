package services

import (
    "context"
    "dockMon/internal/domain/interfaces/infrastructure"
    "dockMon/internal/domain/interfaces/services"
    "dockMon/internal/domain/models"
    "time"
)

type MachinesService struct {
    repo   infrastructure.MachineRepository
    period time.Duration
}

func NewMachinesService(repo infrastructure.MachineRepository, period time.Duration) services.Manager {
    return &MachinesService{repo: repo, period: period}
}

func (m *MachinesService) Save(ctx context.Context, machine *models.Machine) error {
    return m.repo.Put(ctx, machine)
}

func (m *MachinesService) Machines(ctx context.Context) ([]*models.Machine, error) {
    return m.repo.All(ctx)
}
