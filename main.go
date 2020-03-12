package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	toEncodeValue := flag.String("encode", "", "String you want to encode")
	toDecodeValue := flag.String("decode", "", "String you want to decode")

	flag.Parse()

	if *toEncodeValue == "" && *toDecodeValue == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *toEncodeValue != "" {
		fmt.Println(url.QueryEscape(*toEncodeValue))
		os.Exit(1)
	}

	if *toDecodeValue != "" {
		decodedValue, err := url.QueryUnescape(*toDecodeValue)

		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(decodedValue)
		os.Exit(1)
	}
}
