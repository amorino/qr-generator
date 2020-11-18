package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"strconv"

	"./rand"

	"github.com/divan/qrlogo"
)

func main() {
	file, err := os.Open("input.png")
	errcheck(err, "Failed to open logo:")
	defer file.Close()

	logo, _, err := image.Decode(file)
	errcheck(err, "Failed to decode PNG with logo:")

	for i := 1; i < 1000; i++ {
		text := rand.String(10)
		output := "output/" + strconv.Itoa(i) + "_" + text + ".png"

		qr, err := qrlogo.Encode(text, logo, 1024)
		errcheck(err, "Failed to encode QR:")

		out, err := os.Create(output)
		errcheck(err, "Failed to open output file:")
		out.Write(qr.Bytes())
		out.Close()
	}

	fmt.Println("Done! Written QR images")
}

func errcheck(err error, str string) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(1)
	}
}
