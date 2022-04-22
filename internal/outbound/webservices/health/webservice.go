package health

import (
	"context"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
)

const (
	tag = `[HealthWebService]`
)

type (
	WebService interface {
		CheckHealth(ctx context.Context) error
	}
	wsHealth struct {
		resource pkgResource.Resource
	}
)

func New(resource pkgResource.Resource) WebService {
	return &wsHealth{resource: resource}
}
