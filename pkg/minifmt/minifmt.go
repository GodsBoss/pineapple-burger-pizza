// Package minifmt provides some helpers that would usually be trivial with fmt, but compiling fmt
// into the binary makes its size siginificantly bigger (~400 KB).
package minifmt

import (
	"strconv"
	"strings"
)

func FormatInt(i int, size int, prefix string) string {
	s := strconv.Itoa(i)
	if len(s) >= size {
		return s
	}
	return strings.Repeat(prefix, size-len(s)) + s
}
