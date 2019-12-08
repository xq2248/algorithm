/*
 *   Copyright (c) 2019
 *   All rights reserved.
 */

package container

import "testing"

func TestWater(t *testing.T) {
	w := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	r := maxArea(w)
	if r != 49 {
		t.Fail()
	}
}
