package main
import "strconv"

func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
			result = LCM(result, integers[i])
	}

	return result
}

func All[T any](ts []T, pred func(T) bool) bool {
    for _, t := range ts {
        if !pred(t) {
            return false
        }
    }
    return true
}

func first(n int, _ error) int {
	return n
}
func getIntFromString(param string) int {
	return first(strconv.Atoi(param))
}
