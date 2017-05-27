package main

/*
#include <stdio.h>
#include <stdint.h>
int ic = 5;
unsigned int uic =7;
int16_t is = 12345;
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

//http://www.toutiao.com/i6238899664402268674/
func main() {
	var ig int = 10
	igc := int(C.ic) //C int to Go int
	fmt.Println("value:", igc, ",type:", reflect.TypeOf(igc))
	//value: 5 ,type: int

	icg := C.int(ig) //Go int to C int
	fmt.Println("value:", icg, ",type:", reflect.TypeOf(icg))
	//value: 10 ,type: main._Ctype_int

	icp := (*C.int)(unsafe.Pointer(&ig)) //Go int pointer to C int pointer
	fmt.Println("value:", icp, ",type:", reflect.TypeOf(icp))
	//value: 0xc420010268 ,type: *main._Ctype_int

	i16t := int16(C.is) // C int16 to Go int16
	fmt.Println("value:", i16t, ",type:", reflect.TypeOf(i16t))
	//value: 12345 ,type: int16
}
