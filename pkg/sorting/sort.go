package sorting

func Sort(arr []int, algorithm AlgorithmType, sortOrder SortOrder) []int {
	switch algorithm {
	case StupidAlgorithm:
		return StupidSort(arr, sortOrder)
	case BubbleAlgorithm:
		return BubbleSort(arr, sortOrder)
	case BubbleOptimizationAlgorithm:
		return BubbleOptimization(arr, sortOrder)
	default:
		return StupidSort(arr, sortOrder)
	}
}
