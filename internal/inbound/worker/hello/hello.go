package hello

import "context"

func (c *Controller) Run(ctx context.Context) error {
	return c.UseCase.Health.CheckHealth(ctx)
}
