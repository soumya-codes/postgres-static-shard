// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Heartbeat struct {
	MachineID     pgtype.Text `json:"machine_id"`
	LastHeartbeat pgtype.Int8 `json:"last_heartbeat"`
}