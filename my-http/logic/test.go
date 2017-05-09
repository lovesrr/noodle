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
func Kuaipai(inputs []int,start int,end int) {
	if end == start{
		return	
	}
	i,j := start,end
	key := inputs[start]
	for i < j{
		for j>i && inputs[j] >=key{
			j--
		}
		inputs[i] = inputs[j]
		fmt.Println(inputs)
		for i <j && inputs[i] <= key{
			i++
		}
		inputs[j] = inputs[i]
	}
	inputs[i] = key
	fmt.Printf("i is %d,j is %d",i,j)
	if i - start > 1{
		Kuaipai(inputs,start,i-1)
	}
	if end - j > 1{
		Kuaipai(inputs,j+1,end)
	}
	return
}
