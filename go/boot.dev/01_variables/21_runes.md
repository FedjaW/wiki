# Runes and String Encoding

In many programming languages (cough, C, cough), a "character" is a single byte. Using ASCII encoding, the standard for the C programming language, we can represent 128 characters with 7 bits. This is enough for the English alphabet, numbers, and some special characters.

In Go, strings are just sequences of bytes: they can hold arbitrary data. However, Go also has a special type, rune, which is an alias for int32. This means that a rune is a 32-bit integer, which is large enough to hold any Unicode code point.

When you're working with strings, you need to be aware of the encoding (bytes -> representation). Go uses UTF-8 encoding, which is a variable-length encoding.


## What does this mean?

There are 2 main takeaways:

1. When you need to work with individual characters in a string, you should use the rune type. It breaks strings up into their individual characters, which can be more than one byte long.
2. We can use zany characters like emojis and Chinese characters in our strings, and Go will handle them just fine.

