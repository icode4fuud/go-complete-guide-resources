package main

//Exemplify Interfaces, Dynamic Types & Limitations Sec 6:107
//sets up the need for Generics
func main() {}

func add(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	//^repeat code for float and string

	return nil
}
