package main

import (
	"flag"

	grpcsetup "github.com/Chans321/m-highscore/internal/server/grpc"
)

func main() {
	var addressPtr = flag.String("address", ":50051", "address to connest m-highscore service")
	flag.Parse()
	s := grpcsetup.NewServer(*addressPtr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start grpc serverfor m-highscore services")
	}
}
