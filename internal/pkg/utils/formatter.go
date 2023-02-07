package utils

import (
	"github.com/samber/lo"
	"gorm.io/datatypes"
	"time"
)

func GormDatePtToTimePt(input *datatypes.Date) (output *time.Time) {
	if input != nil {
		output = lo.ToPtr(time.Time(*input))
	}

	return
}
