package core

func DP[T comparable, Q any](function func(T, func(T) Q) Q) func(T) Q {
	dp := make(map[T]Q)
	var inner func(param T) Q
	inner = func(param T) Q {
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
