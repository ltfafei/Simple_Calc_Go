//-*- coding: UTF-8 -*-
// Author: afei00123

package simplemath

import "fmt"

func Jian(a int, b int) int {
	if a >= b {
		return  a - b
	} else {
		fmt.Println("The subtraction must be greater than the subtraction")
	}
	return  a - b
}