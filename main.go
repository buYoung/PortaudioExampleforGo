package main

import (
	"github.com/gordonklaus/portaudio"
	"log"
)

type Record struct {
	samplerate float64
	delay      float64
	data       []float32
}

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()
	recordinit := &Record{samplerate: 44100, delay: 0.1, data: make([]float32, 44100*0.1)}

	recordstream, err := portaudio.OpenDefaultStream(1, 0,
		float64(recordinit.samplerate), len(recordinit.data), func(in []float32) {
			for i := range recordinit.data {
				recordinit.data[i] = in[i]
			}
		})
	if err != nil {
		log.Println(err)
	}
	recordstream.Start()

	plystream, err := portaudio.OpenDefaultStream(0, 1,
		float64(recordinit.samplerate), len(recordinit.data), func(out []float32) {
			for i := range recordinit.data {
				out[i] = recordinit.data[i]
			}
		})
	if err != nil {
		log.Println(err)
	}
	plystream.Start()

	select {}
}
