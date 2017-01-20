package main

import (
    "fmt"
	"errors"
)

func Add(a int,b int)(ret int,err error){
    if a<0 || b<0 {
	    err=errors.New("Should be non-negative numbers")
		return
	}
	return a+b,nil
}

func myfunc2(args []int){
   for _,arg := range args{
      fmt.Println(arg)
   }
}

func myfunc3(args ...int){
   for _,arg := range args{
       fmt.Println(arg)
   }
}

func myfunc4(args ...int){
    myfunc3(args...)
	
	myfunc3(args[2:]...)
}

type Stringer interface{
    String() string
}

//当函数可以接受任意的对象实例时，声明为interface{}
func MyPrintf(args ...interface{}){
   for _, arg := range args{
       switch v := arg.(type){
	      case int:
		     fmt.Println(arg,"is an int value")
		  case string:
		     fmt.Println(arg,"is an string value")
		  case int64:
		     fmt.Println(arg,"is an int64 value")
		  default:
		     if v,ok := arg.(Stringer); ok{
			     val := v.String()
			 }else {
			     fmt.Println(arg,"is an unknow type")
			 }
	   }
   }
}

//为类型添加方法
type Integer int
func (a Integer) Less(b Integer) bool{
   return a<b
}
//用指针传递
func (a *Integer) Add(b Integer){
   *a += b
}

//定义接口
type LessAddr interface{
   Less(n Integer) bool
   Add(b Integer)
}

type Lesser interface{
   Less(b Integer) bool
}

type Rect struct{
   x,y float64
   width,height float64
}
//go语言没有构造函数的概念，对象的创建交由一个全局函数完成，以NewXXX命名
func NewRect(x,y,width,height float64) *Rect{
    return &Rect{x,y,width,height}
}

func main() {

fmt.Printf("HelloWorld!\n")
var i = 2
switch i {
  case 0:
     fmt.Printf("0")
  case 1:
     fmt.Printf("1")
  case 2:
     fallthrough//此关键字继续执行下一个case
  case 3:
     fmt.Printf("3")
  default:
     fmt.Printf("Default")
}
  fmt.Printf("\n")
  var Num = 2
  //switch后面的表达式可以省略
  switch{
    case 0 <= Num && Num <= 3:
       fmt.Printf("0-3\n")
    case 4 <= Num && Num <= 6:
       fmt.Printf("4-6\n")
  }
  
  ret,_ := Add(1,2)
  fmt.Println(ret)
  
  myfunc2([]int{1,2,3,4})
  fmt.Println("**********测试不定参数*******")
  myfunc4(1,2,3,4)
  fmt.Println("**********测试任意类型不定参数*******")
  var v1 int = 1
  var v2 int64 = 234
  var v3 string = "hello"
  var v4 float32 = 1.234
  MyPrintf(v1,v2,v3,v4)
  fmt.Println("**********测试闭包匿名函数*******")
  //f := func(x,y,int){
     //return x+y
  //}
  var j int = 5
  a := func()(func()){
      var i int = 10  //变量i只有内部的匿名函数才能访问
	  return func(){
	     fmt.Printf("i,j: %d,%d\n",i,j)
	  }
  }()//括号内可加参数列表
  a()
  j*=2
  a()
  fmt.Println("**********测试为类型添加方法*******")
  var aa Integer = 1
  if aa.Less(2){
     fmt.Println(aa,"小于 2")
  }
  aa.Add(2)
  fmt.Println("aa= ",aa)
  fmt.Println("**********测试引用类型*******")
  var c = [3]int{1,2,3}
  var d = &c
  d[1]++
  fmt.Println(c,*d)
  fmt.Println("**********初始化结构体类型对象实例*******")
  //rect1 := new(Rect)
  //rect2 := &Rect{}
  //rect3 := &Rect{0,0,100,200}
  //rect4 := &Rect{width:100,height:200}
  fmt.Println("**********将对象实例赋值给接口*******")
  //var t Integer = 1
  //var tt LessAddr = &t//要求对象实例实现了接口要求的所有方法
  
  //var b1 Lesser = t 和 var b1 Lesser = &t 均可以
}
