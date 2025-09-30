package analyzer

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
	"word_counter/progress_bar"
	"word_counter/worker_pool"

	"golang.org/x/time/rate"
)

type FetchResult struct {
	WordsCount *Counter[string]
	PoolCount  *Counter[int]
}

func Execute(ctx context.Context, cfg Config) (*Result, error) {
	// measure start time
	start := time.Now()

	// read lexicon
	lexicon, err := newLexiconFromFile(cfg.LexiconFileName)
	if err != nil {
		return nil, err
	}

	// read essays
	essays, err := newEssayFromFile(cfg.UrlsFileName)
	if err != nil {
		return nil, err
	}

	// fetch essays
	fetchResults, err := fetch(ctx, cfg, lexicon, essays)
	if err != nil {
		return nil, err
	}

	// get top n words
	poolFetched, err := CountPoolFetched(ctx, fetchResults.PoolCount)
	if err != nil {
		return nil, err
	}

	// get top n words
	topNWords, err := CountTopNWords(ctx, fetchResults.WordsCount, cfg.TopN)
	if err != nil {
		return nil, err
	}

	// compute elapsed time
	elapsed := time.Since(start)
	h := int(elapsed.Hours())
	m := int(elapsed.Minutes()) % 60
	s := int(elapsed.Seconds()) % 60
	elapsedFormatted := fmt.Sprintf("%02d:%02d:%02d", h, m, s)

	// create a result
	result := Result{
		Statistics: statistics{
			TotalEssay: essays.total(),
			Lexicon: statisticsLexicon{
				TotalValid:   lexicon.totalValid(),
				TotalInvalid: lexicon.totalInValid(),
			},
			PoolFetches: poolFetched,
			AnalyzeTime: elapsedFormatted,
		},
		TopNWords: topNWords,
	}

	return &result, nil
}

func fetch(ctx context.Context, cfg Config, lexicon *lexicon, essays *essay) (*FetchResult, error) {
	// create progress bars
	bars := progress_bar.NewProgressBars(int64(essays.total()))

	// create counter for number of fetched url per pool
	poolFetchCounter := NewCounter[int]()

	// create counter for essay words
	wordsCounter := NewCounter[string]()

	// create a worker pool to fetch urls
	poolSize := cfg.GetWorkerPoolSize()
	pool := worker_pool.NewPool(ctx, poolSize)

	// create a rate limiter
	rateLimit := cfg.GetRateLimit()
	limiter := rate.NewLimiter(rate.Limit(rateLimit), poolSize)

	// fetch essay urls
	for _, url := range essays.urls {
		// wait until rate limit token is available
		if err := limiter.Wait(ctx); err != nil {
			return nil, err
		}

		// send a job to the worker pool
		pool.RunJob(
			func(ctx context.Context, poolId int) error {
				// increment the pool fetch counter
				poolFetchCounter.Increment(poolId)

				// request the essay page
				page, err := requestUrl(ctx, url)
				if err != nil {
					bars.IncError()
					return err
				}

				// split the page into words
				words := strings.Fields(page)
				for _, word := range words {
					// normalize word
					word = normalize(word)

					if lexicon.exists(word) {
						wordsCounter.Increment(word)
					}
				}

				return nil
			},
		)

		// raise the progress bar by one
		bars.IncProgress()
	}

	// wait until all jobs are done
	pool.Close()

	// wait until progress bars are done
	bars.Wait()

	return &FetchResult{
		WordsCount: wordsCounter,
		PoolCount:  poolFetchCounter,
	}, nil
}

func requestUrl(_ context.Context, url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.New("status code " + strconv.Itoa(res.StatusCode))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CountTopNWords(_ context.Context, wordsCounter *Counter[string], topN int) ([]TopNWord, error) {
	words := make([]TopNWord, 0, wordsCounter.Len())

	// iterate over the words counter and build a slice
	wordsCounter.ForEach(func(key string, count int) {
		words = append(words, TopNWord{
			Word:  key,
			Count: count,
		})
	})

	// sort the words slice by count in descending order
	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})

	// handle case when the word slice is smaller than the required top n
	if topN > len(words) {
		topN = len(words)
	}

	return words[:topN], nil
}

func CountPoolFetched(_ context.Context, poolFetchCounter *Counter[int]) ([]PoolFetch, error) {
	poolFetches := make([]PoolFetch, 0, poolFetchCounter.Len())

	// iterate over the pool fetch counter and build a slice
	poolFetchCounter.ForEach(func(key int, count int) {
		poolFetches = append(poolFetches, PoolFetch{
			Id:    key,
			Count: count,
		})
	})

	// sort the pool fetches slice by id in ascending order
	sort.Slice(poolFetches, func(i, j int) bool {
		return poolFetches[i].Id < poolFetches[j].Id
	})

	return poolFetches, nil
}
