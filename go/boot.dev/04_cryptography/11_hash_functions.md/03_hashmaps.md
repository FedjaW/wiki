# Hashmaps

Maps in Go, object literals in JavaScript, dictionaries in Python, and hashes in Ruby are all examples of hashmaps. Hashmaps are a data structure that uses a hash function under the hood to map keys to values.

The hash function used by a map in Go takes the key as an input and hashes it to an address in memory. The value associated with that key is then stored in the memory address.

We can create an empty map by using the make function. Maps in Go must be initialized before they can be used.

```go
// create a map called "ages" with string keys and int values 
ages := make(map[string]int)
```