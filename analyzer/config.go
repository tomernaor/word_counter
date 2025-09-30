package analyzer

const DefaultWorkerPoolSize = 10
const DefaultRateLimit = 10
const DefaultTopN = 50

type Config struct {
	LexiconFileName string
	UrlsFileName    string
	WorkerPoolSize  int
	RateLimit       int
	TopN            int
}

func (c Config) GetWorkerPoolSize() int {
	if c.WorkerPoolSize > 0 {
		return c.WorkerPoolSize
	}

	return DefaultWorkerPoolSize
}

func (c Config) GetRateLimit() int {
	if c.RateLimit > 0 {
		return c.RateLimit
	}

	return DefaultRateLimit
}

func (c Config) GetTopN() int {
	if c.TopN > 0 {
		return c.TopN
	}

	return DefaultTopN
}
