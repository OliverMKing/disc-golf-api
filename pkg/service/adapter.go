package service

import (
	"context"

	"github.com/olivermking/disc-golf-api/pkg/gen/openapi"
)

type resp[B any] struct {
	Code int
	Body B
}

type discRes = resp[openapi.Disc]
type discsRes = resp[[]openapi.Disc]

type servicer interface {
	GetDiscById(context.Context, int64) (discRes, error)
	ListDisc(context.Context, int64, int64) (discsRes, error)
}

// adapter makes our disc service compatible with the openapi generated one
// we do this so our service can have typed responses instead of the openapi response
// body that is simply interface{} (untyped)
type adapter struct {
	s servicer
}

var _ openapi.DiscApiServicer = &adapter{}

func (a *adapter) GetDiscById(ctx context.Context, discId int64) (openapi.ImplResponse, error) {
	resp, err := a.s.GetDiscById(ctx, discId)
	return openapi.ImplResponse{Code: resp.Code, Body: resp.Body}, err
}

func (a *adapter) ListDisc(ctx context.Context, skipToken int64, top int64) (openapi.ImplResponse, error) {
	resp, err := a.s.ListDisc(ctx, skipToken, top)
	return openapi.ImplResponse{Code: resp.Code, Body: resp.Body}, err
}
