// custom_marshal_test.go
package custom_marshal

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUser_MarshalJSON(t *testing.T) {
	user := &User{
		ID:        1,
		Name:      "John",
		Email:     "John@Example.com",
		Password:  "Secret123",
		IsActive:  true,
		CreatedAt: time.Date(2026, 6, 26, 12, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2026, 6, 26, 12, 0, 0, 0, time.UTC),
	}

	data, err := user.MarshalJSON()
	require.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	require.NoError(t, err)

	require.Equal(t, "JOHN", result["username"])
	require.Equal(t, "JOHN@EXAMPLE.COM", result["email"])
	require.Equal(t, "SECRET123", result["password"])
	require.Equal(t, true, result["is_active"])

	// проверяем, что оригинал не мутировал
	require.Equal(t, "John", user.Name)
	require.Equal(t, "John@Example.com", user.Email)
	require.Equal(t, "Secret123", user.Password)
}

func TestUser_UnmarshalJSON(t *testing.T) {
	input := `{"id":1,"username":"JOHN","email":"JOHN@EXAMPLE.COM","password":"SECRET123","is_active":true,"created_at":"2026-06-26T12:00:00Z","updated_at":"2026-06-26T12:00:00Z"}`

	var user User
	err := user.UnmarshalJSON([]byte(input))
	require.NoError(t, err)

	require.Equal(t, "john", user.Name)
	require.Equal(t, "john@example.com", user.Email)
	require.Equal(t, "secret123", user.Password)
}

func TestUser_MarshalUnmarshalRoundTrip(t *testing.T) {
	original := &User{
		ID:       2,
		Name:     "Jane",
		Email:    "Jane@Test.com",
		Password: "Pass456",
		IsActive: false,
	}

	marshaled, err := original.MarshalJSON()
	require.NoError(t, err)

	var restored User
	err = restored.UnmarshalJSON(marshaled)
	require.NoError(t, err)

	require.Equal(t, "jane", restored.Name)
	require.Equal(t, "jane@test.com", restored.Email)
	require.Equal(t, "pass456", restored.Password)
}
