package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidLogin(t *testing.T) {
	require.True(t, ValidLogin("valera"))
	require.True(t, ValidLogin("valera_2011"))
	require.True(t, ValidLogin("ivan2011"))
	require.True(t, ValidLogin("HerabriN-1_23"))

	require.False(t, ValidLogin("валера"))
	require.False(t, ValidLogin("val"))
}

func TestValidName(t *testing.T) {
	require.True(t, ValidName("valera"))
	require.True(t, ValidName("valera_2011"))
	require.True(t, ValidName("na4alnik"))
	require.True(t, ValidName("HerabriN-1_23"))

	require.False(t, ValidName("валера"))
	require.False(t, ValidName("val"))
}

func TestValidPassword(t *testing.T) {
	require.True(t, ValidPassword("asdA12!as"))
	require.True(t, ValidPassword("ghjvjenth12Z!@"))

	require.False(t, ValidPassword("валера"))
	require.False(t, ValidPassword("val"))
}

func TestPasswordLever(t *testing.T) {
	require.Nil(t, CheckPasswordLever("asdA12!as"))
	require.Nil(t, CheckPasswordLever("asa123Z@1"))
	require.Nil(t, CheckPasswordLever("ghjvjenth12Z!@"))

	require.Error(t, CheckPasswordLever("валера"))
	require.Error(t, CheckPasswordLever("val"))
	require.Error(t, CheckPasswordLever("123123123"))
	require.Error(t, CheckPasswordLever("123123asd"))
	require.Error(t, CheckPasswordLever("123asdZXC"))

	require.Nil(t, CheckPasswordLever("123asdZXC!@"))
}

func TestValidateUserData(t *testing.T) {
	require.True(t, ValidateUserData(map[string]string{
		"login":    "ivan2011",
		"name":     "na4alnik",
		"password": "asa123Z@1",
	}))

	require.False(t, ValidateUserData(map[string]string{
		"login": "che",
		"name":  "chelovek",
	}))
	require.False(t, ValidateUserData(map[string]string{
		"login":    "che",
		"password": "asa23Z@1",
	}))
	require.False(t, ValidateUserData(map[string]string{
		"name":     "chelovek",
		"password": "asa23Z@1",
	}))

}

func TestFileExist(t *testing.T) {
	require.True(t, FileExist("utils.go"))
	require.True(t, FileExist("utils_test.go"))

	require.False(t, FileExist("main.go"))
}
