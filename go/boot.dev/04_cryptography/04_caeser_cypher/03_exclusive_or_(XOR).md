# Exclusive or (XOR)

The "exclusive or" operator is very important in cryptography. 

XOR, or “exclusive or” operates on binary data. It returns `true` if both of its inputs are opposites (one false and one true), otherwise, it returns false. You may see the operator written this way: `⊕`.

```go
package main

func xor(lhs, rhs bool) bool {
	return lhs != rhs
}
```
