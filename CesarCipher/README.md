# Caesar's cipher

Julius Caesar protected his confidential information by encrypting it using a cipher.
Caesar's cipher shifts each letter by a number of letters. For example, in the case of
a rotation by 3, w, x, y and z would map to z, a, b and c.

Original alphabet: `abcdefghijklmnopqrstuvwxyz`

Alphabet rotated +3: `defghijklmnopqrstuvwxyzabc`

## REQUIREMENTS

1. If the shift takes you past the end of the alphabet, just rotate back to the front of the alphabet.

2. If the shift number is a negative number, log an error.

3. The cipher only encrypts letters; all other characters remain unencrypted.

4. Build a function CaesarCipher(s string, n int) (string, error) to handle your implementation.

5. Create unit tests to validate previous requirements.

## EXAMPLE

```Go
s := "abc-defghi *jk* %lyz"
n := 3
```

The alphabet is rotated by 3, matching the mapping above.

The encrypted string is: `"def-ghijkl *mn* %obc"`

## BONUS

Implement a concurrent mechanism to handle the execution of Caesar's Cipher over multiple strings
