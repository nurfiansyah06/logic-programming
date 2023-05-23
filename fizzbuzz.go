package logicprogramming

import "strconv"

func fizzBuzz() string {
	sumOfThree := 0
	sumOfFive := 0

	for i := 0; i < 100; i++ {
		if i % 5 == 0 {
			sumOfFive += i
			return "Buzz"
		}
		
		if i % 3 == 0 {
			sumOfThree += i
			return "Fizz"
		}

		return strconv.Itoa(i)
	}

	return ""
}