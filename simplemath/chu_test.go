//-*- coding: UTF-8 -*-
// Author: afei00123

package simplemath

import "testing"

func TestChu(t *testing.T)  {
	r := Chu(8, 2)
	if r != 4 {
		t.Errorf("Chu(8, 2) failed. Got %d, excepted 4.", r)
	}
}