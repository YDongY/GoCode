package _7_benchmark

import "bytes"

func ConcatStringByAdd(elems []string) string {
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	return ret
}

func ConcatStringByBytesBuffer(elems []string) string {
	var buf bytes.Buffer
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	return buf.String()
}
