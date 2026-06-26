package interfaceSize

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Data struct {
	ID    int64
	Name  string
	Flag  bool
	Value float64
}

func MeasureSize(data interface{}) {
	fmt.Println("Размер структуры:", reflect.TypeOf(data).Size()) // например, 40
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		fmt.Println("Размер структуры через указатель:", t.Elem().Size())
	}
	fmt.Printf("Размер через unsafe %d", unsafe.Sizeof(data))
}
