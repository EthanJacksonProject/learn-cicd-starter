package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyNoAuth(t *testing.T) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	// No Auth
	res, err := GetAPIKey(headers)

	ex_res := ""
	ex_err := ErrNoAuthHeaderIncluded

	if res != ex_res {
		t.Errorf("expected no string return but got one")
	}

	if err != ex_err {
		t.Errorf("expected an error but got none")
	}
}

func TestGetApiKeyValid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "ApiKey 123")

	res, err := GetAPIKey(headers)

	ex_res := "123"

	if res != ex_res {
		t.Errorf("expected res but got wrong one")
	}

	if err != nil {
		t.Errorf("received unexpected err")
	}
}

// func TestSplit(t *testing.T) {
//     got := Split("a/b/c", "/")
//     want := []string{"a", "b", "c"}
//     if !reflect.DeepEqual(want, got) {
//          t.Fatalf("expected: %v, got: %v", want, got)
//     }
// }

// // GetAPIKey -
// func GetAPIKey(headers http.Header) (string, error) {
// 	authHeader := headers.Get("Authorization")
// 	if authHeader == "" {
// 		return "", ErrNoAuthHeaderIncluded
// 	}
// 	splitAuth := strings.Split(authHeader, " ")
// 	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
// 		return "", errors.New("malformed authorization header")
// 	}

// 	return splitAuth[1], nil
// }
