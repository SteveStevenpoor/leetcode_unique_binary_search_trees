package numTrees

func numTrees(n int) int {
	uniqueTrees := make([]int, n+1)

	uniqueTrees[0] = 1
	uniqueTrees[1] = 1

	for i := 2; i < n+1; i++ {
		for j := 1; j <= i; j++ {
			uniqueTrees[i] += uniqueTrees[j-1] * uniqueTrees[i-j]
		}
	}

	return uniqueTrees[n]
}
