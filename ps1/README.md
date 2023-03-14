# Lamport signature

### Generatekey()
    Make the key pair

private key -> 256 blocks. Each blocks have 256bits. 

public key -> hashes 512 pairs of random numbers (bit) in private key.
    Hash(private_key) -> will be the public_key

### Sign(private_key, message)
    Signs a message given a private_key. Return a signature.
No matter is signing a message or a file. The first thing is making sure it has fix length.
Therefore, it usually Hash(message) -> it's given a fixed length.
Pick private_key blocks to reveal based on bits of message of sign.

i.e Hash("Hi") -> c01a4cfa25cb895cdd0bb25181ba9c1622e93895a6de6f533a7299f70d6b0cfb
               -> 11000000000110100100110011100101110001001010111001101110..... (256 length)
            Which is 256 bits. 1 -> reveal private_key row 1
                               0 -> reveal private_key row 0
                1 -> 0001
                a -> 1010

Then Return the signature which are revealed private_keys.

### Verify(public_key, message, signature): Boolean
    Verify a signature on a message from a public_key. Return a boolean.
Hash(signature of each blocks) and check it maps into the public_key.

---

### Summary
Generatekey() make personal/computer/information idenitfy 
which produce private_keys and public_keys

Sign() using the message itself to reveal private_keys. That produces unique signature.
In other words, Hash(message).toBinary + choose part of private_keys => signature.
Note:
    message just reveal which part should be revealed. 
    signature is part of private_keys.

Verify() Hash(signature) and check public_key. Since signature is part of private_keys 
therefore, it can be checked by public_key to verify if the message is from myself.


