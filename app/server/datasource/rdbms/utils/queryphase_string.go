// Code generated by "stringer -type=QueryPhase"; DO NOT EDIT.

package utils

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[QueryPhaseUnspecified-0]
	_ = x[QueryPhaseDescribeTable-1]
	_ = x[QueryPhaseListSplits-2]
	_ = x[QueryPhaseReadSplits-3]
}

const _QueryPhase_name = "QueryPhaseUnspecifiedQueryPhaseDescribeTableQueryPhaseListSplitsQueryPhaseReadSplits"

var _QueryPhase_index = [...]uint8{0, 21, 44, 64, 84}

func (i QueryPhase) String() string {
	if i < 0 || i >= QueryPhase(len(_QueryPhase_index)-1) {
		return "QueryPhase(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _QueryPhase_name[_QueryPhase_index[i]:_QueryPhase_index[i+1]]
}
