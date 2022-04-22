package hello

import (
	"context"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
)

const (
	tag = `[HelloMessaging]`
)

type (
	Messaging interface {
		PublishMessage(ctx context.Context, message string) error
	}
	mHello struct {
		resource pkgResource.Resource
	}
)

func New(resource pkgResource.Resource) Messaging {
	return &mHello{
		resource: resource,
	}
}
