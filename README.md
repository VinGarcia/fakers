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
	"testing"
	"time"

	"github.com/vingarcia/fakers"
)

type User struct {
	ID        int
	Name      string
	IsAdmin   bool
	CreatedAt time.Time
}

func fake[V any](t *testing.T, customValues map[string]any) V {
	var v V
	err := fakers.New(&v, customValues)
	if err != nil {
		t.Fatalf("error creating fake type: %s", err)
	}

	return v
}

func TestListUsers(t *testing.T) {
	tests := []struct {
		adminsOnly    bool
		dbUsers       []User
		expectedUsers []User
	}{
		{
			adminsOnly: true,
			dbUsers: []User{
				fake[User](t, map[string]any{
					"ID":      1,
					"IsAdmin": true,
				}),
				fake[User](t, map[string]any{
					"ID": 2,
				}),
			},
			expectedUsers: []User{
				fake[User](t, map[string]any{
					"ID":      1,
					"IsAdmin": true,
				}),
			},
		},
	}

	for _, test := range tests {
		truncateUsersTable(t)

		insertTestUsers(test.dbUsers)

		svc := NewUserService()

		users, err := svc.ListUsers(test.adminsOnly)
		assertNoError(t, err)
		assertEquals(t, users, test.expectedUsers)
	}
}
```
