package test

// import (
// 	"regexp"
// 	"testing"
// )

// func TestAddUser(t *testing.T) {
// 	type TestCase struct {
// 		Name     string
// 		Email    string
// 		password string
// 	}

// 	tests := []TestCase{
// 		{normWord: "JUMP", loweredWord: "jump"},
// 		{normWord: "Alphabet", loweredWord: "alphabet"},
// 		{normWord: "cRaZy", loweredWord: "crazy"},
// 	}

// 	for _, test := range tests {
// 		value := test.normWord
// 		want := regexp.MustCompile(test.loweredWord)
// 		msg, err := go_reloaded.LowerCase(value)
// 		if !want.MatchString(msg) || err != nil {
// 			t.Errorf(`%v = %q, %v, want match for %#q`, value, msg, err, want)
// 		}
// 	}
// }
