# Bcrypt Review

Let's take a look at the output of the Bcrypt KDF that we just implemented:

```txt
$2a$13$Z9FCUt.58BSe5RdBQmPhPOfMQO4WDb8LwlY11Oy0wJ8HNXbSiNYSa
```

This isn't just a hash of the password input. It's actually a string with a format:

```txt
$ VERSION $ COST_FACTOR $ SALT HASH
```

To break it down:

- Version: `2a`
- Cost Factor: `13`
- Salt: `Z9FCUt.58BSe5RdBQmPhPO` (first 16-bytes of the last section)
- Hash: `fMQO4WDb8LwlY11Oy0wJ8HNXbSiNYSa`

## Version

The version is a string that identifies the version of the Bcrypt algorithm used to generate the hash. This is stored in the output so that we can use different versions of the algorithm in the future without breaking old hashes.

## Cost Factor

The cost factor is a number that determines how many rounds of hashing are performed. The higher the cost factor, the more rounds of hashing are performed. This makes the hash slower to compute, but also makes it more secure. We need to store the cost factor in the output so that we can use it to check the hash later.

## Salt

A salt is a random chunk of data added to a password before it is hashed so that its output hash will differ from the hash of the same password with a different salt. We'll talk about these in more detail soon. Again, we need it to be stored in the output so that we can use it to check the hash later.

## Hash

The hash is the output of the Bcrypt hashing algorithm. It is the result of hashing the password with the salt and cost factor.