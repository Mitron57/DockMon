package postgres

import (
    "context"
    "database/sql"
    "dockMon/internal/domain/interfaces/infrastructure"
    "dockMon/internal/domain/models"
    _ "github.com/lib/pq"
    "log"
)

const InsertStmt = `INSERT INTO HealthCheck VALUES ($1, $2, $3, $4) 
ON CONFLICT(IP) DO UPDATE SET PingTime = $2, Success = $3, LastSuccess = 
CASE WHEN $3 THEN $4 ELSE HealthCheck.LastSuccess END`

const SelectStmt = "SELECT IP, PingTime, Success, LastSuccess FROM HealthCheck"

type MachineRepository struct {
    insert   *sql.Stmt
    retrieve *sql.Stmt
    db       *sql.DB
}

func NewMachineRepository(db *sql.DB) infrastructure.MachineRepository {
    insert, err := db.Prepare(InsertStmt)
    if err != nil {
        log.Fatal(err)
    }
    retrieve, err := db.Prepare(SelectStmt)
    if err != nil {
        log.Fatal(err)
    }
    return MachineRepository{
        insert:   insert,
        retrieve: retrieve,
        db:       db,
    }
}

func (h MachineRepository) Put(ctx context.Context, machine *models.Machine) error {
    tx, err := h.db.Begin()
    if err != nil {
        return err
    }
    defer func() {
        if err != nil {
            tx.Rollback()
        }
    }()
    _, err = tx.StmtContext(ctx, h.insert).ExecContext(
        ctx,
        machine.IP, machine.PingTime, machine.Success, machine.LastSuccess,
    )
    if err != nil {
        return err
    }
    err = tx.Commit()
    if err != nil {
        return err
    }
    return nil
}

func (h MachineRepository) All(ctx context.Context) ([]*models.Machine, error) {
    rows, err := h.retrieve.QueryContext(ctx)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    result := make([]*models.Machine, 0, 5) //let cap be 5, to amortize reallocation frequency
    for rows.Next() {
        machine := models.Machine{}
        err = rows.Scan(&machine.IP, &machine.PingTime, &machine.Success, &machine.LastSuccess)
        if err != nil {
            return nil, err
        }
        result = append(result, &machine)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return result, nil
}
