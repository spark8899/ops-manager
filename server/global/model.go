package global

import (
	"time"

	"gorm.io/gorm"
)

type OPM_MODEL struct {
	ID        uint           `gorm:"primarykey"` // id
	CreatedAt time.Time      // create time
	UpdatedAt time.Time      // update time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // delete time
}
