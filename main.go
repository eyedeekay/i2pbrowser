//go:generate go run -tags generate gen.go extensions.go

/*
Released under the The MIT License (MIT)
see ./LICENSE
*/

package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	. "github.com/eyedeekay/i2pbrowser/import"
)

func portable() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(path, "i2p")
}

func main() {
	chromium := flag.Bool("chromium", false, "use a chromium-based browser instead of a firefox-based browser.")
	chat := flag.Bool("chat", true, "Start an IRC client and configure it to use with I2P")
	rundir := flag.String("-i2p-profile", "", "override the normal profile directory")
	flag.Parse()

	err := os.Setenv("RHZ_PROFILE_OVERRIDE", portable())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(portable())
	Main(*chromium, *chat, *rundir, flag.Args())
}
