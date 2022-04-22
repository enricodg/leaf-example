package inbound

import (
	"github.com/enricodg/leaf-example/internal/inbound/http"
	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Http http.Inbound
}
