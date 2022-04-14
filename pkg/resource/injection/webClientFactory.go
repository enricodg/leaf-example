package injection

import (
	leafHeimdall "github.com/paulusrobin/leaf-utilities/webClient/integrations/heimdall"
	leafWebClient "github.com/paulusrobin/leaf-utilities/webClient/webClient"
)

func NewWebClientFactory() leafWebClient.Factory {
	return leafHeimdall.NewClientFactory()
}
