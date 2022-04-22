package health

import (
	"context"
)

func (u *ucHealth) CheckHealth(ctx context.Context) error {
	return u.Outbound.WebServices.Health.CheckHealth(ctx)
}
