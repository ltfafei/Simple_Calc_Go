//-*- coding: UTF-8 -*-
// Author: afei00123

package main

import (
	"Simple_Calc/simplemath"
	"fmt"
	"os"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values. \n\tsqrt\tAquare root of a non-negative value. \n\tchu\tdo the division. \n\tcheng\tdo multiplication. \n\tjian\tdo the subtraction. \n\tmo\tdo the modulo operation.")
}

func main() {
	//自动推导类型声明变量
	args := os.Args
	//nil默认值为零
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1> <integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1> <integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result is ", ret)

	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: calc add <integer>")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calc add <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result is: ", ret)
	case "chu":
		if len(args) != 4 {
			fmt.Println("USAGE: calc chu <integer1> <integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc chu <integer1> <integer2>")
			return
		}
		ret := simplemath.Chu(v1, v2)
		fmt.Println("Result is ", ret)
	case "cheng":
		if len(args) != 4 {
			fmt.Println("USAGE: calc cheng <integer1> <integer2>")
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc cheng <integer1> <integer2>")
			return
		}
		ret := simplemath.Cheng(v1, v2)
		fmt.Println("Result is ", ret)
	case "jian":
		if len(args) != 4 {
			fmt.Println("USAGE: calc jian <integer1> <integer2>")
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc jian <integer1> <integer2>")
			return
		}
		if v1 >= v2 {
			ret := simplemath.Jian(v1, v2)
			fmt.Println("Result is ", ret)
		} else {
			fmt.Println("The subtraction must be greater than the subtraction")
		}
	case "mo":
		if len(args) != 4 {
			fmt.Println("USAGE: calc mo <integer1> <integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc mo <integer1> <integer2>")
			return
		}
		ret := simplemath.Mo(v1, v2)
		fmt.Println("Result is ", ret)
	default:
		fmt.Println("程序运行错误，请检查！")
		Usage()
	}
}