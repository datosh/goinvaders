package main

import (
	"flag"
	_ "image/png"
	"log"
	"os"
	"runtime/pprof"
	"spaceinvaders"

	"github.com/hajimehoshi/ebiten"

	_ "spaceinvaders/statik"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(spaceinvaders.NewGame()); err != nil {
		log.Fatal(err)
	}
}
