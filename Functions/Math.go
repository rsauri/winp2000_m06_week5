package functions

func Add(iNum1, iNum2 int) int {
	return iNum1 + iNum2
}

func Divide(iNum1, iNum2 int) (int, int) {
	return iNum1 / iNum2, iNum1 % iNum2
}
