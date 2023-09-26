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
    str := z.NewStringSchema()
    result, err := str.Parse("Luna")
    
    // expect an error when string is invalid and return custom error message
    str2 := z.NewStringSchema()
    result2, err2 := str2.Message("custom error message").Parse(12)
    
    // expect an error when string length exceed Max value
    str3 := z.NewStringSchema()
    result2, err := str3.Max(3).Message("custom error message").Parse("Luna")
    
    str4 := z.NewStringSchema()
    result4, err4 := str4.Min(5).Message("custom error message").Parse("Hi")
```


`Create a schema for number`

```go
    // don't expect error when number is valid
    number := z.NewNumberSchema()
    result, err := number.Parse(12)
    
    // expect an error when number is invalid and return custom error message
    number2 := z.NewNumberSchema()
    result2, err2 := number2.Message("custom error message").Parse("not a number")
```

---

### Coercion


`Coercion to string`

```go
	schema := z.NewCoerceStringSchema()
	result, err := schema.Parse("Luna") // "Luna"

	schema2 := z.NewCoerceStringSchema()
	result2, err2 := schema2.Parse(12) // "12"

	schema3 := z.NewCoerceStringSchema()
	result3, err3 := schema3.Parse(true) // "true"
```


`Coercion to bool`

```go
	schema := z.NewCoerceBoolSchema()
	result, err := schema.Parse("true") // true

	schema2 := z.NewCoerceBoolSchema()
	result2, err2 := schema2.Parse("FalSe") // true

	schema3 := z.NewCoerceBoolSchema()
	result3, err3 := schema3.Parse(true) // true

	schema4 := z.NewCoerceBoolSchema()
	result4, err4 := schema4.Parse(false) // false

	schema5 := z.NewCoerceBoolSchema()
	result5, err5 := schema5.Parse("Hello") // true

	schema6 := z.NewCoerceBoolSchema()
	result6, err6 := schema6.Parse("") // false

	schema7 := z.NewCoerceBoolSchema()
	result7, err7 := schema7.Parse(10) // true

	schema8 := z.NewCoerceBoolSchema()
	result8, err8 := schema8.Parse(nil) // false

	schema9 := z.NewCoerceBoolSchema()
	result9, err9 := schema9.Parse(1) // true

	schema10 := z.NewCoerceBoolSchema()
	result10, err10 := schema10.Parse(0) // false

	schema11 := z.NewCoerceBoolSchema()
	result11, err11 := schema11.Parse([]interface{}{}) // true
```

---
  

![Screen Shot 2023-09-25 at 01 43 04](https://github.com/mkafonso/go-schema/assets/73212666/761bdaea-20df-4555-9f97-3f5986b7443d)
