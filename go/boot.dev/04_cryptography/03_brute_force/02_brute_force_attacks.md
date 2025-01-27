# Brute Force Attacks

Encryption keys must be kept secret, but there is another possible problem with keys: they can be guessed.

## What Is a Brute-Force Attack?

Brute-force attackers guess passwords, passphrases, and private keys in an attempt to eventually get the right answer and crack the security of a system. A brute-force attack is simple: the attacker systematically guesses every possible combination.

For example, if the attacker were guessing telephone numbers in the US, they would start guessing these values:

```
(000) 000-0000
(000) 000-0001
(000) 000-0002
...
(000) 000-0010
(000) 000-0011
```

## How Long Should Cryptographic Keys Be?

The longer the key, the harder it will be to guess. However, remember that we learned in the last exercise that the security of the key grows exponentially with the length of the key. A 256-bit key is not twice as hard to guess as a 128-bit key, it is `2^128` times harder to guess!
