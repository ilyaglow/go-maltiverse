package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/ilyaglow/maltiverse"
)

var ip = flag.String("ip", "", "IPv4 to lookup in Maltiverse")

func main() {
	flag.Parse()
	if *ip == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	mc := maltiverse.NewAPIClient(maltiverse.NewConfiguration())
	report, _, err := mc.IPv4.GetIP(context.Background(), *ip)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(report)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(os.Stdout)
}
