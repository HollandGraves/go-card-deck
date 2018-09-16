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



*/