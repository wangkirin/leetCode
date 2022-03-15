package main

import "strconv"

func main() {

}
func fizzBuzz(n int) []string {
	out := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			out[i-1] = "FizzBuzz"
		} else if i%3 == 0 && i%5 != 0 {
			out[i-1] = "Fizz"
		} else if i%3 != 0 && i%5 == 0 {
			out[i-1] = "Buzz"
		} else {
			out[i-1] = strconv.Itoa(i)
		}
	}
	return out
}
