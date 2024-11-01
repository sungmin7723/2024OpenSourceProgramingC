package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Input Number : ")
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i) // 줄바꿈등 제거. 파이썬의 strip 함수와 비슷
	n, _ := strconv.Atoi(i)  // 10진 정수형(32bit)으로 변환
	if err != nil {
		log.Fatal(err)
	}
	// count := 0
	var isPrime bool = true // 가독성 up, 메모리 down
	// bug fix 1 == prime number
	if n <= 1 { // A prime numbers is a natural number greater than 1 that has only 1 and itself as divisors
		isPrime = false
	} else {
		j := 2
		for j < n {
			if n%j == 0 {
				// count++
				isPrime = false // 더하기 연산 제거
				break           // performance up
			}
			fmt.Printf("%d ", j) // check j loop
			j++
		}
	}

	// if count == 0 {
	if isPrime { // == 비교 연산자 제거
		fmt.Printf("%d is prime number", n)
	} else {
		fmt.Printf("%d is Not prime number", n)
	}
}
