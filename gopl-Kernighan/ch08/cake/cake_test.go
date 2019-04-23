package cake

import (
	"testing"
	"time"
)

var defaults = shop{
	Verbose:      testing.Verbose(),
	Cakes:        20,
	BakeTime:     10 * time.Millisecond,
	NumIcers:     1,
	IceTime:      10 * time.Millisecond,
	InscribeTime: 10 * time.Millisecond,
}

func Benchmark(b *testing.B) {
	cakeshop := defaults
	cakeshop.Run(b.N)
}

func BenchmarkBuffers(b *testing.B) {
	// Adding buffers has no effect
	cakeshop := defaults
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Run(b.N)
}

func BenchmarkVariable(b *testing.B) {
	// Adding variability to rate of each step
	// increases total time due to channel delays.
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.Run(b.N)
}

func BenchmarkVariableBuffers(b *testing.B) {
	// Adding channel buffers reduces
	// delays resulting from variability.
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Run(b.N)
}

func BenchmarkSlowIcing(b *testing.B) {
	// Making the middle stage slower
	// adds directly to the critical path.
	cakeshop := defaults
	cakeshop.IceTime = 500 * time.Millisecond
	cakeshop.Run(b.N)
}

func BenchmarkSlowIcingManyIcers(b *testing.B) {
	// Adding more icing cooks reduces the cost of icing
	// to its sequential component, following Amdahl's Law.
	cakeshop := defaults
	// cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.NumIcers = 5
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Run(b.N)
}
