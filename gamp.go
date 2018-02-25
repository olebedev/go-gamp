package gamp

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/olebedev/go-gamp/client"
	"github.com/olebedev/go-gamp/client/gampops"
	"golang.org/x/time/rate"
)

// Add wrapper for API throttling
type wrapper struct {
	tid string
	*rate.Limiter
	*httptransport.Runtime
	context.Context
}

type stid interface {
	SetTid(string)
}

// Submit wraps httpclient.Submit for throttling
func (w *wrapper) Submit(op *runtime.ClientOperation) (interface{}, error) {
	if op.Context != nil {
		w.Wait(op.Context)
	} else if w.Context != nil {
		w.Wait(w.Context)
	} else if w.Runtime.Context != nil {
		w.Wait(w.Runtime.Context)
	} else {
		w.Wait(context.Background())
	}

	if w.tid != "" {
		if st, ok := op.Params.(stid); ok {
			st.SetTid(w.tid)
		}
	}

	return w.Runtime.Submit(op)
}

// New returns gamp client
func New(ctx context.Context, tid string) *gampops.Client {
	transport := httptransport.New("www.google-analytics.com", "/", []string{"https"})
	if ctx == nil {
		ctx = context.Background()
	}
	transport.Context = ctx
	transport.Consumers["image/gif"] = runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		_, err := ioutil.ReadAll(reader)
		return err
	})

	c := client.New(&wrapper{
		tid:     tid,
		Limiter: rate.NewLimiter(30, 1),
		Runtime: transport,
		Context: ctx,
	}, nil)
	return c.Gampops
}
