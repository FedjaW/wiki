# Caesar Cipher Security

The Caesar cipher is not secure in today's digital world.

Assuming that an attacker knows or guesses that a Caesar cipher has been used, several key pieces of information are leaked immediately:

- The language and character set of the original plaintext
- The length of the original plaintext

Frequency analysis can then be employed to quickly crack the code. On modern hardware, a Caesar cipher can be cracked instantly.

For example, in English, the letter `e` is by far the most commonly used. If the attacker can find the most common letter in the ciphertext, then that letter very likely maps to the plaintext `e`. Because a Caesar cipher uses a fixed-length shift, the key is immediately discovered and the ciphertext can be compromised.
