package cake

import (
	"fmt"
	"math/rand"
	"time"
)

type shop struct {
	Verbose bool

	Cakes int // number of cakes to bake

	BakeTime   time.Duration // time to bake one cake
	BakeStdDev time.Duration // standard deviation of baking time
	BakeBuf    int           // buffer slots between baking and icing

	NumIcers  int           // number of cooks doing icing
	IceTime   time.Duration // time to ice one cake
	IceStdDev time.Duration // standard deviation of icing time
	IceBuf    int           // buffer slots between icing and inscribing

	InscribeTime   time.Duration // time to inscribe one cake
	InscribeStdDev time.Duration // standard deviation o inscribing time
}

type cake int

func (s *shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		if s.Verbose {
			fmt.Println("baking ", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

func (s *shop) icer(iced chan<- cake, baked <-chan cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Println("icing ", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
	// close(iced)
}

func (s *shop) inscriber(iced <-chan cake) {
	// for c := range iced {
	// 	if s.Verbose {
	// 		fmt.Println("inscribing ", c)
	// 	}
	// 	work(s.InscribeTime, s.InscribeStdDev)
	// 	if s.Verbose {
	// 		fmt.Println("finished ", c)
	// 	}
	// }
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Println("inscribing", c)
		}
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("finished", c)
		}
	}
}

// work blocks the calling goroutine for a period of time that is normally distributed around d
// with a standard deviation of stddev
func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}

// Run runs the simulation `runs` times.
func (s *shop) Run(runs int) {
	for i := 0; i < runs; i++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for j := 0; j < s.NumIcers; j++ {
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}
