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
  

![Screen Shot 2023-09-25 at 01 43 04](https://github.com/mkafonso/go-schema/assets/73212666/761bdaea-20df-4555-9f97-3f5986b7443d)
