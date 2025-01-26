package main

import "fmt"

type Streamer interface {
	Stream()
}

type NetflixStreamer struct{}

func (n *NetflixStreamer) Stream() {
	fmt.Println("streaming from netlix server")
}

type AppleTvStreamer struct{}

func (a *AppleTvStreamer) Stream() {
	fmt.Println("streaming from apple tv server")
}

func main() {
	netflix := &NetflixStreamer{}
	netflix.Stream()

	appltTv := &AppleTvStreamer{}
	appltTv.Stream()
}
