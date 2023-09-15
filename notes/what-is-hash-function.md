# What is a Hash Function?

*All code will be linked to on the courses page.*

A hash function:
- Accepts arbitrary data
- Generates a result of fixed size using the data
- Generates the same result given the same input

Applying a hash function is often called "hashing"

The value returned from a bash function may be called a "hash value" or just a
"hash".

Example:

```go
// An example of a hash function, albeit a bad one
func hash(s string) int {
    return len(s) % 5
}
```

Uses length and mod 5 to calculate a hash.

All returns will be in the set: [0, 1, 2, 3 ,4]

Can have same output for two different inputs:
- "hi" => 2
- "Hello, world" => 2

*Run this on the Go Playground: <https://go.dev/play/p/yclBda-YR_y>*

Two inputs mapping to the same hash is called a "hash collision" or just a
"collision".

Collisions are inevitable with fixed output size due to infinite inputs with
fixed outputs.

In our `hash` func collisions are highly likely.

When we hash passwords, we will use a hashing func where collisions are very
unlikely.


## Hashes cannot be reversed

- Cannot take a hash value and a hash function and calculate the input
- This is partially because multiple inputs could result in the same
output

In our example we don't know if the original input length was `2`, `7`, `12`,
or something else where `% 5` == `2`.

## HMAC

Let's look at an example using [HMAC](https://en.wikipedia.org/wiki/HMAC).

```go
package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

func main() {
    // White HMAC uses a secret key, this is not the same as encryption
    // and cannot be used to decrypt a password.
    secretKeyForHash := "secret-key"

    password := "this is a totally secret password nobody will guess"

    // Setup our hashing function
    h := hmac.New(sha256.New, []byte(secretKeyForHash))

    // Write data to our hashing function
    h.Write([]byte(password))

    // Get the resulting hash
    result := h.Sum(nil)

    // The resulting hash is binary, so encode it to hex so we can read it as
    // letters and numbers
    fmt.Println(hex.EncodeToString(result))
}
```

*Run this on the Go playground: <https://go.dev/play/p/57Lh03BauIK>*

`HMAC` requires a secert key that is used to hash data.

HMAC is NOT encryption! It cannot be reversed.

The HMAC key gives us a way to generate unique hashes that others cannot
replicate without the key.
We will see why this matters later when we learn about session.

In the code we:
- Setup the hash func
- Write data to the hash func
- Get the result value
- Encode with hex so we can print out to the terminal

It is not a single func call, which is fine.

It is setup this way so we can stream input to it and eventually get the result
when we are ready.

*If we wanted to hash multiple things we would need to reset the hash func to clear the data.*

Not all binary data will map to strings, so we use hex encoding to ensure it
prints out fine.

## Uses for hash functions

- Hash maps (`map` in Go); see [Go maps in action](https://go.dev/blog/maps)
- Securing password
- Digitally signing data

Hash functions used for each will vary.

Digital signing can use HMAC because it has a key.

Passwords will use functions like `bcrypt`

maps will likely use a hashing func without a key.
Also looking for hash funcs faster than passwords use, since mitigating pw
cracking attacks isn't a goal for map hashing.

When setting up an authentication system, we should always use a
password-specific hashing function.

`bcrypt` will not only be more secure, it will handle some of the details (like 
salting for us).
