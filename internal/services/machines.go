package services

import (
    "context"
    "dockMon/internal/domain/interfaces/infrastructure"
    "dockMon/internal/domain/interfaces/services"
    "dockMon/internal/domain/models"
    "maps"
    "slices"
    "sync"
    "time"
)

const future = 10 * time.Second

type MachinesService struct {
    repo  infrastructure.MachineRepository
    mx    sync.Mutex
    cache map[string]*models.Machine
    last  time.Time
}

func NewMachinesService(repo infrastructure.MachineRepository) services.Manager {
    return &MachinesService{repo: repo}
}

func (m *MachinesService) Save(ctx context.Context, machine *models.Machine) error {
    if time.Now().Before(m.last.Add(future)) {
        m.mx.Lock()
        m.cache[machine.IP] = machine
        m.mx.Unlock()
    }
    return m.repo.Put(ctx, machine)
}

func (m *MachinesService) Machines(ctx context.Context) ([]*models.Machine, error) {
    if time.Now().After(m.last.Add(future)) {
        m.last = time.Now()
        cache, err := m.repo.All(ctx)
        if err != nil {
            return nil, err
        }
        m.mx.Lock()
        defer m.mx.Unlock()
        m.cache = cache
    }
    return slices.Collect(maps.Values(m.cache)), nil
}
