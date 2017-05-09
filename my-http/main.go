package main
import (
	"net/http"
	"fmt"
	"log"
	"strings"
	"strconv"
	"noodle/my-http/logic"
)
var s int = k + 1
func init(){
	fmt.Println(s)
	fmt.Println(k)
}
func exec(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	path := r.Form["path"][0]
	data := r.Form["data"][0]
	rs := []int{}
	input := strings.Split(data,",")
	inputs := make([]int,len(input),len(input))
	for i:=0;i<len(input);i++{
		inputs[i],_ = strconv.Atoi(input[i])
	}
	if path == "maopao"{
		fmt.Println(inputs)
		rs = logic.Maopao(inputs)
	}
	if path == "kuaipai"{
		fmt.Println(inputs)
		logic.Kuaipai(inputs,0,len(inputs)-1)
		rs = inputs
	}
	fmt.Fprintln(w,"result is ",rs)
}
func main(){
	http.HandleFunc("/say",exec)
	err := http.ListenAndServe(":9998",nil)
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}
var k int = 2
