//-*- coding: UTF-8 -*-
// Author: afei00123

package simplemath

import "math"

func Sqrt(i int) int {
	v := math.Sqrt(float64(i))
	return int(v)
}