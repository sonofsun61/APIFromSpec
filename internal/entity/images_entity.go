package entity

import "github.com/google/uuid"

type Images struct {
	ID    uuid.UUID `db:"id"`
	Image []byte    `db:"image"`
}
