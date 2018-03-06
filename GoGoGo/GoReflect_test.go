/*
author: ZengJJ
date: 2018/3/6  
*/
package GoGoGo

import (
	"reflect"
	"testing"
	"fmt"
)

// Go Reflect : https://studygolang.com/articles/1251
// 利用reflect设置struct的字段 : http://blog.csdn.net/pkueecser/article/details/50422533

type Foo struct {
	X string
	Y int
}

func (f Foo) Do(i int) string {
	fmt.Printf("X is: %s, Y is: %d", f.X, f.Y)
	return ""
}

func TestTypeOf(t *testing.T) {
	var s = "abc"
	fmt.Println(reflect.TypeOf(s).String()) //string
	fmt.Println(reflect.TypeOf(s).Name())   //string
	fmt.Println()

	var f Foo
	typ := reflect.TypeOf(f)
	fmt.Println(typ.String()) //GoGoGo.Foo
	fmt.Println(typ.Name())   //Foo ，返回结构体的名字
	fmt.Println(typ.Kind())   //struct
	fmt.Println()

	var f2 = &Foo{}
	typ2 := reflect.TypeOf(f2)
	fmt.Println(typ2)        //*GoGoGo.Foo
	fmt.Println(typ2.Kind()) //ptr
	fmt.Println()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fmt.Printf("%s type is :%s\n", field.Name, field.Type)
	}
	//x type is :string
	//y type is :int
	field2, _ := typ.FieldByName("X") //等价于typ.Field(0)，返回的也是StructField对象
	fmt.Println(field2.Name)          // X
}

func TestMethod(t *testing.T) {
	var f Foo
	typ := reflect.TypeOf(f)
	fmt.Println(typ.NumMethod()) //1， Foo 方法的个数  (exported methods)
	m := typ.Method(0)
	fmt.Println(m.Name) //Do
	fmt.Println(m.Type) //func(main.Foo, int) string
	fmt.Println()

	var i = 1
	var foo = &Foo{"abc", 123}
	reflect.ValueOf(foo).MethodByName("Do").Call([]reflect.Value{reflect.ValueOf(i)})
}

func TestValueOf(t *testing.T) {
	var i = 123
	var f = Foo{"abc", 123}
	var s = "abc"
	fmt.Println(reflect.ValueOf(i).String())    //<int Value>
	fmt.Println(reflect.ValueOf(f).Interface()) //<main.Foo Value>
	fmt.Println(reflect.ValueOf(s))             //abc

	val := reflect.ValueOf(f)
	fmt.Println(val.FieldByName("Y")) //<int Value>  interface.Value对象
	//
	typ := reflect.TypeOf(f)
	fmt.Println(typ.FieldByName("Y")) //{  <nil>  0 [] false} false StructField对象
}

func TestCanSet(t *testing.T) {
	var s = "abc"
	fv := reflect.ValueOf(s)
	fmt.Println(fv.CanSet()) //false
	// fv.SetString("edf")   //panic
	fv2 := reflect.ValueOf(&s)
	fmt.Println(fv2.CanSet()) //false
	// fv2.SetString("edf")      //panic
	fmt.Println()

	var f = Foo{"abc", 123}
	val := reflect.Indirect(reflect.New(reflect.TypeOf(f))) // true
	fmt.Println(val.FieldByName("Y").CanSet())
	val = reflect.ValueOf(&f).Elem() // true
	fmt.Println(val.FieldByName("Y").CanSet())
	val = reflect.New(reflect.TypeOf(f)).Elem() // true
	fmt.Println(val.FieldByName("Y").CanSet())
}

//=======================================other test
func say(text string) {
	fmt.Println(text)
}

func Call(fun interface{}, params ...interface{}) (result []reflect.Value) {
	f := reflect.ValueOf(fun)
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func TestOther(t *testing.T) {
	Call(say, "hello")
	fmt.Println()

	var qqq = []int{1, 2, 3}
	fmt.Println(reflect.TypeOf(qqq).Elem())
}
