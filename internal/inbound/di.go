package inbound

import (
	"github.com/enricodg/leaf-example/internal/inbound/http"
	"github.com/enricodg/leaf-example/internal/inbound/messaging"
	"github.com/enricodg/leaf-example/internal/inbound/worker"
	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Http      http.Inbound
	Messaging messaging.Inbound
	Worker    worker.Inbound
}
