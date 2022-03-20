package restApi

import (
	"PortalClient/pkg/tracing"
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var client *http.Client

type ApiInfo struct {
	Url           string
	OperationName string
	Body          []byte
	Response      chan []byte
	Ctx           context.Context
	Err           chan error
}

func init() {
	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 5 * time.Second,
		},
	}
}

func Get(info ApiInfo) {
	req, err := http.NewRequest("GET", info.Url, bytes.NewReader(info.Body))
	info.Err <- errors.Wrap(nil, "test")
	return

	if err != nil {
		info.Err <- err
		info.Response <- nil
		return
	}
	//req.Header.Add("key", "value")
	//info.Ping(req)
	span, _ := opentracing.StartSpanFromContext(info.Ctx, info.OperationName)
	defer span.Finish()
	// defer close(info.Response)
	// defer close(info.Err)

	if err := tracing.Inject(span, req); err != nil {
		log.Warn("Error during trace inject: ", err)
	}

	// Send request and read data
	resp, err := client.Do(req)
	if err != nil {
		info.Err <- err
		info.Response <- nil
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		info.Err <- err
		info.Response <- nil
		return
	}

	//if resp.StatusCode != 200 {
	//	return "", fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	//}
	info.Response <- body
	info.Err <- nil
}

// Ping sends a ping request to the given hostPort, ensuring a new span is created
// for the downstream call, and associating the span to the parent span, if available
// in the provided context.
func (info *ApiInfo) Ping(req *http.Request) {
	span, _ := opentracing.StartSpanFromContext(info.Ctx, info.OperationName)
	defer span.Finish()
	// defer close(info.Response)
	// defer close(info.Err)

	if err := tracing.Inject(span, req); err != nil {
		log.Warn("Error during trace inject: ", err)
	}

	// Send request and read data
	resp, err := client.Do(req)
	if err != nil {
		info.Err <- err
		info.Response <- nil
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		info.Err <- err
		info.Response <- nil
		return
	}

	//if resp.StatusCode != 200 {
	//	return "", fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	//}
	info.Response <- body
	info.Err <- nil
}
