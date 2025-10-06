package main

import "fmt"

// Example 1: Basic pointer operations
func basicPointers() {
	fmt.Println("\n=== Example 1: Basic Pointer Operations ===")
	
	num := 42
	ptr := &num
	
	fmt.Printf("Value of num: %d\n", num)
	fmt.Printf("Address of num: %p\n", &num)
	fmt.Printf("Value stored in ptr: %p\n", ptr)
	fmt.Printf("Value pointed by ptr: %d\n", *ptr)
	
	// Modify through pointer
	*ptr = 100
	fmt.Printf("After modification, num: %d\n", num)
}

// Example 2: Pointers with functions (pass by reference)
func modifyValue(x *int) {
	*x = *x * 2
}

func modifyValueCopy(x int) {
	x = x * 2
}

func functionPointers() {
	fmt.Println("\n=== Example 2: Pointers with Functions ===")
	
	value := 10
	fmt.Printf("Original value: %d\n", value)
	
	// Pass by value (copy)
	modifyValueCopy(value)
	fmt.Printf("After modifyValueCopy: %d (unchanged)\n", value)
	
	// Pass by reference (pointer)
	modifyValue(&value)
	fmt.Printf("After modifyValue: %d (changed)\n", value)
}

// Example 3: Pointers with structs
type Person struct {
	Name string
	Age  int
}

func updatePerson(p *Person, newName string, newAge int) {
	p.Name = newName
	p.Age = newAge
}

func structPointers() {
	fmt.Println("\n=== Example 3: Pointers with Structs ===")
	
	person := Person{Name: "Alice", Age: 25}
	fmt.Printf("Before: %+v\n", person)
	
	updatePerson(&person, "Bob", 30)
	fmt.Printf("After: %+v\n", person)
	
	// Direct pointer creation
	personPtr := &Person{Name: "Charlie", Age: 35}
	fmt.Printf("Person via pointer: %+v\n", *personPtr)
}

// Example 4: Pointer to pointer
func pointerToPointer() {
	fmt.Println("\n=== Example 4: Pointer to Pointer ===")
	
	num := 100
	ptr := &num
	ptrToPtr := &ptr
	
	fmt.Printf("Value: %d\n", num)
	fmt.Printf("Pointer value: %d\n", *ptr)
	fmt.Printf("Pointer to pointer value: %d\n", **ptrToPtr)
	
	// Modify through pointer to pointer
	**ptrToPtr = 200
	fmt.Printf("After modification: %d\n", num)
}

// Example 5: Nil pointers and checking
func nilPointers() {
	fmt.Println("\n=== Example 5: Nil Pointers ===")
	
	var ptr *int
	fmt.Printf("Nil pointer: %v\n", ptr)
	
	if ptr == nil {
		fmt.Println("Pointer is nil, initializing...")
		value := 50
		ptr = &value
	}
	
	fmt.Printf("After initialization: %d\n", *ptr)
}

// Example 6: Pointers with slices and maps
func slicePointers() {
	fmt.Println("\n=== Example 6: Pointers with Slices ===")
	
	// Slices are reference types, but we can still use pointers
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original slice: %v\n", slice)
	
	modifySlice(&slice)
	fmt.Printf("After modification: %v\n", slice)
}

func modifySlice(s *[]int) {
	*s = append(*s, 6, 7, 8)
	(*s)[0] = 100
}

// Example 7: Returning pointers from functions
func createPerson(name string, age int) *Person {
	p := Person{Name: name, Age: age}
	return &p // Safe in Go due to escape analysis
}

func returningPointers() {
	fmt.Println("\n=== Example 7: Returning Pointers ===")
	
	person := createPerson("David", 28)
	fmt.Printf("Created person: %+v\n", *person)
}

// Example 8: Pointer receivers in methods
type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Decrement() {
	c.count--
}

func (c Counter) GetCount() int {
	return c.count
}

func methodPointers() {
	fmt.Println("\n=== Example 8: Pointer Receivers in Methods ===")
	
	counter := Counter{count: 0}
	fmt.Printf("Initial count: %d\n", counter.GetCount())
	
	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Printf("After 3 increments: %d\n", counter.GetCount())
	
	counter.Decrement()
	fmt.Printf("After 1 decrement: %d\n", counter.GetCount())
}

// Example 9: Array of pointers
func arrayOfPointers() {
	fmt.Println("\n=== Example 9: Array of Pointers ===")
	
	a, b, c := 10, 20, 30
	ptrArray := []*int{&a, &b, &c}
	
	fmt.Println("Values through pointer array:")
	for i, ptr := range ptrArray {
		fmt.Printf("Index %d: %d\n", i, *ptr)
	}
	
	// Modify through pointers
	*ptrArray[0] = 100
	fmt.Printf("After modification, a = %d\n", a)
}

// Example 10: Swap function using pointers
func swap(x, y *int) {
	*x, *y = *y, *x
}

func swapExample() {
	fmt.Println("\n=== Example 10: Swap Function ===")
	
	a, b := 5, 10
	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	
	swap(&a, &b)
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)
}

func main() {
	fmt.Println("=== Go Pointers Comprehensive Examples ===")
	
	basicPointers()
	functionPointers()
	structPointers()
	pointerToPointer()
	nilPointers()
	slicePointers()
	returningPointers()
	methodPointers()
	arrayOfPointers()
	swapExample()
	
	fmt.Println("\n=== All examples completed! ===")
}