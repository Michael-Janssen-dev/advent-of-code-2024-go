package core

func DP[T comparable](function func(T, func(T) int) int) func(T) int {
	dp := make(map[T]int)
	var inner func(param T) int
	inner = func(param T) int {
		cache, ok := dp[param]
		if ok {
			return cache
		}
		result := function(param, inner)
		dp[param] = result
		return result
	}
	return inner
}
