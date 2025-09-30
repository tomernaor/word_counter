package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof" // <-- this registers /debug/pprof handlers automatically
	"word_counter/analyzer"
)

func main() {
	log.Println("service is up")

	ctx := context.Background()

	// parse flags
	LexiconFileName := flag.String("lexicon", "input_files/words.txt", "Path to the lexicon file")
	UrlsFileName := flag.String("urls", "input_files/endg-urls-small.txt", "Path to the URLs file")
	WorkerPoolSize := flag.Int("workers", 10, "Number of concurrent workers")
	rateLimit := flag.Int("rate", 10, "Rate limit (requests per second)")
	topN := flag.Int("top", 10, "Top N words to output")
	showStats := flag.Bool("stats", false, "Show extended statistics")
	flag.Parse() // parse command-line flags

	// validate workers arg
	if *WorkerPoolSize <= 0 {
		fmt.Println("Error: `workers` must be greater than zero")
		return
	}

	// validate rate arg
	if *rateLimit <= 0 {
		fmt.Println("Error: `rate` must be greater than zero")
		return
	}

	// validate top n arg
	if *topN <= 0 {
		fmt.Println(fmt.Sprintf("Error: `top` must be greater than zero"))
		return
	}

	// set analyzer config
	cfg := analyzer.Config{
		LexiconFileName: *LexiconFileName,
		UrlsFileName:    *UrlsFileName,
		WorkerPoolSize:  *WorkerPoolSize,
		RateLimit:       *rateLimit,
		TopN:            *topN,
	}

	// execute analyzer
	res, err := analyzer.Execute(ctx, cfg)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
		return
	}

	// print result
	var p any = res.TopNWords
	if *showStats {
		p = res
	}

	// marshal to pretty json
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
		return
	}

	// Write to stdout
	fmt.Println(string(b))

	log.Println("service is down")
}
