package health

import (
	"context"
	pkgError "github.com/enricodg/leaf-example/pkg/constant/error"
	pkgPath "github.com/enricodg/leaf-example/pkg/constant/path"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	leafWebClient "github.com/paulusrobin/leaf-utilities/webClient/webClient"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

func (w *wsHealth) CheckHealth(ctx context.Context) error {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	baseURI, err := url.Parse(w.resource.ConfigWebServiceExample.Host)
	if err != nil {
		return err
	}
	baseURI.Path = path.Join(baseURI.Path, pkgPath.HealthPath)
	webClient := w.resource.WebClientFactory.Create(
		leafWebClient.NewWebClientOptionWithRetry(
			w.resource.ConfigWebServiceExample.TimeoutDuration,
			w.resource.ConfigWebServiceExample.RetryCount)...)
	response, _ := webClient.Get(ctx, baseURI.String(), header, nil)

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		dataByte, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		var resp any
		if err := json.Unmarshal(dataByte, &resp); err != nil {
			return err
		}
		return nil
	default:
		return pkgError.ErrOutbound
	}
}
