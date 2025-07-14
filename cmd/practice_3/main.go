package main

import "fmt"

func main() {
	fmt.Println("A simple program to demonstrate arrays, slices and maps in Go")

	var arr [3]int32 = [...]int32{10, 20, 30} // Create an array of int32 with 3 elements and initialize it with "..." being autosized

	fmt.Println("Array initialized with values:", arr)
	fmt.Println("Array elements and contiguous memory:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element at index %d: %d ; is stored at: %d\n", i, arr[i], &arr[i])
	}
	fmt.Println("Array length:", len(arr))

	// Omit the size of the array to create a slice
	var slice []int32 = []int32{40, 50, 60}
	fmt.Printf("Slice created from the array: %v\nLength is %v with capacity %v\n", slice, len(slice), cap(slice))
	// Note: The capacity of a slice is the maximum number of elements it can hold without reallocating memory
	// Use make(int32[], 3, 5) to create a slice with a specific length 3 and capacity 5

	slice = append(slice, 70) // Append a new element to the slice
	fmt.Printf("Updated slice: %v\nLength is %v with capacity %v\n", slice, len(slice), cap(slice))

	slice = append(slice, arr[:]...) // Append elements from the array to the slice using spread operator
	fmt.Printf("Slice after appending array elements: %v\nLength is %v with capacity %v\n", slice, len(slice), cap(slice))

	var map1 map[string]int32 = make(map[string]int32) // Create a map with string keys and int32 values
	fmt.Println("Map created:", map1)

	var radiation_map map[string]int32 = map[string]int32{"Alpha": 40, "Beta": 20, "Gamma": 30} // Create and initialize a map
	fmt.Println("Map created and initialized:", radiation_map)
	var radiation, ok = radiation_map["Alpha"] // Access a value from the map using a key
	if ok {
		fmt.Println("Radiation level of Alpha:", radiation)
	} else {
		fmt.Println("Alpha not found in radiation map")
	}

	delete(radiation_map, "Beta") // Delete a key-value pair from the map
	fmt.Println("Radiation map after deleting Beta:", radiation_map)

	for name := range radiation_map {
		fmt.Printf("Radiation level of %s: %d\n", name, radiation_map[name])
	}

	for name, level := range radiation_map {
		fmt.Printf("Radiation level of %s: %d\n", name, level)
	}
}
