# Deterministic

Hash functions must be deterministic. Deterministic just means not random. The same input to a deterministic function will always produce the same output.

Think of a Rubic's cube.

If I start with the cube unscrambled and twist edges at random, by the end I will end up with something that does not resemble anything close to what I started with. However, if I were to start over and do the exact same series of moves, I would be able to repeatedly get the same outcome. Even though the outcome may appear random, it isn’t at all. That is essentially what a good hash function does.

Determinism is important for securely storing a password. For instance, let’s pretend my password is `ex@mpl3P@ss`

I can use a hash function to scramble it:

```txt
ex@mpl3P@ss → "2f73f64c5da77728d135069d6e6e67363961efa7c27ebed4f3230f29096c8dc1
```

Now, if anyone finds the scrambled version, they won't be able to know my original password! This is important because it means that as a website developer, I only need to store the hash (scrambled data) of my user’s password to be able to verify it. When a user signs up, I hash the password and store it in my database. When the user logs in, I just hash what they typed in and compare the two hashes. Because a given input always produces the same hash, this works every time.

If a website stores passwords in plain text (not hashed) it is a huge breach of security. If someone were to hack that site’s database and find all the emails stored with plain-text passwords, they could then use those combinations and try them on other websites.