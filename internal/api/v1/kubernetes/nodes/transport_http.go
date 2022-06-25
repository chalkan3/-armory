package nodes

import (
	"context"
	"encoding/json"
	"net/http"

	nodesv1beta1 "scheduler/pkg/manifest/kubernetes/v1beta1"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(svc Service, logger log.Logger) *mux.Router {
	opt := options(logger)

	createHandler := httptransport.NewServer(
		CreateEndpoint(svc),
		decodeCreate,
		encodeResponse,
		opt...,
	)

	patchHandler := httptransport.NewServer(
		PatchEndpoint(svc),
		decodePatch,
		encodeResponse,
		opt...,
	)

	listHandler := httptransport.NewServer(
		ListEndpoint(svc),
		decodeList,
		encodeResponse,
		opt...,
	)

	deleteHandler := httptransport.NewServer(
		RemoveEndpoint(svc),
		decodeDelete,
		encodeResponse,
		opt...,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodPost).Path(Base).Handler(createHandler)
	r.Methods(http.MethodPatch).Path(Base).Handler(patchHandler)
	r.Methods(http.MethodGet).Path(Base).Handler(listHandler)
	r.Methods(http.MethodDelete).Path(BaseWithID).Handler(deleteHandler)

	return r

}

func decodeCreate(ctx context.Context, r *http.Request) (interface{}, error) {
	var request CreateWorkerNodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodePatch(ctx context.Context, r *http.Request) (interface{}, error) {
	var request PatchWorkerNodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeDelete(ctx context.Context, r *http.Request) (interface{}, error) {
	request := RemoveNodeRequest{
		nodesv1beta1.KubeNodes{
			Metadata: &nodesv1beta1.Metadata{
				ID: pathParametersID(r),
			},
		},
	}
	return request, nil
}

func decodeList(ctx context.Context, r *http.Request) (interface{}, error) {

	return nil, nil
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

func codeFrom(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}

func options(logger log.Logger) []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}
}

func pathParametersID(r *http.Request) string {
	vars := mux.Vars(r)
	id := vars["id"]

	return id
}
