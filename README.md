## go-schema

- Zod-inspired project: https://github.com/colinhacks/zod
- Star the repository
- Each contribution has a significant impact within the Golang community

---

### How to use

add the package

> go get github.com/mkafonso/go-schema

import in your project

```go
import (
  "github.com/mkafonso/go-schema/z"
)
```

--- 

### Basic Usage

`Create a schema for string`

```go
    // don't expect error when string is valid
    schema := z.NewStringSchema()
    result, err := schema.Parse("Luna")
    
    // expect an error when string is invalid and return custom error message
    schema := z.NewStringSchema()
    result, err := schema.Parse(42, "Custom error message")
    
    // expect an error when string is invalid and return the first custom error message
    schema := z.NewStringSchema()
    result, err := schema.Parse(42, "First custom error message", "Second custom error message")
    
    // check the string length with Min Method
    schema := z.NewStringSchema()
    result, err := schema.Min(2, "length must be at least 2 characters").Parse("Luna")
    
    // check the string length with Max Method
    schema := z.NewStringSchema()
    result, err := schema.Max(10, "length must not exceed 2 characters").Parse("Luna")
    
    // check if the string is a valid email address
    schema := z.NewStringSchema()
    result, err := schema.Email("custom error message").Parse("email@email.com")
    
    // can combine methods
    schema := z.NewStringSchema()
    result, err := schema.Min(100, "length custom error").Email("email error message").Parse("me@there.com")

```

`Create a schema for number`

```go
    // don't expect error when number is valid
    schema := z.NewNumberSchema()
    result, err := schema.Parse(42.5)
    
    // expect an error when number is invalid and return custom error message
    schema := z.NewNumberSchema()
    result, err := schema.Parse("Hi", "Custom error message")
```

---

### Comprehensive description for each method

![Screen Shot 2023-09-27 at 01 55 14](https://github.com/mkafonso/go-schema/assets/73212666/17e90457-3585-4dfd-8a43-37aec816ce60)


