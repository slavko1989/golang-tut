package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func convert() {

	cV3 := "500000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

}
