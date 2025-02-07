package postgres

import (
	"context"
	"database/sql"
	"dockMon/internal/domain/interfaces/infrastructure"
	"dockMon/internal/domain/models"
	_ "github.com/lib/pq"
)

const InsertStmt = `INSERT INTO HealthCheck VALUES ($1, $2, $3, $4) 
ON CONFLICT(IP) DO UPDATE SET PingTime = $2, Success = $3, LastSuccess = 
CASE WHEN $3 THEN $4 ELSE HealthCheck.LastSuccess END`

const SelectStmt = "SELECT IP, PingTime, Success, LastSuccess FROM HealthCheck"

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
	stmt, err := tx.PrepareContext(ctx, InsertStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(
		ctx,
		health.IP, health.PingTime, health.Success, health.LastSuccess,
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
	stmt, err := tx.PrepareContext(ctx, SelectStmt)
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
		err = rows.Scan(&machine.IP, &machine.PingTime, &machine.Success, &machine.LastSuccess)
		if err != nil {
			return nil, err
		}
		result[machine.IP] = &machine
	}
	return result, nil
}
