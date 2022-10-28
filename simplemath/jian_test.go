//-*- coding: UTF-8 -*-
// Author: afei00123

package simplemath

import "testing"

func TestJian(t *testing.T)  {
	r := Jian(8, 2)
	if r != 6 {
		t.Errorf("Chu(8, 2) failed. Got %d, excepted 6.", r)
	}
}