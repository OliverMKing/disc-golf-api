package service

import (
	"context"
	"net/http"

	"github.com/olivermking/disc-golf-api/pkg/gen/openapi"
)

type service struct{}

var _ servicer = &service{}

func New() openapi.DiscApiServicer {
	return &adapter{s: &service{}}
}

var disc = openapi.Disc{
	Id:          1,
	Name:        "Teebird3",
	Distributor: "Innova",
	MaxWeight:   53,
	Diameter:    35,
	Height:      3,
	RimDepth:    5,
	Speed:       8,
	Glide:       5,
	Turn:        0,
	Fade:        2,
	PrimaryUse:  "Fairway",
}

func (s *service) GetDiscById(ctx context.Context, discId int64) (discRes, error) {
	return discRes{Code: http.StatusOK, Body: disc}, nil
}

func (s *service) ListDisc(ctx context.Context, skipToken int64, top int64) (discsRes, error) {
	return discsRes{Code: http.StatusOK, Body: []openapi.Disc{disc}}, nil
}
