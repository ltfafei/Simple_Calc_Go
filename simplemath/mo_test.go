//-*- coding: UTF-8 -*-
// Author: afei00123

package simplemath

import "testing"

func TestMo(t *testing.T)  {
	r := Mo(8, 5)
	if r != 3 {
		t.Errorf("Mo(8, 5) failed. Got %d, excepted 3.", r)
	}
}