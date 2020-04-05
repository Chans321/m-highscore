package main

import (
	"context"
	"flag"
	"time"

	pbhighscore "github.com/Chans321/m-apis/m-highscore/v1"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:50051", "address to connect")
	flag.Parse()
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	con, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial establish connection with m-highscore server")
	}
	c := pbhighscore.NewGameClient(con)
	if c == nil {
		log.Info().Msg("Client Nil")
	}
	r, err := c.GetHighScore(timeoutCtx, &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to fetch highscore from m-highscore server")
	}
	if r != nil {
		log.Info().Interface("highscore", r.GetHighScore()).Msg("Highscore from m-highscore microservice")
	} else {
		log.Fatal().Err(err).Msg("Failed to fetch highscore from m-highscore server")
	}
	defer func() {
		err := con.Close()
		if err != nil {
			log.Info().Msg("Failed to close connection")
		}
	}()
}
