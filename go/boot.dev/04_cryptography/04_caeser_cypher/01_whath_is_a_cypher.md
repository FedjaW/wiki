# What Is a Cipher?

> A cipher is an algorithm for performing encryption or decryption.

Ciphers are simply a series of computational steps that result in the scrambling and de-scrambling of a secret message. All of the encryption we've done so far in this course has involved ciphers.

## The Caesar Cipher

A Caesar cipher is an extremely simple substitution cipher where each letter in the plaintext is replaced by a different letter a fixed number of positions away in the alphabet. Caesar ciphers are named after Julius Caesar, so we've known about them for a long time!

Ralphie's Decoder ring in A Christmas Story is a great example of a Caesar cipher:

## Example

Using the English alphabet, and a key of left shift 3, we would get the following lookup table:

```
Plain:    ABCDEFGHIJKLMNOPQRSTUVWXYZ
Cipher:   XYZABCDEFGHIJKLMNOPQRSTUVW
```

We could then encrypt some text:

```
Plaintext:  THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG
Ciphertext: QEB NRFZH YOLTK CLU GRJMP LSBO QEB IXWV ALD
```

Decrypting the message is as simple as reversing the process and shifting the letters back to the rig
