package model

import (
	leafModel "github.com/paulusrobin/leaf-utilities/common/model"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type (
	ServiceParameter struct {
		Variable    string `gorm:"column:variable"`
		Value       string `gorm:"column:value"`
		Description string `gorm:"column:description"`
		BaseAuditable[uint64]
	}

	ServiceParameterPagingResponse struct {
		leafModel.BasePagingResponse
		ServiceParameters []ServiceParameter
	}

	ServiceParameterPagingRequest struct {
		leafModel.PagingParams
	}
)

func (ServiceParameter) TableName() string {
	return "service_parameters"
}

func (s ServiceParameter) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *ServiceParameter) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}

type SoftDeleteDeleteClause struct {
	Field *schema.Field
}

func (sd SoftDeleteDeleteClause) Name() string {
	return ""
}

func (sd SoftDeleteDeleteClause) Build(clause.Builder) {
}

func (sd SoftDeleteDeleteClause) MergeClause(*clause.Clause) {
}
