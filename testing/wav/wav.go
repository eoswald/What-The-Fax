package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/cryptix/wav"
	"github.com/eoswald/What-The-Fax"
)

const (
	bits = 32
	rate = 44100
)

func main() {
  //bs := make([]uint8, 1000, 10000)
	bpsk := qam.NewBPSKBuilder(1, 1000, rate)
	waveform := bpsk.ModulateByteSlice([]uint8{0xAB})

	wavOut, err := os.Create("Test.wav")
	checkErr(err)
	defer wavOut.Close()

	meta := wav.File{
		Channels:        1,
		SampleRate:      rate,
		SignificantBits: bits,
	}

	writer, err := meta.NewWriter(wavOut)
	checkErr(err)
	defer writer.Close()

	start := time.Now()

	for i := 0; i < len(waveform); i++ {
		y := int32(0.8 * math.Pow(2, bits-1) * waveform[i])

		err = writer.WriteInt32(y)
		checkErr(err)
	}

	fmt.Printf("Simulation Done. Took:%v\n", time.Since(start))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
