package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	var res []string
	keyArr := keyArr(input)
	sort.Ints(keyArr)
	for _, v := range keyArr {
		res = append(res, input[v])
	}
	return res
}

func keyArr(input map[int]string) []int {
	var res []int
	for i := range input {
		res = append(res, i)
	}
	return res
}

// func main() {
// 	input := map[int]string{
// 		0:  "M",
// 		1:  "A",
// 		2:  "Y",
// 		3:  " ",
// 		4:  "T",
// 		5:  "H",
// 		6:  "E",
// 		7:  " ",
// 		8:  "F",
// 		9:  "O",
// 		10: "R",
// 		11: "C",
// 		12: "E",
// 		13: " ",
// 		14: "B",
// 		15: "E",
// 		16: " ",
// 		17: "W",
// 		18: "I",
// 		19: "T",
// 		20: "H",
// 		21: " ",
// 		22: "Y",
// 		23: "O",
// 		24: "U",
// 	}

// 	fmt.Println(sortMapValues(input))
// }
