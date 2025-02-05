# Bcrypt - Storing Passwords

KDFs are the best way to store passwords in web applications! As a back-end developer, this is critical to understand.

## Can I Store Passwords in Plain Text?

Storing passwords in a database in plain text is a huge security risk. If someone gets access to your database, they can see all of your users' passwords.

## Can I Hash Passwords With SHA-256?

No. SHA-256 is a hash function, but it's not a KDF. SHA-256 is a very fast hash function. Good KDFs like Bcrypt are designed to be slow. We'll talk more about why that's important soon.

## Assignment

At Passly, we store passwords securely (it would be sad if we didn't). Each user has a master password that they use to log into their cloud account. That password is hashed with Bcrypt before being stored.

Use the `golang.org/x/crypto/bcrypt` package to complete the `hashPassword()` and `checkPasswordHash()` functions. You do not need to modify the function signatures, just implement the Bcrypt library's API and do the `[]byte` <-> `string` conversions.

- Use a cost factor of `13`
- Docs for bcrypt.GenerateFromPassword
- Docs for bcrypt.CompareHashAndPassword

### Solution

```go
package main

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
```