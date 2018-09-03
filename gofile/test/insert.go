 package main

import (
            "fmt"
)
// func main(){
// 	fmt.Println(124)
// 	slicetest()
// }
func Slicetest(){
	var sli1 = []byte {'1','2'}
	sli1 = append(sli1,'1','2','3')

	fmt.Println(sli1)
	sliss:= []int{1,2,3}
	sliss = append(sliss,2,3)
	fmt.Println(sliss)

}
// func main() {
//         a := insertSort(1, 2)
//         fmt.Println(a)
//         result, ok := mysqrt(1121)
//         if ok {
//                 fmt.Println(result)
//         }
//         fmt.Println(IsNumPosi(17.0))
//         season, err := Season(3)
//         if err {
//                 fmt.Println("not in range")
//                 return
//         }
//         fmt.Println(season)
//         itera(7)
//         iteraStr("This is a very import question!")
//         gotointera(9)
//         printIttera(10)
//         printGoIttera(10)
//         fmt.Println(MultiPly3Num(1,2,3))
//         fmt.Println(add(1,2))
//         fmt.Println(add_3(1,2))
//         fmt.Println(sub(1,2))
//         fmt.Println(sub_3(1,2))
//         fmt.Println(multi(1,2))
//         fmt.Println(multi_3(1,2))
//         fmt.Println(mysqrts_3(123))
// }





