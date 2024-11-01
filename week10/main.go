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

	count := 0

	j := 2
	for j < n {
		if n%j == 0 {
			count++
		}
		j++
	}
	if count == 0 {
		fmt.Printf("%d is prime number", n)
	} else {
		fmt.Printf("%d is Not prime number", n)

	}
}
