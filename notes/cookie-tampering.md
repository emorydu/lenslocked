# Digitally Signing

```json
{
    "id": 123,
    "signature": "60fc0ff9d3f4132cbe7b9dd12ef51837d8e"
}
```

```go
HMAC("secret-key-1",`{"id": 123}`) = 
// 60fc0ff9d3f4132cbe7b9dd12ef51837d8e...

// Different key, different result
HMAC("secret-key-2", `{"id": 123}`) = 
// b91c4b45a02dc9f92725d93ec4a7bfd6a4e...

HMAC("secret-key-3", `{"id": 333}`) = 
// 78257ef3e303a46b934fccf7ed9c01d7f77...
```

Playground Link: <https://go.dev/play/p/942QmEci-ZT>
*Available on course website for this lesson.*

JWTs - a standard for digitally signing JSON data

## Obfuscation:
```
user           | random_string
-------------- | --------------
jon@calhoun.io | 7r1mpIjJcl
bob@bob.com    | UkpEUzFPJy
```

This approach is commonly referred to as *sessions*, and the random string is a
*session token*.

We will use this approach and learn more about it in the next section of the course.

# Wht not JWTs?

Complexity without enough benefits.

- Expiration
- `{"alg": "none"}`
- Refresh Token, so we need sessions anyway?


Benefits?
- No need to query DB table to see who a user is from session token

This isn't really slow and caching can make it even faster if it becomes an 
issue with way less complexity.

All of these things can be dealt with, but my point is that this introduces
complexity without any real benefits for us.