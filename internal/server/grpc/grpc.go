package grpc

import (
	"context"
	"net"

	pbhighscore "github.com/Chans321/m-apis/m-highscore/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

var HighScore = 9999.00

func (g *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse, error) {
	log.Info().Msg("Setting High Score in m-highscore")
	HighScore = input.HighScore
	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil

}

func (g *Grpc) GetHighScore(ctx context.Context, input *pbhighscore.GetHighScoreRequest) (*pbhighscore.GetHighScoreResponse, error) {
	log.Info().Msg("Getting High Score from m-highscore")
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil

}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port in m-highscore microoservice")
	}
	g.srv = grpc.NewServer()

	pbhighscore.RegisterGameServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("starting grpc service for m-highscore microservice")

	err = g.srv.Serve(lis)

	if err != nil {
		return errors.Wrap(err, "failed to start grpc server for m-highscore micoservice")
	}

	return nil

}
