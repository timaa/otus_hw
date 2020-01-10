package linkedlist

import (
	"fmt"
	"strings"
)

func (l List) ToString() string {
	curElem := l.first
	str := strings.Builder{}

	for curElem != nil {
		str.WriteString(fmt.Sprintf("%v \n", curElem.Value()))
		curElem = curElem.Next()
	}
	return str.String()
}
