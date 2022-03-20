package handlers

import (
	"context"
	"net/http"

	"PortalClient/pkg/tracing"

	"github.com/opentracing/opentracing-go"
)

type ServerResp struct {
	span opentracing.Span
	ctx  context.Context
	err  error
}

func ErrorHandler(f func(http.ResponseWriter, *http.Request, *ServerResp)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Start a span for request
		span := tracing.StartSpanFromRequest(r)
		ctx := opentracing.ContextWithSpan(r.Context(), span)

		tracing.LogURL(span, r.URL.Path, r.Method)

		// Inorder to capture endTime of request, we must call finish()
		defer span.Finish()

		resp := ServerResp{span: span, ctx: ctx}
		f(w, r, &resp)

		if resp.err != nil {
			tracing.LogError(span, resp.err)
			http.Error(w, resp.err.Error(), http.StatusInternalServerError)
		}
	}
}
