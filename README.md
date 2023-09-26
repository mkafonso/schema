## go-schema

- Zod-inspired project: https://github.com/colinhacks/zod
- Star the repository
- I am accepting new contribuition
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
    result, err := schema.Message("custom error message").Parse(12)
    
    // expect an error when string length exceed Max value
    schema := z.NewStringSchema()
    result, err := schema.Max(3).Message("custom error message").Parse("Luna")
    
    schema := z.NewStringSchema()
    result, err := schema.Min(5).Message("custom error message").Parse("Hi")
```


`Create a schema for number`

```go
    // don't expect error when number is valid
    schema := z.NewNumberSchema()
    result, err := schema.Parse(12)
    
    // expect an error when number is invalid and return custom error message
    schema := z.NewNumberSchema()
    result, err := schema.Message("custom error message").Parse("not a number")
```

---

### Coercion


`Coercion to string`

```go
	schema := z.NewCoerceStringSchema()
	result, err := schema.Parse("Luna") // "Luna"

	schema := z.NewCoerceStringSchema()
	result, err := schema.Pare(12) // "12"

	schema := z.NewCoerceStringSchema()
	result, err := schema.Parse(true) // "true"
```


`Coercion to bool`

```go
	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse("true") // true

	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse("FalSe") // true

	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse(true) // true

	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse(false) // false

	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse("Hello") // true

	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse("") // false

	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse(10) // true

	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse(nil) // false

	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse(1) // true

	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse(0) // false

	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse([]interface{}{}) // true
```

---
  

![Screen Shot 2023-09-25 at 01 43 04](https://github.com/mkafonso/go-schema/assets/73212666/761bdaea-20df-4555-9f97-3f5986b7443d)
