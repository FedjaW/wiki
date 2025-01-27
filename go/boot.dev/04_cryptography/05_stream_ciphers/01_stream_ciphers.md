# Stream Ciphers

The Ceasar cipher, other substitution ciphers, and the "one-time pad" are all examples of stream ciphers.

A stream cipher is a symmetric key cipher where plaintext digits are combined with a key stream. In a stream cipher, each plaintext digit is encrypted one at a time with the corresponding digit of the keystream, to give a digit of the ciphertext stream.

## Assignments

We've been asked to update our "One Time Pad". For performance reasons, we'll now be reading and writing the data using Go channels. This will allow us to encrypt the data as it's read in from an external database. Rather than storing the entire message in memory (which could be many Gigabytes) and then decrypting it all at once, this new `crypt` function can do it one character at a time.

### crypt(textCh, keyCh <-chan byte, result chan<- byte)

Read one `byte` at a time from the `textCh` and `keyCh` channels. Perform an XOR operation on the two bytes and write the result to the `result` channel. Stop once either channel is closed.

Be sure to close the `result` channel when you're done writing to it.

#### Solution

 ```go
 package main

func crypt(textCh, keyCh <-chan byte, result chan<- byte) {
	defer close(result)
	for {
		textChar, textOk := <-textCh
		if !textOk {
			return
		}
		keyChar, keyOk := <-keyCh
		if !keyOk {
			return
		}
		result <- textChar ^ keyChar
	}
}
 ```