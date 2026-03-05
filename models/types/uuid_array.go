package types

import "github.com/google/uuid"

type UUIDArray []uuid.UUID

func (a *UUIDArray) Scan(value interface{}) error {
	
}