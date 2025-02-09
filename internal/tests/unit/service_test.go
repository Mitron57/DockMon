package unit

import (
    "context"
    "dockMon/internal/domain/models"
    "dockMon/internal/services"
    mockInfrastructure "dockMon/internal/tests/mocks"
    "go.uber.org/mock/gomock"
    "slices"
    "testing"
)

func TestNilSave(t *testing.T) {
    ctrl := gomock.NewController(t)
    repo := mockInfrastructure.NewMockMachineRepository(ctrl)
    ctx := context.Background()
    manager := services.NewMachinesService(repo)
    if manager.Save(ctx, nil) == nil {
        t.Fatal("Service saved nil value!")
    }
}

func TestSave(t *testing.T) {
    ctrl := gomock.NewController(t)
    repo := mockInfrastructure.NewMockMachineRepository(ctrl)
    ctx := context.Background()
    dummy := &models.Machine{}
    repo.EXPECT().Put(ctx, dummy).Return(nil)
    manager := services.NewMachinesService(repo)
    if err := manager.Save(ctx, dummy); err != nil {
        t.Fatal("Save error:", err)
    }
}

func TestMachines(t *testing.T) {
    ctrl := gomock.NewController(t)
    repo := mockInfrastructure.NewMockMachineRepository(ctrl)
    ctx := context.Background()
    dummyResult := []*models.Machine{{}}
    repo.EXPECT().All(ctx).Return(dummyResult, nil)
    manager := services.NewMachinesService(repo)
    machines, err := manager.Machines(ctx)
    if err != nil || !slices.Equal(machines, dummyResult) {
        t.Fatal("Get machines error:", err)
    }
}
