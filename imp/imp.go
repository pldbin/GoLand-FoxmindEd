package imp

import "fmt"

func Imp() {
	arr := [5]int{5, 2, 14, 21, 20}
	min, max, minsum, maxsum := MaxMinSum(arr)
	//fmt.Printf("minimum: %d, maximum: %d.", min, max)
	fmt.Printf("minimumsum: %d, maximumsum: %d.", minsum-max, maxsum-min)
}
func MaxMinSum(arr [5]int) (min, max, minsum, maxsum int) {
	mi := arr[0]
	//miindex := 0
	misum := 0
	//fmt.Println("miindex", miindex, "mi", mi)
	ma := arr[0]
	//maindex := 0
	masum := 0
	//fmt.Println("miindex", maindex, "ma", ma)
	//for ind, numb := range arr {
	for _, numb := range arr {
		if numb < mi {
			mi = numb
			//miindex = ind
		}
		misum += numb
		if numb > ma {
			ma = numb
			//maindex = ind
		}
		masum += numb
		//fmt.Println("ind for:", ind, "numb for:", numb)
		//fmt.Println("miindex:", miindex, ", ", "mi - мінімальна змінна:", mi)
		//fmt.Println("maindex:", maindex, ", ", "ma - максимальна змінна:", ma)
		//fmt.Println("misum:", misum, "masum:", masum)
	}
	return mi, ma, misum, masum
}
