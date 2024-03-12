package semaphorep

import (
	"ajika/pkg/nullables"
	"ajika/pkg/wraperr"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Returned not-nil semaphore means semaphore is locked successfully.
func Lock(db *gorm.DB, semaphoreName string) (*Semaphore, error) {

	// TODO:
	// нужен будет контроль что семафор не завис в занятом состоянии: если семафор занят больше например минуты (читать из .env) - принудительно освобождать его (с записью ошибки в лог для алертинга)
	// принудительное освобождение сделать в этом же таймере, расписать
	// при этом в процессе проверки проверять - если он продолжается дольше например минуты (читать из .env) - прерывать ее и писать ошибку в лог для алертинга

	guid := uuid.NewString()

	result := db.Exec(
		`update semaphores
		set consumer_guid = ?, is_released = false, last_lock_msec = (CURRENT_TIMESTAMP(3) * 1000)
		where name = ? and is_released = true`,
		guid,
		semaphoreName,
	)
	if result.Error != nil {
		return nil, wraperr.Wrap(result.Error)
	}

	semaphores := []Semaphore{}

	if result := db.Where("consumer_guid = ?", guid).Find(&semaphores); result.Error != nil {
		// review such errors
		return nil, wraperr.Wrap(result.Error)
	}

	if len(semaphores) == 0 {
		return nil, nil
	}

	return &semaphores[0], nil
}

func Release(db *gorm.DB, semaphore *Semaphore) error {
	if semaphore.IsReleased.Bool { // for a case when semphore was released before call
		return nil
	}

	semaphore.IsReleased = nullables.NewNullBool(true)
	if result := db.Save(&semaphore); result.Error != nil {
		return wraperr.Wrap(result.Error)
	}

	return nil
}
