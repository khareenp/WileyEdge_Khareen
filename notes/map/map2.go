package map1

import (
	"fmt"
	"sort"
)

func Map2() {
	// initialize like this
	ages := make(map[string]int) // map[key]value
	ages["joe"] = 25
	ages["jane"] = 32
	ages["alice"] = 31
	ages["bob"] = 27

	// can also use a literal
	// of course unicode strings can be keys
	// (can even be variable names)
	年龄 := map[string]int{
		"小明": 18,
		"小李": 20,
		"小张": 22,
		"小赵": 24,
	}

	fmt.Println(ages)
	fmt.Println(年龄)

	// delete a key with the delete function
	delete(ages, "bob")
	fmt.Println(ages)

	// the value of a key that is not found is the zero value of the map's value type.
	// bob is born again
	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages)

	// if you need to distinguish between an key's value being the zero
	// value and the key not being there, use the second return value
	// of the lookup
	if age, ok := ages["kevin"]; !ok {
		fmt.Printf("kevin is not in the map\n")
	} else {
		fmt.Printf("kevin is %d\n", age)
	}

	// we cannot take the address of a map element
	// growing a map might rehash existing elements into new storage locations
	// thus the pointer could become dangling
	// an important design goal of Go was to prevent dangling pointers
	// _ = &ages["bob"] // error

	// instead, use a map of pointers
	type Person struct {
		name  string
		title string
	}

	a := Person{name: "Jeff", title: "Mr"}
	b := Person{name: "Oz", title: "Mr"}
	c := Person{name: "May", title: "Miss"}
	var myMap = map[int]*Person{
		1: &a,
		2: &b,
		3: &c,
	}
	(*myMap[2]).title = "Dr"
	fmt.Printf("b: %v\n", b)

	// although when printing, the default behavior is to print in sorted order,
	// while iterating there is no guaranteed order, so you must manually sort the keys
	// (it's often the order of insertion, but this is implementation specific)

	fmt.Println("Iterated order:")
	// pre-allocating a slice for the keys with the length of the map is slightly more efficient
	// than using append because it doesn't need to compare the capacity every time or do any
	// reallocations
	names := make([]string, len(ages))
	i := 0
	for k := range ages {
		fmt.Printf("%s: %d\n", k, ages[k])
		names[i] = k
		i++
	}

	fmt.Println("Sorted order:")
	sort.Strings(names)
	for _, v := range names {
		fmt.Printf("%s: %d\n", v, ages[v])
	}

	// while delete, len, range, and lookups on a nil map are all safe,
	// storing to a nil map will panic
	// you must allocate before storing
	// var animals map[string]string
	// animals["Koko"] = "gorilla" // panic

	// maps are not comparable, so you cannot use == or !=
	// but we can check for deep equality
	âge := map[string]int{
		"joe":   25,
		"jane":  29,
		"alice": 31,
		"bob":   1,
	}
	fmt.Printf("equal(ages, âge): %t\n", equal(ages, âge))

	// check the badEqual function for explanation
	zeroAges := map[string]int{
		"小明": 0,
		"小李": 0,
		"小张": 0,
		"小赵": 0,
	}
	fmt.Printf("badEqual(zeroAges, ages): %t\n", badEqual(zeroAges, ages))

	// there is no set type in go, but you can
	// use a map of booleans
	// Can you think of a use for a set which contains true and false values?
	people := []string{"alice", "bob", "charlie", "alice", "dave", "elizabeth", "bob", "jane", "bob", "jim"}
	seen := make(map[string]bool)

	for _, name := range people {
		if _, ok := seen[name]; !ok {
			seen[name] = true
		}
	}
	fmt.Printf("len(people): %v\n", len(people))
	fmt.Printf("len(seen): %v\n", len(seen))
	fmt.Printf("seen: %v\n", seen)
}

// we can compare two maps by using a deep equal function
// we must check if elements exist and are the same
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		// note this condition.
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

// if you simply check xv != y[k], then
// equal(map[string]int{"A": 0}, map[string]int{"B": 42})
// will report as true because the map will return the zero value
// of the value type
func badEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		// bad
		if xv != y[k] {
			return false
		}
	}
	return true
}