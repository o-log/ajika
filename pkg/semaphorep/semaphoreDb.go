package semaphorep

import (
	"ajika/pkg/wraperr"

	"gorm.io/gorm"
)

func CreateSemaphoreIfNotExist(db *gorm.DB, semaphoreName string) error {
	semaphores := []Semaphore{}

	// use Find instead of First because First writes error to log if object not found, also it's cleaner to check array size than gorm.ErrRecordNotFound error
	if result := db.Where("name = ?", semaphoreName).Find(&semaphores); result.Error != nil {
		// die if can not select semaphores - can't operate without them
		return wraperr.Wrap(result.Error)
	}

	if len(semaphores) > 0 { // semaphore exists
		return nil
	}

	semaphore := NewSemaphore(semaphoreName)
	if result := db.Create(&semaphore); result.Error != nil {
		// die if can not select semaphores - can't operate without them
		return wraperr.Wrap(result.Error)
	}

	return nil
}

func DeleteAll(db *gorm.DB) error {
	if result := db.Where("1 = 1").Delete(&Semaphore{}); result.Error != nil {
		return wraperr.Wrap(result.Error)
	}

	return nil
}
