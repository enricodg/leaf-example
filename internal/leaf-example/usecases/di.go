package usecases

import (
	"go.uber.org/dig"
)

type UseCase struct {
	dig.In

	//BiddingUpload    biddingUpload.UseCase
}

func Register(container *dig.Container) error {
	//if err := container.Provide(shared.New); err != nil {
	//	return fmt.Errorf("[DI] error provide shared use case: %+v", err)
	//}
	return nil
}
