package domain

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	Id            uuid.UUID
	Name          string
	Amount        string
	TransactionAt time.Time
}
