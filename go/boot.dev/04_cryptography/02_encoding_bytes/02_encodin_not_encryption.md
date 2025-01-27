# Encoding != Encryption

Encoding and encryption are not the same.

## Encoding

When we talk about encoding in the context of cryptography, we are talking about converting raw data (binary) into a format that can be stored or transmitted as text.

It is not safe to simply encode a secret message. Anyone who has access to the encoded message can decode it, there are no encryption keys involved. Anyone can trivially figure out the alphabet used to encode the message and decode it.

## Encryption

Encryption, on the other hand, is a secure way to transmit secret information. The critical difference is that encryption requires a key to decrypt the message. Without the key, the message is unreadable.
