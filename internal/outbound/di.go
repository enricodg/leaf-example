package outbound

import (
	"fmt"
	"github.com/enricodg/leaf-example/internal/outbound/messaging"
	"github.com/enricodg/leaf-example/internal/outbound/repositories"
	"github.com/enricodg/leaf-example/internal/outbound/webservices"
	"go.uber.org/dig"
)

type Outbound struct {
	dig.In

	//Cache        cache.Cache
	Messaging    messaging.Messaging
	Repositories repositories.Repositories
	WebServices  webservices.WebServices
}

func Register(container *dig.Container) error {
	//if err := cache.Register(container); err != nil {
	//	return fmt.Errorf("[DI] error provide outbound cache: %+v", err)
	//}
	if err := messaging.Register(container); err != nil {
		return fmt.Errorf("[DI] error provide outbound messaging: %+v", err)
	}
	if err := repositories.Register(container); err != nil {
		return fmt.Errorf("[DI] error provide outbound repositories: %+v", err)
	}
	if err := webservices.Register(container); err != nil {
		return fmt.Errorf("[DI] error provide outbound webservice: %+v", err)
	}
	return nil
}
