package calc

import "errors"

func Add(numbers ...int) (error, int) {
	if len(numbers) < 2 {
		return errors.New("Accepts at least 2 args"), 0
	}
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}
	return nil, sum
}
