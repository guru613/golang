package main

import "fmt"

func isprime(n int) bool {

	count := 0

	if n < 2 {
		return false
	}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count == 2

}

func main() {
	var num int
	fmt.Println("enter the number")
	fmt.Scan(&num)

	if isprime(num) {
		fmt.Println("it is prime number ")
	} else {
		fmt.Println("it is not prime")
	}
}
