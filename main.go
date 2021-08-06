package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isPrime(n int) {
	isPrime := true
	if n == 0 || n == 1 {
		fmt.Printf(" %d is not a  prime number\n\n", n)
	} else {
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				fmt.Printf(" %d is not a  prime number\n\n", n)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Printf(" %d is a prime number\n\n", n)
		}
	}
}

func isExist(n int, arr []int)  {
	for i := 0; i < len(arr); i++ {
		if n == arr[i] {
			fmt.Printf(" %d is exists in the file\n\n", arr[i])
		} else {
			fmt.Printf(" %d is not exists in the file\n\n", arr[i])
		}
	}
}

func main()  {
	arr := []int{13,2}
	text, err := ioutil.ReadFile("text.txt")
	check(err)
	str := string(text)
	n, _:= strconv.Atoi(str)
	fmt.Println(str, reflect.TypeOf(str))
	fmt.Println(n, reflect.TypeOf(n))
	isPrime(n)
	isExist(n, arr)
}