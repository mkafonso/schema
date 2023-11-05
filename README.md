  <h2 align="center">
   ##  GO SCHEMA ## 
  </h2>


This open-source project is a schema validation toolkit inspired by [zod.dev](https://zod.dev/). Feel free to use and modify it for any purpose. If you have any suggestions for improvement, please don't hesitate to reach out. Thank you!


## Contents

- [Getting started](#how-to-use)
- Basic Usage
  - [String Validation](#string-validation)
 

## <div id="how-to-use" />How to use

Add the package to your project using the following command:

```go
go get github.com/mkafonso/go-schema
```

Import the package in your project:

```go
import (
  "github.com/mkafonso/go-schema/z"
)
```

Once you've added the package and imported it into your project, you can start using the schema validation toolkit for your specific needs.


## <div id="string-validation" />String Validation

In this section, we'll explore the fundamental usage of the schema validation toolkit for strings using the go-schema package. You'll learn how to create a schema and perform various string validations.

In the first example, we create a string schema using z.NewStringSchema() and attempt to parse the string "Luna." Since "Luna" is a valid string, no error is expected.
```go
// Create a schema for string
schema := z.NewStringSchema()
result, err := schema.Parse("Luna")
```

In the second example, we again create a string schema. This time, we attempt to parse the value 42 as a string. Since 42 is not a valid string, we expect an error with the custom error message provided.
```go
// Expect an error when the string is invalid and return a custom error message
schema := z.NewStringSchema()
result, err := schema.Parse(42, "Custom error message")
```

In the third example, we use the same schema to parse a non-string value (42) and expect an error. The provided custom error messages are checked in order, and the first one is returned in the error message.
```go
// Expect an error when the string is invalid and return the first custom error message
schema := z.NewStringSchema()
result, err := schema.Parse(42, "First custom error message", "Second custom error message")
```

In the next example, we create a string schema and use the Min method to check the length of the string. We expect the string to be at least 2 characters long, and if it's not, a custom error message will be returned.
```go
// Check the string length with the Min Method
schema := z.NewStringSchema()
result, err := schema.Min(2, "length must be at least 2 characters").Parse("Luna")
```

Similarly, in the following example, we use the Max method to ensure that the string does not exceed a length of 10 characters. If it does, a custom error message is returned.
```go
// Check the string length with the Max Method
schema := z.NewStringSchema()
result, err := schema.Max(10, "length must not exceed 10 characters").Parse("Luna")
```

In the penultimate example, we check whether the string is a valid email address using the Email method. If the string is not a valid email, a custom error message is provided.
```go
// Check if the string is a valid email address
schema := z.NewStringSchema()
result, err := schema.Email("custom error message").Parse("email@email.com")
```

Finally, you can combine multiple validation methods to perform various checks on the string. In this example, we first check the string's length and then verify if it's a valid email address. If any validation fails, the respective custom error message is returned.
```go
// Combine methods
schema := z.NewStringSchema()
result, err := schema.Min(100, "length custom error").Email("email error message").Parse("me@there.com")
```
