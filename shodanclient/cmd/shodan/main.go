package main

import (
	"fmt"
	"gophersize/shodanclient/shodan"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("USAGE: shodan searchterm")
	}
	apikey := "1a4RzgPSZAMHhDDna1uCwHsVGKJu5Wl4"
	s := shodan.New(apikey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\n Name: %s\n",
		info.Credits, info.DisplayName,
	)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
