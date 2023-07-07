package dto

import (
	"time"

	"github.com/google/uuid"
)

type InsertFileDTOModel struct {
	Filename  uuid.UUID
	Filetype  string
	Filesize  int
	Binary    []byte
	CreatedAt time.Time
}
