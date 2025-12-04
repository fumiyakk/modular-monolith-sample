package entity

import "github.com/google/uuid"

type Contract struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Status ContractStatus
}
