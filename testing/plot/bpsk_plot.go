package main

import (
	"github.com/eoswald/What-The-Fax"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

func main() {
	bpsk := qam.NewBPSKBuilder(1, 1000, 44100)
	waveform := bpsk.ModulateByteSlice([]uint8{0xAA})
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"Modulated Signal", WaveformToPoints(waveform))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func WaveformToPoints(waveform []float64) plotter.XYs {
	pts := make(plotter.XYs, len(waveform))
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = waveform[i]
	}
	return pts
}
