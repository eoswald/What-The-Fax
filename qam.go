package qam

import ()

type QAM interface {
	ModulateByte(uint8) []float64
	ModulateByteSlice([]uint8) []float64
}
