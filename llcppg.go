package py

import (
	"github.com/goplus/lib/c"
)

// llgo:link (*Object).CStr C.PyUnicode_AsUTF8
func (u *Object) CStr() *c.Char { return nil }
