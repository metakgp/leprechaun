package leprechaun

type Field struct {
    key string
    validator func(string) bool
}

var fields = []Field{
    Field{
        "roll",
        validRoll,
    },
    Field{
        "email",
        validEmail,
    },
}
