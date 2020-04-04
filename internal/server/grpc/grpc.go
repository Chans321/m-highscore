package grpc

import(
	pbhighscore "github.com/Chans321/m-apis/m-highscore/v1"
	"google.golang.org/grpc"
	"context"
	"github.com/rs/zerolog/log"
)

type Grpc struct{
	address string
	srv *.grpc.Server
}

var HighScore=9999.00


func (g *Grpc)SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse, error){
	log.Info().Msg("Setting High Score in m-highscore")
	HighScore=input.HighScore
	return &pbhighscore.SetHighScoreResponse{
		Set:true
	},nil

}

func (g *Grpc)GetHighScore(ctx context.Context, input *pbhighscore.GetHighScoreRequest) (*pbhighscore.GetHighScoreResponse, error){
	log.Info().Msg("Getting High Score from m-highscore")
	return &pbhighscore.GetHighScoreResponse{
		HighScore:HighScore
	},nil

}