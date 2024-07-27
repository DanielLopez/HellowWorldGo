package GoBasics

import "fmt"

func ExecuteBasics() {

	//Hello World
	fmt.Println("Hello World")

	//Const
	PrintSection("CONST")
	const LabelName = "NAME"
	fmt.Println(LabelName)

	//Strings
	PrintSection("STRING")
	var name = "Daniel"
	fmt.Println(name)  //Printing Strings
	fmt.Println(&name) //Printing String Address Location '&' Pointers

	fmt.Println("A", "B", "C", "D", "E", "F", "G", "H", "I", "J")
	fmt.Println("Length", len(name)) //Length Function (Length of Strings or Arrays).
	Line()

	//Formatted Text
	name = "Denise + Daniel + Lana"
	fmt.Printf("Value: %s, Length: %d", name, len(name))

	//Integers
	PrintSection("MATH")

	Label("Integers")
	var intNum uint16 = 32767
	fmt.Println(intNum)

	Label("Float")
	var floatNum float32 = 123123.55
	fmt.Println(floatNum)

	Label("Adding Mixed Types")
	fmt.Println("NOTE: You cant add mixed types. Types must be cast to ensure they are compatible")
	fmt.Println(floatNum + float32(intNum))

	/*ARRAYS*/
	PrintSection(`ARRAYS (FIXED)

Fixed Size: An array has a fixed size. The size is part of the array's type, meaning once defined, it cannot be resized.
Defined Length: When you declare an array, you specify its length (number of elements it can hold) as part of its type.
Value Type: Arrays are value types. Assigning an array to another variable copies all the elements. Similarly, passing an array to a function copies the entire array, unless you use pointers.
Rarely Used Directly: In Go, direct use of arrays is less common because their fixed size makes them less flexible than slices.
	`)

	Label("ARRAYS")
	var array = [5]int{0, 1, 2, 3, 4}
	fmt.Println(array[0:5])

	array[0] = 0
	array[1] = 10
	array[2] = 20
	array[3] = 30
	array[4] = 40

	fmt.Println(array[0:5])
	fmt.Printf("BEFORE: Length=%d; Capaci 	ty=%d\n", len(array), cap(array))

	/*SLICES*/
	PrintSection(`SLICES (DYNAMIC)
Are Essentially Wrappers around an array that gives additional functions. Like the ability to use append.  (Think ArrayList)

Dynamic Size: A slice is a dynamically-sized, flexible view into the elements of an array. It can grow and shrink within the capacity of the underlying array.
Reference Type: Slices are reference types. When you assign a slice to another variable, both refer to the same underlying array. Any modifications made through one slice are visible through the other.
Built-in Functions: Slices are supported by built-in functions like append, which can add elements to a slice, potentially increasing its size.
More Commonly Used: Slices are more commonly used in Go because they are more flexible and provide the functionality needed for most array-like operations without the restrictions of a fixed size.
	`)

	Label("SLICES")
	var slice1 = []int{0, 100, 200, 300}
	fmt.Printf("[slice1]\t\tValue=%d Length=%d; Capacity=%d\n", slice1, len(slice1), cap(slice1))

	var slice2 = []int{0, 1000, 2000, 3000}
	fmt.Printf("[slice2]\t\tValue=%d Length=%d; Capacity=%d\n", slice2, len(slice2), cap(slice2))

	sliceCombo := append(slice1, slice2...) //NOTE: The := operator automatically knows the values is a variable and creates the valid .
	fmt.Printf("[sliceCombo]\tValue=%d Length=%d; Capacity=%d\n", sliceCombo, len(sliceCombo), cap(sliceCombo))

	/*BOOLEANS */
	Label("BOOLEANS")
	var boolFlag = true
	boolFlag = false
	fmt.Println(boolFlag)

	/*CONTROL STRUCTURES*/
	PrintSection("CONTROL STRUCTURES")
	Label("IF/ELSE")
	controlStructureIfElse(true)
	controlStructureIfElse(false)

	/*SWITCH*/
	Label("SWITCH")
	controlStructureSwitch(1)
	controlStructureSwitch(2)
	controlStructureSwitch(3)

	/*FOR LOOP*/
	//The for loop is the only loop available in Go. Here are some examples of how to use the for loop.
	Label("FOR LOOP")
	arrayLoop := [5]int{0, 1, 2, 3, 4}
	for i := 0; i < len(arrayLoop); i++ {
		fmt.Println(arrayLoop[i])
	}

	/*FOR RANGE LOOP*/
	//The range keyword is used in for loops to iterate over items of an array, slice, channel, or map.
	Label("FOR RANGE LOOP")
	for index, value := range arrayLoop {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	/*MAPS*/
	//Maps are a collection of key-value pairs. They are used to look up a value by its associated key.
	Label("MAPS")
	var map1 = map[string]int{"A": 1, "B": 2, "C": 3}
	fmt.Println(map1)

	//Adding a new key-value pair
	map1["D"] = 4
	fmt.Println(map1)

	valueExistsInMapByReference(&map1, "A")
	valueExistsInMapByReference(&map1, "B")
	//valueExistsInMapByReference(&map1, "X")
}

func valueExistsInMapByReference(map1 *map[string]int, key string) {
	valueExistsInMap(*map1, key)
}

func valueExistsInMap(map1 map[string]int, key string) {

	fmt.Printf("Address of Map: %p\n", &map1)

	value, exists := map1[key]
	if exists {
		fmt.Printf("Value Exists: %d\n", value)
	} else {
		fmt.Println("Value Does Not Exist")
	}
}

func controlStructureSwitch(switchVar int) {
	switch switchVar {
	case 1:
		fmt.Println("Case: One")
	case 2:
		fmt.Println("Case: Two")
	default:
		fmt.Println("Case: Unknown")
	}
}

func controlStructureIfElse(isEnabled bool) {
	if isEnabled {
		fmt.Println("Flag is Enabeld (TRUE)")
	} else {
		fmt.Println("Flag is Disabled (FALSE)")
	}
}
