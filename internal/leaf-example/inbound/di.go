package inbound

import (
	"github.com/enricodg/leaf-example/internal/leaf-example/inbound/http"
	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Http http.Inbound
	//Messaging messaging.Inbound
	//Worker    worker.Inbound
}
