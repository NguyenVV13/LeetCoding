package main

func main() {
	println(Multiple3And5(10))
}

func Multiple3And5(number int) int {
	var sum int = 0
	if number%3 == 0 || number%5 == 0 {
		sum += number
	}
	return sum
}
