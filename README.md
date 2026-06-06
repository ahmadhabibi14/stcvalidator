# Go Struct Validator

```sh
go get -u github.com/ahmadhabibi14/stcvalidator
```

### Usage
```go
package main

import (
  "fmt"
  "github.com/ahmadhabibi14/stcvalidator"
)

type User struct {
  Name  string  `json:"name" validate:"required"`
  Age   int64   `json:"age" validate:"required,max=200"`
}

func main() {
  myUser := User{
    Name: "Thor",
    Age: 1500,
  }

  err := stcvalidator.Validate(myUser, stcvalidator.MapErrMsg{
    "Name": {
      "required": "Name cannot be empty",
    },
    "Age": {
      "required": "Age cannot be empty",
      "max": "Maximal age is 200 year old",
    },
  })

  if err != nil {
    fmt.Println(err)
  }
}
```