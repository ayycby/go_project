package main

import (
	"fmt"
	"go_project/datafile"
	"go_project/mystruct"
	"go_project/mytype"
	"log"
)

//4
func main() {
	//getVotes()
	//testStruct()
	//testType()
	getDate()
}

func getDate() {
	date := mytype.Date{}
	date.SetYear(2021)
	date.SetMonth(9)
	date.SetDay(5)
	fmt.Println(date)

	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}

func testType() {
	number := mytype.Number(2)
	fmt.Println(number)
	(&number).Double()
	fmt.Println(number)
	//此处无需&number.Double()，编译器会自动将非指针类型转换为指针类型
	//前提条件时存在number这个变量定义
	number.Double()
	fmt.Println(number)
}

func testStruct() {
	myCar := mystruct.Car{Description: "Audi Q5", Price: 199, TopSpeed: 299}
	fmt.Println(myCar)
	mystruct.ChangeCarPrice(&myCar, 599)
	fmt.Println(myCar)
}

func getVotes() {
	lines, err := datafile.GetStrings("E:\\go_code\\src\\go_project\\count\\votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	//创建一个空map(初始化)
	voteMap := map[string]int{}
	for _, line := range lines {
		voteMap[line]++
	}

	//ok用来表示该键是否已被赋值过
	//value, ok := voteMap["b"]
	//_, ok := voteMap["b"]也可用来判断键b是否存在
	//fmt.Println(voteMap)

	//map遍历
	for key, value := range voteMap {
		//fmt.Println(key, ": ", value)
		fmt.Printf("Vote for %s : %d\n", key, value)
	}
}
