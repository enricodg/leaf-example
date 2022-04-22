package forwarder

import (
	"context"
	"github.com/enricodg/leaf-example/internal/outbound"
	"github.com/enricodg/leaf-example/internal/usecases/shared"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
)

const (
	tag = `[ForwardMessageUseCase]`
)

type (
	UseCase interface {
		ForwardMessage(ctx context.Context, message string) error
	}
	ucForwardMessage struct {
		outbound.Outbound
		shared   shared.UseCase
		resource pkgResource.Resource
	}
)

func New(
	outbound outbound.Outbound,
	shared shared.UseCase,
	resource pkgResource.Resource,
) UseCase {
	return &ucForwardMessage{
		Outbound: outbound,
		shared:   shared,
		resource: resource,
	}
}
