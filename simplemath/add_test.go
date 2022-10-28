//-*- coding: UTF-8 -*-
// Author: afei00123

package simplemath

import "testing"

func TestAdd(t *testing.T)  {
	r := Add(1, 9)
	if r != 11 {
		t.Errorf("Add(1, 9) failed. Got %d, excepted 10.", r)
	}
}