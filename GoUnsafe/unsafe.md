### unsafe

#### 定义

##### Pointer
> unsafe.Pointer在golang中是各种指针相互转换的桥梁,允许随意的读写内存，所以是不安全的
>
> 独有的四个操作:
> * 任何类型的指针都考科一转换为Pointer
> * Pointer 可以转换为任何类型的指针
> * uintptr 类型可以转换为Pointer
> * Pointer 可以转换为 uintptr
>

##### uintptr
> uintptr是golang的内置类型，是能存储指针的值的整型，不是指针类型，所以
> * 当uintptr的值所指的数据被GC所移动，uintptr的值并不会同时被更改，所以uintptr 不能存储在变量中，在变量被使用之前可能地址已改变
> * uintptr 无法持有对象，当他所指的数据没有被其他指针所指，就可能被GC回收
>
> uintptr和unsafe.Pointer的区别就是：
> * unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
> * uintptr是用于指针运算的，GC 不把 uintptr 当指针
>

##### 方法
* `func Sizeof(x ArbitraryType) uintptr`
    > 整个数据类型占用的内存大小，比如slice类型，sizeof返回的是整个slice描述器的大小（data、len、cap）
* ```func Offsetof(x ArbitraryType) uintptr```
    > 函数的参数必须是一个字段 x.f, 然后返回 f 字段相对于 x 起始地址的偏移量, 包括可能的空洞.
* ```func Alignof(x ArbitraryType) uintptr```
    > unsafe.Alignof 函数返回对应参数的类型需要对齐的倍数

##### valid patterns (go vat)
* 某种指针类型转换成Pointer 再转换成另一种指针类型
    > 后一种类型大小不能大于前一种类型，并且要有相同的内存布局
* Pointer转换成uintptr，但不能转换回来，通常这个用法是打印该地址
* Pointer转换成uintptr，进行一些指针地址运算，再转换回来
    > 如 ```p = unsafe.Pointer(uintptr(p) + offset)```，注意 必须在同一表达式中完成，不能讲uintptr赋值给变量，再使用，因为在使用之前，地址可能变了
* Conversion of a Pointer to a uintptr when calling syscall.Syscall
* Conversion of the result of reflect.Value.Pointer or reflect.Value.UnsafeAddr from uintptr to Pointer.
* Conversion of a reflect.SliceHeader or reflect.StringHeader Data field to or from Pointer

#### 其他
##### 内存对齐
* 结构体成员的起始地址是自己Alignment的整倍数
* 整个结构体的Alignment是数据成员的Alignment中的最大值，整个结构体的起始地址也是自己Alignment的整倍数

##### go:linkname
* [golang进阶(八)——隐藏技能go:linkname](http://blog.csdn.net/lastsweetop/article/details/78830772)

#### 参考资料
* [Go语言 unsafe的妙用](https://studygolang.com/articles/1414)
* [Go unsafe包](https://studygolang.com/articles/9446)
