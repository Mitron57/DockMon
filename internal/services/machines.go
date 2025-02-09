package services

import (
    "context"
    "dockMon/internal/domain/interfaces/infrastructure"
    "dockMon/internal/domain/interfaces/services"
    "dockMon/internal/domain/models"
    "errors"
)

var NilMachine = errors.New("machine is nil")

type MachinesService struct {
    repo infrastructure.MachineRepository
}

func NewMachinesService(repo infrastructure.MachineRepository) services.Manager {
    return &MachinesService{repo: repo}
}

func (m *MachinesService) Save(ctx context.Context, machine *models.Machine) error {
    if machine == nil {
        return NilMachine
    }
    return m.repo.Put(ctx, machine)
}

func (m *MachinesService) Machines(ctx context.Context) ([]*models.Machine, error) {
    return m.repo.All(ctx)
}
