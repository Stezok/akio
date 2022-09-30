package main

import (
	"NFTProject/internal/generator/clrchanger"
	"NFTProject/internal/meta"
	"image/png"
	"log"
	"os"
)

func main() {

	file, err := os.Open("tmp/body.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	proxy := clrchanger.NewColorProxy(img)

	cnf, err := os.Open("tmp/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer cnf.Close()

	spec, err := meta.LoadColorSpec(cnf)
	if err != nil {
		log.Fatal(err)
	}

	err = proxy.SetSpec(&spec)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("tmp/output.png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(out, proxy)
	if err != nil {
		log.Fatal(err)
	}
}
