package model

import (
	obModel "github.com/enricodg/leaf-example/internal/outbound/model"
	leafModel "github.com/paulusrobin/leaf-utilities/common/model"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
)

type (
	ServiceParameterUpsertRequest struct {
		Variable    string
		Value       string
		Description string
	}

	ServiceParameterUpsertResponse struct {
		ID          uint64 `json:"id"`
		Variable    string `json:"variable"`
		Value       string `json:"value"`
		Description string `json:"description"`
		Version     int64  `json:"version"`
	}

	ServiceParameterPagingRequest struct {
		leafModel.PagingParams
	}

	ServiceParameterPagingResponse struct {
		leafModel.BasePagingResponse
		ServiceParameters []obModel.ServiceParameter `json:"serviceParameters"`
	}
)

func (s ServiceParameterUpsertResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *ServiceParameterUpsertResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}
