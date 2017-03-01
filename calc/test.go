package main
import ("fmt"
	//"sync"
	//"runtime"
	//"strconv"
)
type tree struct{
	left *tree
	right *tree
	data int
}
func create(index int,value []int) (T *tree){
	T = &tree{}
	T.data = value[index - 1]
	if index < len(value) - 1 && 2*index <= len(value) && 2*index + 1 <= len(value){
		T.left = create(2*index,value)
		T.right = create(2*index+1,value)
	}
	return T
}
func showTree(T *tree){
	if T != nil{
		fmt.Println(T.data)
		if T.left != nil{
			showTree(T.left)
		}
		if T.right != nil{
			showTree(T.right)
		}
	}
}
func show(T *tree,level int){
	if level < 0 || T == nil{
		return
	}
	if level == 1{
		fmt.Println(T.data)
		return 
	}
	show(T.left,level-1)
	show(T.right,level-1)
}
func showAddTree(T *tree){
	level := getLevelTree(T)
	for i:=1;i<=level;i++{
		show(T,i)
	}
}
func getLevelTree(T *tree) (i int){
	s := &tree{}
	s = T.left
	i = 2
	for ;;i++{
		if s.left == nil{
			break
		}
		s = s.left
	}
	fmt.Println("level is")
	fmt.Println(i)
	fmt.Println("level is")
	return
}
func main(){
	value := []int{1,2,3,4,5,6,7,8,9}
	for i :=10;i<=100000;i++{
		value = append(value,i)
	}
	T := create(1,value)
	showTree(T)
	fmt.Println("________________________________________________________")
	showAddTree(T)
}
