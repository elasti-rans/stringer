// Generated by: setup
// TypeWriter: stringer
// Directive: +test on Gap

package main

import (
	"fmt"
)

const (
	_Gap_name_0 = "gTwogThree"
	_Gap_name_1 = "gFivegSixgSevengEightgNine"
	_Gap_name_2 = "gEleven"
)

var (
	_Gap_index_0 = [...]uint8{0, 4, 10}
	_Gap_index_1 = [...]uint8{0, 5, 9, 15, 21, 26}
	_Gap_index_2 = [...]uint8{0, 7}
)

func (i Gap) String() string {
	switch {
	case 2 <= i && i <= 3:
		i -= 2
		return _Gap_name_0[_Gap_index_0[i]:_Gap_index_0[i+1]]
	case 5 <= i && i <= 9:
		i -= 5
		return _Gap_name_1[_Gap_index_1[i]:_Gap_index_1[i+1]]
	case i == 11:
		return _Gap_name_2
	default:
		return fmt.Sprintf("Gap(%d)", i)
	}
}
