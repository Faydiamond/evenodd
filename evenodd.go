package evenodd

func EvenOdd(number1 int) string {
	X := number1 % 2
	var result string
	if X == 0 {
		result = " El numero es par "
	} else {
		result = " El numero es impar"
	}

	return result

}
