# Perfect Security

A cipher is said to have perfect security if an attacker who has access to only the ciphertext can infer absolutely nothing of interest about the plaintext. Such perfect ciphers do exist, one such example is the "one-time pad".

## One Time Pad

The One Time Pad has perfect security because an attacker who has access to only the ciphertext can infer absolutely nothing of interest about the plaintext.

## Ciphers Used in Production Are Usually Not Perfectly Secure

Most production ciphers are not perfectly secure, but are "close enough". In short, trade-offs are made that add to the practical security of the system while sacrificing the perfect theoretical security of the cipher itself.

The big problem with the One Time Pad is that the key needs to be the same length as the message. That means to encrypt a 128 GB hard drive, I'd need a 128 GB key!! That's just not practical.

## Some Ciphers Are Just Crap

The Caesar cipher is a great example of a cipher that is NOT perfectly secure OR practically secure. As we demonstrated earlier, when given access to the ciphertext of a Caesar cipher, an attacker can see the positions and patterns of characters in the plaintext.

### The XOR Operator in Go

The `^` operator will XOR two `byte`s.

```go
byte1 := 0b01110000
byte2 := 0b10101000
xorRes := byte1 ^ byte2
fmt.Printf("%b\n", xorRes)
// 11011000
```
