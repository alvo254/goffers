# goffers
This is my personal go documentation, learning etc

This will be side note all detailed notes/explanations will be found in the notes section 
Tool used for better notes creation is obsidian

## Run in go 
 - run go run <filename>

## Build in go 
 - run go build <filename>


Typed and Untyped Constants
Constants can be typed or untyped. An untyped constant works exactly like a literal; it
has no type of its own, but does have a default type that is used when no other type can be
inferred. A typed constant can only be directly assigned to a variable of that type.
Whether or not to make a constant typed depends on why the constant was declared. If
you are giving a name to a mathematical constant that could be used with multiple
numeric types, then keep the constant untyped.


Blocks
Go lets you declare variables in lots of places. You can declare them outside of functions, as the parameters to functions, and as local variables within functions.


