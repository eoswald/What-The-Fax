package qam

import (
	"math"
)

type BPSKBuilder struct {
	Fc         float64
	Amplitude  float64
	SampleRate int
	Sin0       []float64
	Sin1       []float64
}

func NewBPSKBuilder(amplitude float64, fc float64, samplerate int) *BPSKBuilder {
	points0 := make([]float64, int(float64(samplerate)/fc))
	for i := 0; i < int(float64(samplerate)/fc); i++ {
    t := float64(i)/float64(samplerate)
		points0[i] = amplitude * math.Sin(2*math.Pi*fc*t+math.Pi)
	}

	points1 := make([]float64, int(float64(samplerate)/fc))
	for i := 0; i < int(float64(samplerate)/fc); i++ {
    t := float64(i)/float64(samplerate)
		points1[i] = amplitude * math.Sin(2*math.Pi*fc*t)
	}

	b := BPSKBuilder{
		Amplitude:  amplitude,
		Fc:         fc,
		SampleRate: samplerate,
		Sin0:       points0,
		Sin1:       points1,
	}
	return &b
}

func (bpsk *BPSKBuilder) ModulateByte(b uint8) []float64 {
	var mod_byte []float64
	for i := 0; i < 8; i++ {
		n := b & 1
		if n == 0 {
			mod_byte = append(mod_byte, bpsk.Sin0...)
		} else {
			mod_byte = append(mod_byte, bpsk.Sin1...)
		}
		b = b >> 1
	}
	return mod_byte
}

func (bpsk *BPSKBuilder) ModulateByteSlice(bs []uint8) []float64 {
	var mod_bs []float64
	for i := 0; i < len(bs); i++ {
		mod_bs = append(mod_bs, bpsk.ModulateByte(bs[i])...)
	}
	return mod_bs
}
