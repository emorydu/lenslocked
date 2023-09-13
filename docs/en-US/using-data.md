Using this data...

```go
user := User{
	Name: "Emory.Du",
	Email: "orangeduxiaocheng@gmail.com",
	...
}
```

We want to render the following.

```html
<body>
    <a href="/account">orangeduxiaocheng@gmail.com</a>
    ...
    <h1>Hello, Emory.Du</h1>
</body>
```

## Server-side

Server uses this 

```html
<body>
    <!-- A link to the users account details -->
    <a href="/account">{{.Email}}</a>
    <!-- ... -->
    <h1>Hello, {{.Name}}</h1>
</body>
```

Returns this

```html
<body>
    <a href="/account">orangeduxiaocheng@gmail.com</a>
    ...
    <h1>Hello, Emory.Du</h1>
</body>
```

## API

Server returns this

```json
{
  "name": "Emory.Du",
  "email": "orangeduxiaocheng@gmail.com",
}
```

## JS Frontend

```js
import React from 'react';

function Example() {
    const { name, email } = // data from our Go server
    
    return (
        <body>
            <h1>Hello, {name}!</h1>
            ...
            <!-- A link to the users account details -->
            <a href="/account">{email}</a>
        </body>
    );
}
```