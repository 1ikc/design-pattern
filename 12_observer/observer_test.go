package observer

func ExampleSubject_Notify() {
	sub := &Subject{}
	sub.Add(&Observer1{})
	sub.Add(&Observer2{})
	sub.Notify("hi")

	// unordered Output:
	// Observer1: hi
	// Observer2: hi
}