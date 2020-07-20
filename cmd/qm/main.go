package main

import (
	"fmt"
	"log"
	"qm"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	log.SetFlags(0)
}

func main() {
	flags := parseFlags()
	defer func() { flags.File.Close() }()

	err := qm.EachJsonLine(flags.File, flags.QueryKey, flags.FingerprintKey, flags.SHA1, func(jl map[string]interface{}) {
		line, err := jsoniter.ConfigFastest.MarshalToString(jl)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(line)
	})

	if err != nil {
		log.Fatal(err)
	}
}
