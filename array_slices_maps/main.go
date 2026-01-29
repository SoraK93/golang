package main

import "fmt"

func main() {
	arrayInt()
	arrayString()
	sliceArray()
	sliceAppend()
	sliceCopy()
	mapFound()
	mapDelete()
	mapReassign()
}

func arrayInt() {
	arr := [10]int{10, 20, 30, 50}
	fmt.Println(arr)
	fmt.Println(len(arr))
}

func arrayString() {
	arr := [5][2]string{{"a"}, {"b"}, {"c"}}
	for index1, element1 := range arr {
		for index2, element2 := range element1 {
			fmt.Printf("Array value at index [%d][%d] is [%q]\n", index1, index2, element2)
		}
	}
}

func sliceArray() {
	var arr [5]int = [5]int{10, 20, 90, 70, 60}
	slice := arr[:3]
	slice[2] = 900

	fmt.Println(arr)
	fmt.Println(slice)
}

func sliceAppend() {
	arr := [5]int{10, 20, 90, 70, 60}
	slice := append(arr[:2], arr[3:]...)
	fmt.Println(slice)
}

func sliceCopy() {
	arr := []int{10, 20, 90, 70, 60}
	slice := make([]int, 10)
	copy(slice, arr)
	slice[1] = 1000
	fmt.Println(arr)
	fmt.Println(slice)
}

func mapFound() {
	ascii_codes := map[string]int{}
	ascii_codes["A"] = 65
	_, found := ascii_codes["B"]
	if found {
		fmt.Println("key B was not found")
	}
}

func mapDelete() {
	new_map := make(map[string]int)
	new_map["A"] = 65
	new_map["F"] = 70
	new_map["K"] = 75
	delete(new_map, "F")
	fmt.Println(new_map)
}

func mapReassign() {
	ascii_codes := make(map[string]int, 10)
	ascii_codes["A"] = 65
	ascii_codes["F"] = 70
	ascii_codes["K"] = 75
	fmt.Println(len(ascii_codes))
	ascii_codes = make(map[string]int)
	ascii_codes["U"] = 85
	fmt.Println(len(ascii_codes))
}

