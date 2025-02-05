package dto

import "dockMon/internal/domain/models"

type Machines struct {
    List []*models.Machine `json:"machines"`
}
