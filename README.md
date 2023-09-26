## go-schema

- Zod-inspired project: https://github.com/colinhacks/zod
- Star the repository
- I am accepting new contribuition
- Each contribution has a significant impact within the Golang community

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
