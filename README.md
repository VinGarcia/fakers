# Fakers Creation Library

This project facilitates creating fake instances of Go structs for
testing purposes.

The library creates these fakers using a set of deterministic fake
values by default while still allowing the user to overwrite specific
fields.

## Usage Example

```golang
package main

import (
	"fmt"

	"github.com/vingarcia/fakers"
)

type User struct{
  ID int `fakeAs:"zero"`
  Name string
  IsAdmin bool
  CreatedAt time.Time
}

func TestListUsers(t *testing.T) {
  tests := struct{
      adminsOnly bool
      dbUsers []User
      expectedUsers []User
  }{
    {
      adminsOnly: true,
      dbUsers: []User{
        fakers.New[User](map[string]any{
          "ID": 1,
          "IsAdmin": true,
        }),
        fakers.New[User](map[string]any{
          "ID": 2,
        }),
      },
      expectedUsers: []User{
        fakers.New[User](map[string]any{
          "ID": 1,
          "IsAdmin": true,
        }),
      },
    },
  }

  for _, test := range tests {
    truncateUsersTable(t)

    insertTestUsers(test.dbUsers)

    users, err := userService.ListUsers(adminsOnly)
    assertNoError(t, err)
    assertEquals(t, users, test.expectedUsers)
  }
}
```
