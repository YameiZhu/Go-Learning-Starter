package main

import "fmt"
func updateSlice(arr []int) {
	arr[0] = 100
}

func main()  {
		arr1 := [...]int{0,1,2,3,4,5,6,7}
		//slice
		s2 := arr1[2:]
		updateSlice(s2)
		fmt.Println("After updateSlice:(s2):")
		fmt.Println(s2)
		fmt.Println(arr1)
	
		//reslice
		fmt.Println("Reslice")
		fmt.Println(s2)
		s2 = s2[:5]
		fmt.Println(s2)
		s2 = s2[2:]
		fmt.Println(s2)
	
		//extend slice
		fmt.Println("Extending slice")
		arr := [...]int{0,1,2,3,4,5,6,7,8}
		fmt.Printf("arr = %v\n", arr)
		s1 := arr[2:6]
		s2 = s1[3:5]
		// slice can extend backforward, but not forward
		fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
		fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2))
		fmt.Println(s1[3:7])
	
		//append
		s3 := append(s2, 10)
		s4 := append(s3, 11)
		s5 := append(s4, 12)
		fmt.Println("s3, s4, s5: ", s3, s4, s5)
		// s4 and s5 no longer view the original array
		fmt.Println("arr: ", arr)

		// make/copy/delete
}