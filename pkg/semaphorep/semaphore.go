package semaphorep

import (
	"ajika/pkg/nullables"
	"database/sql"
	"time"
)

type Semaphore struct {
	ID           int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ConsumerGuid string
	Name         sql.NullString // not null because has no sense without value, we do not encourage garbage. NullString to prevent saving unset value as ""
	IsReleased   sql.NullBool   // not null with default true to have unambiguous value for selectors and on creating semaphore is released. NullBool to prevent saving unset value as false
	LastLockMsec uint
}

func NewSemaphore(name string) Semaphore {
	return Semaphore{
		Name:       nullables.NewNullString(name),
		IsReleased: nullables.NewNullBool(true),
	}
}
