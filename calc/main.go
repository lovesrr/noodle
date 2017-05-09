package main
import(
	"fmt"
	"reflect"
)

func main(){
	var a int
	a = -1
	var b rune = 234
	var c interface{} = 'c'
	fmt.Println("%b,%d",a>>4,a<<4)
	fmt.Println("%v,%v,%v",reflect.TypeOf(a),reflect.TypeOf(b),reflect.TypeOf(c))
	if s,ok := c.(rune);ok{
		fmt.Printf("%c",s)
	}else{
		fmt.Println("error")
	}
}
