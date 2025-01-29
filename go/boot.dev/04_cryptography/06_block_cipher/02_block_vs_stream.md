# Block vs. Stream

## Block Cipher

A block cipher is a deterministic algorithm operating on fixed-length groups of data, called blocks. Block ciphers are an example of symmetric encryption.

Popular block ciphers include:

- AES
- DES

## Stream Cipher

In a stream cipher, each plaintext digit is encrypted one at a time with the corresponding character of the keystream, to give a character of the ciphertext stream.

Popular stream ciphers include:

- RC4
- Salsa20

## Which Is Best?

One isn't necessarily better than the other. Stream ciphers are typically used when a stream of data must be encrypted bit by bit. For example, when encrypting a stream of video data in transit. Block ciphers are more typically used on static data, things like passwords in a password manager (like Passly).
