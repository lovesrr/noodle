package logic
import(
	"fmt"
)
var Ss int = 20000
func init(){
	fmt.Println(Ss)
}
func Test(){
	fmt.Println("from logic.test")
}
func Maopao(inputs []int) (ret []int){
	fmt.Println("xxxx",inputs)
	input := inputs
	end := len(input) - 1
	index := 0
	for index=end;index>=0;index-- {
		for j:=1;j<=index;j++ {
			if input[j] < input[j-1] {
				input[j-1] = input[j] + input[j-1]
				input[j] = input[j-1] - input[j]
				input[j-1] = input[j-1] - input[j]
			}
		}
	}
	ret = input
	return
}
