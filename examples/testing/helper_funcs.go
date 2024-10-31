package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This file is for functions I don't want to appear on the main example

func truncateUsersTable(t *testing.T) {

}

func insertTestUsers([]User) {

}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func assertEquals(t *testing.T, got any, expected any) {
	assert.Equal(t, expected, got)
}
