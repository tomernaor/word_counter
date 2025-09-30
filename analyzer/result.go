package analyzer

type Result struct {
	Statistics statistics `json:"statistics"`
	TopNWords  []TopNWord `json:"top_n_words"`
}

type statistics struct {
	TotalEssay  int               `json:"total_essay"`
	Lexicon     statisticsLexicon `json:"lexicon"`
	PoolFetches []PoolFetch       `json:"pool_fetch"`
	AnalyzeTime string            `json:"analyze_time"`
}

type statisticsLexicon struct {
	TotalValid   int `json:"total_valid"`
	TotalInvalid int `json:"total_invalid"`
}

type PoolFetch struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}

type TopNWord struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}
