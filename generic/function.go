package generic

func gMap[T1 any, T2 any](arr []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func gReduce[T1 any, T2 any](arr []T1, init T2, f func(T2, T1) T2) T2 {
	result := init
	for _, v := range arr {
		result = f(result, v)
	}
	return result
}

func gFilter[T any](arr []T, in bool, f func(T) bool) []T {
	result := []T{}
	for _, v := range arr {
		choose := f(v)
		if (in && choose) || (!in && !choose) {
			result = append(result, v)
		}
	}
	return result
}

func gFilterIn[T any](arr []T, f func(T) bool) []T {
	return gFilter(arr, true, f)
}
func gFilterOut[T any](arr []T, f func(T) bool) []T {
	return gFilter(arr, false, f)
}

func gCountIf[T any](arr []T, f func(T) bool) int {
	cnt := 0
	for _, v := range arr {
		if f(v) {
			cnt++
		}
	}
	return cnt
}

type Sumable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func gSum[T any, U Sumable](arr []T, f func(T) U) U {
	var sum U
	for _, v := range arr {
		sum += f(v)
	}
	return sum
}
