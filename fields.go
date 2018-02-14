package main

type Field struct {
	key       string
	validator func(string) bool
}

var fields = []Field{
	{
		"roll",
		validRoll,
	},
	{
		"email",
		validEmail,
	},
}
