// Code generated by "stringer -type=status"; DO NOT EDIT

package unit

import "fmt"

const _status_name = "InactiveLoadingActiveExitedFailed"

var _status_index = [...]uint8{0, 8, 15, 21, 27, 33}

func (i status) String() string {
	if i < 0 || i >= status(len(_status_index)-1) {
		return fmt.Sprintf("status(%d)", i)
	}
	return _status_name[_status_index[i]:_status_index[i+1]]
}
