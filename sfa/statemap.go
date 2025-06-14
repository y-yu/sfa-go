package sfa

import (
	"github.com/y-yu/sfa-go/common"
)

type StateMap map[common.State]common.State

func (lhs StateMap) Equal(rhs StateMap) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for k, rv := range rhs {
		if lv, ok := lhs[k]; !ok || lv != rv {
			return false
		}
	}

	return true
}
