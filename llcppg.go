package py

import (
	"github.com/goplus/lib/c"
)

//llgo:link (*Object).CStr C.PyUnicode_AsUTF8
func (self *Object) CStr() *c.Char { return nil }
