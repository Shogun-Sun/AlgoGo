package sorting

type SortOrder string
type AlgorithmType string

const (
	Ascending  SortOrder = "asc"
	Descending SortOrder = "desc"
)

const (
	StupidAlgorithm             AlgorithmType = "stupid"
	BubbleAlgorithm             AlgorithmType = "bubble"
	BubbleOptimizationAlgorithm AlgorithmType = "bubble-optimization"
)
