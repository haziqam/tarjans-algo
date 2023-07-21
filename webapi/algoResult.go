package webapi

type algoResult struct {
	Result          [][]string `json:"result"`
	TimeNanoseconds int        `json:"timeNanoseconds"`
}