package program

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"log"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cravenceiling/indexer/cli/cmd/util"
	"github.com/cravenceiling/indexer/cli/internal/document"
	"github.com/cravenceiling/indexer/cli/internal/parser"
)

// Indexer
type Indexer struct {
	Parser parser.Parser
}

// Index indexes files in the given directory
func (in *Indexer) Index(dir string, req HttpRequest) {
	start := time.Now()
	var totalIndexed int64

	fileChan := make(chan string, 100)
	var wg sync.WaitGroup

	const numWorkers = 4
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			in.processFiles(fileChan, req, &totalIndexed)
		}()
	}

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Printf("error accessing file %s: %v", path, err)
			return nil
		}

		if !info.IsDir() {
			fileChan <- path
		}

		return nil
	})

	if err != nil {
		log.Printf("error walking directory: %v", err)
	}

	close(fileChan)
	wg.Wait()

	log.Printf(
		"process finished: %d files indexed in %.2f minutes\n",
		totalIndexed,
		time.Since(start).Minutes(),
	)
}

// processFiles processes files from the channel and uploads data in batches
func (in *Indexer) processFiles(
	fileChan <-chan string,
	req HttpRequest,
	totalIndexed *int64,
) {
	const batchSize = 500

	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	counter := 0

	flush := func() {
		if err := Upload(req, buf); err != nil {
			log.Printf("error uploading files: %v", err)
		} else {
			atomic.AddInt64(totalIndexed, int64(counter))
		}
		buf.Reset()
		counter = 0
	}

	for path := range fileChan {
		if isEmpty, err := util.CheckEmpty(path); err != nil {
			log.Printf("error checking file %s: %v", path, err)
			continue
		} else if isEmpty {
			log.Printf("skipping empty file: %s", path)
			continue
		}

		em, err := in.Parser.Parse(path)
		if err != nil {
			log.Printf("error parsing file %s: %v", path, err)
			continue
		}

		err = encoder.Encode(IndexAction{Index: IndexDocument{Index: req.Index}})
		if err != nil {
			log.Printf("error encondig index action")
		}

		err = encoder.Encode(document.Document{Path: path, Email: em})
		if err != nil {
			log.Printf("error encondig documents")
		}

		counter++

		if counter == batchSize {
			flush()
		}
	}

	// Upload any remaining data
	if counter > 0 {
		flush()
	}
}
