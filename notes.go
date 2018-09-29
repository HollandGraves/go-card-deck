package main

/*



GO NOTE:
:= sign

	// the two statements:
	//		var card string = "Ace of Spades"
	//		card := "Ace of Spades"
	// are totally equal to each other. The first case is a very specific form of declaration. Although, you do not need this level of specificity as the
	// go compiler will be told to intuit the data type being stored when using the := notation



GO NOTE:
= sign

	// after a variable has been declared with := then to reassign a new value to the same variable you use simply the = sign



GO NOTE:
intializing variables

	// you can initialize a variable outside of the func main() {} function, although you cannot assign a value to it
	// you can also initialize a variable, e.g. var card int
	//		without immediately assigning a variable to it; or later assigning a variable to it.



GO NOTE:
creating new functions outside of func main

	// whenever you create a function outside of main to be called inside of main, you use normal syntax as with most languages when writing functions
	// one thing, is when you're using a function that e.g. only returns a value e.g. a string you will have to put the data type of the item being returned
	// e.g. func newCard() string {return "Ace of Spades"}



GO NOTE:
Arrays in go

	// there are two basic data structures for handeling records in Go: Arrays and Slices
	// Arrays are of fixed length
	// Slices are arrays that can grow or shrink at will
	// Slices and Arrays must both be defined with ONE data type
	// 		repeating: a slice and/or an array may only have one data type in it
	// e.g. newSlice := []string{"one", "two", "three"}



GO NOTE:
Adding new elements to slices

	// to add a new element to a slice you would take the variable the slice is stored in and set it equal to the append() function,
	// 		while using at least two arguments:
	// 		one argument will be the variable that contains the slice (so as to re-store the original
	// 		slice's values into the new iteration of the cards variable; very similar to how to store values into an array using a spread
	// 		operator)
	// 		the other arugment will be the new values
	// e.g.
	// newSlice := []string{"one", "two", "three"}
	// newSlice = append(newSlice, "four", "five")
	//
	// something especially important to note here is that the newSlice variable that is equal to the append statement will be returning
	// 		a whole new version of newSlice. Go essentially goes back to the space where newSlice was originally written, deletes that entry
	// 		and writes and new iteration of newSlice, elsewhere in the RAM, database, etc



GO NOTE:
Iterating over a slice

	// to iterate over a slice you will use the following syntax
	// e.g.
	// newSlice := []string{"one", "two", "three"}
	// for i, variablePlaceholder := range newSlice {
	// 		fmt.Println(i, variablePlaceholder)
	// }
	// the important thing to notice here is that at each iteration, just as mentioned in the last section in the above note, the data for i and
	// 		variablePlaceholder are comepletely being deleted and recreated with new values
	// 		such that, at the first iteration of i:
	// 			i := 0
	// 			variablePlaceholder := "one"
	// 		then at the second iteration of i:
	//			i := 1
	// 			variablePlaceholder := "two"
	// at each step above i is being deleted from memory and recreated, same as variablePlaceholder
	//
	// also important, breaking down each step of the syntax: for i, placeHolderVariable := range newSlice {}
	// 		here the most important part is the "range" keyword, the range keyword is one of many keywords that signifies the for loop iterating over
	// 		every item in the slice.
	// also notice how a := symbol is being used, this supports that Go will delete each entry each iteration, search for the value type of the new
	// 		iteration, and resave the new iteration under the same placeHolderVariable name
	//
	// finally, to mention, in the syntax: for i, card := range cards {
	//											fmt.Println(card)
	// 										}
	// 		WONT compile, because the i variable is not being used.
	//		something about Go is that if a variable is being declared IT MUST BE USED, so you must actually use the syntax as follows:
	//			for i, card := range cards {
	// 				fmt.Println(i, card)
	// 			}



GO NOTE:
The Go version of a Class and Extending A Class

	// first the Go syntax for defining a "class" and "extending" the class to another "class"
	//
	// e.g.
	// type deck []string
	//
	// in this example we're saying that the "class" (i.e. type) named deck is "extending" the "class" []string (i.e. a slice that only contains strings)
	// in other words we're adding deck as a new data type and giving it the properties that a (e.g. []string) slice that only returns strings has



GO NOTE:
Creating a receiver function

	// a receiver function is a function of a type (e.g. type deck []string) that can take an argument and is tied access as a method of anything that
	// 		uses the deck type of slice
	// this is the syntax for a receiver function
	//
	// e.g.
	// func (d deck) print() {
	//
	// }
	//
	// in this case you will define the data type that print() is
	// 		it is a function, so we will use "func"
	// second you will use the receiver notation
	// 		(d deck)
	// 		this will allow the function to be used through dot notation on any variable that contains the use of the data type slice (which remember is
	// 		a slice data type that contains strings at heart with a new name due to the type deck []string syntax)
	// third we will write the name of the function we want to be able to use through dot notation



GO NOTE:
Using a receiver function

	// in the case of using a receiver function that is tied to a "class" (aka type that is uniquely named and extending the properties of another keyword
	// 		data type), you use the function through dot notation
	//
	// so imagining we have a type deck []string and a receiver function func (d deck) print() {} already defined in another file
	// 		we would use the print() function as so:
	//
	// e.g. start
	//
	// 			deck.go file
	// package main
	//
	// import "fmt"
	//
	// type deck []string
	//
	// func (d deck) print() {
	// 		for i, card := range d {
	// 			fmt.Println(i, card)
	// 		}
	// }
	//
	//			main.go file
	// package main
	//
	// import "fmt"
	//
	// func main() {
	// 		cards := deck{"Ace of Spades", "King of Hearts"}
	// 		cards = append(cards, "Queen of Clubs")
	// 		cards.print() 		// notice how the print() method is being accessed from the deck.go file, and is being used on a variable that has the deck data type stored in it, thus giving the variable access to the methods that have deck within their receiver function
	// }
	//
	// e.g. finish
	//
	// Because of both files using package main, and presumably being in the same folder, they both have access to each other's code
	// In this case the values of the cards variable are being passed as an argument into the print() function, in this case the value of cards is being
	// 		passed in as d, within the receiver function print(), and is being used as the range of the print() function.
	//
	// It is common practice to name the receiver argument (e.g. d) as the first letter, or first two letters, of the data type the receiver function is tied to [e.g. (d deck)]



GO NOTE:
using specific ranges of a slice

	// Very similarly to how you can access an item in an array (e.g. array[0] ) you can access items in a slice!
	// even more so, you can access a range of items in a slice with the syntax
	//
	// e.g.
	// slice[0:2]
	//
	// this would access the items STARTING FROM 0 and UP TO BUT NOT INCLUDING 2
	// so the items 0 and 1 would get selected from the slice[0:2] array
	//
	// you can also access all the items of a slice STARTING FROM A SPECIFIC POINT using the syntax
	// e.g.
	// slice[0:]
	// this would start from the 0 indexed item and select all items up to the end of the slice INCLUDING THE LAST ITEM
	// you could also do
	// e.g.
	// slice[2:]
	// this would start from the 2 indexed item and include all the indexed items starting from 2 onward, including the last item
	//
	// you can also access all the items from the beggining of the slice up until the specified point, NOT INCLUDING THE SPECIFIED POINT
	// e.g.
	// slice[:7]
	// this would start from the 0 indexed item and include all items UP TO BUT NOT INCLUDING the 7th indexed item



GO NOTE:
setting arguments in a function

	// when setting arguments in a function you must specify the name of the argument and it's type. Go is a statically typed langauge after all
	// e.g.
	// func dealCards (d deck, handSize int) {}
	//
	// in this case above the first argument is referenced in the function by the letter d and is of custom type defined deck
	// in the case of the second argument, it is referenced by handSize and is of data type int



GO NOTE:
remembering to set return types for functions

	// just like you declare a data type for arguments in functions, if your function returns at least one value, the value must be specified
	// 		at the end of the function declaration right before the curly braces
	//
	// e.g.
	// func deal(d deck, handSize int) (deck, deck) {}
	//
	// in this case above the function's name is deal
	// it has an argument d of type deck, and an argument handSize of type int
	// and returns two values, the first value will be of type deck and the second value will be of type deck




*/
