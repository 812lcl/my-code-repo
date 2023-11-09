package structure

import "fmt"

func ArrayFunc() {
	var array1 [5]int
	var slice1 = array1[1:3]

	array1[0] = 0
	array1[1] = 5
	array1[2] = 2
	array1[3] = 3
	array1[4] = 4

	fmt.Printf("Value of a is %v\n", array1)
	fmt.Printf("Type of a is %T\n", array1)
	fmt.Printf("Address of array1[1] is %p\n", &array1[1])
	fmt.Println()

	fmt.Printf("Value of slice1 is %v\n", slice1)
	fmt.Printf("Type of slice1 is %T\n", slice1)
	fmt.Printf("Address of slice1 is %p\n", &slice1)
	fmt.Printf("Array address of slice1 is %p\n", &slice1[0])
	fmt.Println()

	array1[1] = 1
	fmt.Println("Change the value of the array, the slice will change too")
	fmt.Printf("Value of array1 is %v\n", array1)
	fmt.Printf("Length of array1 is %d\n", len(array1))
	fmt.Printf("Capacity of array1 is %d\n", cap(array1))
	fmt.Printf("Address of array1[1] is %p\n", &array1[1])

	fmt.Println()
	fmt.Printf("Value of slice1 is %v\n", slice1)
	fmt.Printf("Length of slice1 is %d\n", len(slice1))
	fmt.Printf("Capacity of slice1 is %d\n", cap(slice1))
	fmt.Printf("Address of slice1 is %p\n", &slice1)
	fmt.Printf("Array address of slice1 is %p\n", &slice1[0])

	slice1 = append(slice1, []int{3, 4, 5, 6}...)
	fmt.Println()
	fmt.Println("Append the slice, the slice refer to another new array")
	fmt.Printf("Value of array1 is %v\n", array1)
	fmt.Printf("Length of array1 is %d\n", len(array1))
	fmt.Printf("Capacity of array1 is %d\n", cap(array1))
	fmt.Printf("Address of array1[1] is %p\n", &array1[1])

	fmt.Println()
	fmt.Printf("Value of slice1 is %v\n", slice1)
	fmt.Printf("Length of slice1 is %d\n", len(slice1))
	fmt.Printf("Capacity of slice1 is %d\n", cap(slice1))
	fmt.Printf("Address of slice1 is %p\n", &slice1)
	fmt.Printf("Array address of slice1 is %p\n", &slice1[0])

	fmt.Println()
	array1[1] = 5
	fmt.Println("Change the value of the original array, the slice will NOT change")
	fmt.Printf("Value of array1 is %v\n", array1)
	fmt.Printf("Length of array1 is %d\n", len(array1))
	fmt.Printf("Capacity of array1 is %d\n", cap(array1))
	fmt.Printf("Address of array1[1] is %p\n", &array1[1])

	fmt.Println()
	fmt.Printf("Value of slice1 is %v\n", slice1)
	fmt.Printf("Length of slice1 is %d\n", len(slice1))
	fmt.Printf("Capacity of slice1 is %d\n", cap(slice1))
	fmt.Printf("Address of slice1 is %p\n", &slice1)
	fmt.Printf("Array address of slice1 is %p\n", &slice1[0])

	array2 := [10000]int{}
	slice1 = append(slice1, array2[0:10000]...)
	fmt.Println()
	fmt.Println("Append a large slice, the slice header address will not change")
	fmt.Printf("Address of slice1 is %p\n", &slice1)
	fmt.Printf("Array address of slice1 is %p\n", &slice1[0])
	slice1 = slice1[:0]
	fmt.Println()
	fmt.Println("Clear a slice")
	fmt.Printf("Value of slice1 is %v\n", slice1)
	fmt.Printf("Length of slice1 is %d\n", len(slice1))
	fmt.Printf("Capacity of slice1 is %d\n", cap(slice1))
	slice1 = append(slice1, array2[0:10000]...)
	slice1 = slice1[:0:0]
	fmt.Printf("Length of slice1 is %d\n", len(slice1))
	fmt.Printf("Capacity of slice1 is %d\n", cap(slice1))
	fmt.Printf("Value of slice1 is %v\n", slice1)
	fmt.Println()
}
