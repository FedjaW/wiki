# Data Encryption Standard

The Data Encryption Standard is an example of a symmetric-key block cipher that utilizes a Feistel network.

DES was developed in the early 1970s at IBM and in 1977 a slightly modified version was published as an official Federal Information Processing Standard for the United States.

In 1997 the DESCHALL Project was the first group to publicly break DES encryption and won $10,000 for their efforts! DES was broken by a simple brute force attack, which is possible due to its small key size of just 56 bits.

## What's an IV?

An IV, or initialization vector, is a random value that is used to initialize a block cipher. It is used to ensure that the same plaintext always encrypts to a different ciphertext. Without an IV, the same plaintext would always encrypt to the same ciphertext which is a big security vulnerability.

## Assignment

For Cryptanalysis purposes, Passly keeps a DES implementation around so that our engineers can practice breaking it. Complete the encrypt function. It should use the standard library's crypto/des package to encrypt the plaintext with the given key.

Complete the encrypt function and its helper padMsg function. The decrypt function is already written for you.

### padMsg(plaintext []byte, blockSize int) []byte

The padWithZeros function is provided for you, but it only pads a single block. You'll need to find the last block in the message and pad that one. Essentially you need to ensure that the entire message length is a multiple of the block size.

### encrypt(key, plaintext []byte) ([]byte, error)

We'll be using DES in CBC mode. Here's an example from the Go documentation that shows how to encrypt a message.

1. Create a new cipher block
2. Pad the plaintext with zeros using padMsg
3. Generate a random iv and append it to the beginning of the ciphertext. It should be the same length as the block size.
4. Create a new encrypter
5. Encrypt the blocks and return the entire ciphertext

Return any errors that occur.

### Solution

```go
package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"io"
)

func encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext = padMsg(plaintext, des.BlockSize)

	ciphertext := make([]byte, des.BlockSize+len(plaintext))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[des.BlockSize:], plaintext)

	return ciphertext, nil
}

func padMsg(plaintext []byte, blockSize int) []byte {
	if len(plaintext)%blockSize != 0 {
		indexOfLastBlock := (len(plaintext) / blockSize) * blockSize
		fullBlocks := plaintext[:indexOfLastBlock]
		lastBlockPadded := padWithZeros(plaintext[indexOfLastBlock:], blockSize)
		fullBlocks = append(fullBlocks, lastBlockPadded...)
		return fullBlocks
	}
	return plaintext
}
```

// helper.go

```go
package main

import (
	"crypto/cipher"
	"crypto/des"
	"errors"
)

func decrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]
	if len(ciphertext)%des.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return ciphertext, nil
}

func padWithZeros(block []byte, desiredSize int) []byte {
	for len(block) < desiredSize {
		block = append(block, 0)
	}
	return block
}
```