package postgres

import (
    "context"
    "database/sql"
    "dockMon/internal/domain/interfaces/infrastructure"
    "dockMon/internal/domain/models"
    _ "github.com/lib/pq"
)

type MachineRepository struct {
    db *sql.DB
}

func NewMachineRepository(db *sql.DB) infrastructure.MachineRepository {
    return MachineRepository{db: db}
}

func (h MachineRepository) Put(ctx context.Context, health *models.Machine) error {
    tx, err := h.db.Begin()
    if err != nil {
        return err
    }
    stmt, err := tx.PrepareContext(
        ctx, `INSERT INTO HealthCheck VALUES ($1, $2, $3) ON CONFLICT(ip) DO UPDATE SET pingtime = $2, lastcheck = $3`,
    )
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.ExecContext(
        ctx,
        health.IP, health.PingTime, health.LastCheck,
    )
    if err != nil {
        tx.Rollback()
        return err
    }
    err = tx.Commit()
    if err != nil {
        return err
    }
    return nil
}

func (h MachineRepository) All(ctx context.Context) (map[string]*models.Machine, error) {
    tx, err := h.db.Begin()
    if err != nil {
        return nil, err
    }
    stmt, err := tx.PrepareContext(ctx, `SELECT ip, pingtime, lastcheck FROM HealthCheck`)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    rows, err := stmt.QueryContext(ctx)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    result := make(map[string]*models.Machine, 5) //let cap be 5, to amortize reallocation frequency
    for rows.Next() {
        machine := models.Machine{}
        err = rows.Scan(&machine.IP, &machine.PingTime, &machine.LastCheck)
        if err != nil {
            return nil, err
        }
        result[machine.IP] = &machine
    }
    return result, nil
}
