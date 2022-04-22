package model

import (
	leafTypes "github.com/paulusrobin/leaf-utilities/common/types"
	leafGormSoftDelete "github.com/paulusrobin/leaf-utilities/database/sql/integrations/gorm/plugin/softDelete"
	"gorm.io/plugin/optimisticlock"
	"time"
)

type (
	ID interface {
		uint64 | string
	}
	BaseAuditable[T ID] struct {
		ID        T                            `gorm:"primaryKey"`
		CreatedBy string                       `gorm:"column:created_by"`
		CreatedAt time.Time                    `gorm:"column:created_at"`
		UpdatedBy leafTypes.NullString         `gorm:"column:updated_by"`
		UpdatedAt time.Time                    `gorm:"column:updated_at"`
		DeletedBy leafTypes.NullString         `gorm:"column:deleted_by"`
		DeletedAt leafGormSoftDelete.DeletedAt `gorm:"column:deleted_at_unix;softDelete:deletedByField:DeletedBy"`
		Version   optimisticlock.Version       `gorm:"column:version"`
	}
)
