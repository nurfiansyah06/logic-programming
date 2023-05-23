package logicprogramming

func isPalindrome(x int) bool {
	reverse :=0
	originalNumber := x
	if x < 0 {
		return false
	}

	for x !=0 {
		reverse =(reverse*10)+(x%10)
		x/=10
	}

	return originalNumber==reverse
}