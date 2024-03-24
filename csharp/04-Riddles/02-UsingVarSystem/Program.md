```CSHARP
using System;

using var = System.Int32; // type alias

var var = default!;

Console.WriteLine(var);

// Riddle: 
// What will be printed to the console?

// For answere scroll down.
```


























































A: 0

In C# you can name a variable with a reserved keyword. To do so you need put an @ symbol before like so:
int @int = 5; this compiles successfully.
A special case is the var keyword, this can be used without an @ symbol. This has something to do with downwards compatibality.
The default means the default value if int32 which is 0. 
default <-> default(var) or default(Int32)
The ! can be ignored because it tells the compiler that this is not null, but it cant be null in any case because int32 is not nullable.
```