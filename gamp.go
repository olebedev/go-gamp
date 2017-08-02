package gamp

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/olebedev/go-gamp/client"
	"github.com/olebedev/go-gamp/client/gampops"
	"github.com/pkg/errors"
)

// New returns gamp client
func New(ctx context.Context, tid string) *gampops.Client {
	transport := httptransport.New("www.google-analytics.com", "/", []string{"https"})
	if ctx == nil {
		ctx = context.Background()
	}
	transport.Context = ctx
	transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		err := r.SetFormParam("tid", tid)
		if err != nil {
			return errors.Wrap(err, "set form param 'tid'")
		}
		return errors.Wrap(r.SetFormParam("v", "1"), "set form param 'v'")
	})
	transport.Consumers["image/gif"] = runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		_, err := ioutil.ReadAll(reader)
		return err
	})

	c := client.New(transport, nil)
	return c.Gampops
}
