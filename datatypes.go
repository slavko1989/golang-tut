package data

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println

func data() {

	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(25.2))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))

}
