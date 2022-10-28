//-*- coding: UTF-8 -*-
// Author: afei00123

package simplemath

import "testing"

func TestCheng(t *testing.T)  {
	r := Cheng(2, 9)
	if r != 18 {
		t.Errorf("Cheng(2, 9) failed. Got %d, excepted 18.", r)
	}
}