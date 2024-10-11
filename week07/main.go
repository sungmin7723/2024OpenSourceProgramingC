package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// //오늘은 2024년 10월 11일 입니다.
	// var now time.Time = time.Now()
	// //var month int = int(now.Month()) //var month time.Month = now.Month()
	// //fmt.Println(month)
	// fmt.Printf("오늘은 %d년 %d월 %d일 입니다.\n", now.Year(), int(now.Month()), now.Day())
	// fmt.Printf("지금은 %d시 %d분 %d초 입니다.", now.Hour(), now.Minute(), now.Second())

	// var army string = "우리는 $가와 $민에 충성을 다하는 대한민$ 육군이다"
	// armyfixed := strings.NewReplacer("$", "국")
	// fmt.Println(armyfixed.Replace(army))

	in := bufio.NewReader(os.Stdin)
	fmt.Print("input your name: ")
	name, err := in.ReadString('\n')
	fmt.Println(name)
	fmt.Println(err)

}
