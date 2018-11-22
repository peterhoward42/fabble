package main

import (
	"flag"
	"log"
	"os"

	"github.com/peterhoward42/fabble/ingest"
	"github.com/peterhoward42/fabble/ingest/archive"
	"github.com/peterhoward42/fabble/ingest/csv"
	"github.com/peterhoward42/fabble/ingest/storer"
)

// Single shot reading and processing of a tar file, and using gRPC to
// notify a storage service of each *User* encountered.
func main() {

	var tarName string
	flag.StringVar(&tarName, "tarname", "", "Specify a .tar filepath")
	flag.Parse()
	if tarName == "" {
		log.Fatal("You must provide a tar pathname")
	}

	tar, err := os.Open(tarName)
	if err != nil {
		log.Fatalf("Cannot open archive: %v", err)
	}
	defer tar.Close()

	// The Ingestor we use below is decoupled from where the inputs are coming
	// from or their format, and requires only something that satisfies
	// the UserParser interface.
	inputSource, err := archive.AccessCSVInTar(tar)
	if err != nil {
		log.Fatalf("archive.ReaderForCSVFile(): %v", err)
	}
	inputParser := csv.NewUserParser(inputSource)

	// Simiarly it takes an abstracted Storer.
	storer, err := storer.NewGRPCStorer()
	if err != nil {
		log.Fatalf("storer.NewGRPCStorer(): %v", err)
	}

	// Now can construct the ingestoer.
	ingestor := ingest.NewIngestor(inputParser, storer)

	numStored, err := ingestor.ParseInputAndStore()
	if err != nil {
		log.Fatalf("ingestor.ParseInputAndStore(): %v", err)
	}
	log.Printf("Ingested and stored %d users", numStored)
}
