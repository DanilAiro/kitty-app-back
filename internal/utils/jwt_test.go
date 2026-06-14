package utils

import (
	"testing"
)

func TestVerifyToken(tt *testing.T) {
	tests := []struct {
		Name     string
		Token    string
		Expected string
	}{
		{"Test 1", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3ODE0Nzg5MDYsImlhdCI6MTc4MTM5MjUwNiwic3ViIjoidGVzdEB0ZXN0LnJ1In0.zjgiI08SWTrrdIK3Cl2T3zZAWS9KNhVkaNMW57VVHLs", "test@test.ru"},
		{"Test 2", ".eyJleHAiOjE3ODE0Nzg5MDYsImlhdCI6MTc4MTM5MjUwNiwic3ViIjoidGVzdEB0ZXN0LnJ1In0.zjgiI08SWTrrdIK3Cl2T3zZAWS9KNhVkaNMW57VVHLs", ""},
		{"Test 3", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..zjgiI08SWTrrdIK3Cl2T3zZAWS9KNhVkaNMW57VVHLs", ""},
		{"Test 4", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3ODE0Nzg5MDYsImlhdCI6MTc4MTM5MjUwNiwic3ViIjoidGVzdEB0ZXN0LnJ1In0.", ""},
		{"Test 5", "", ""},
	}

	for _, test := range tests {
		tt.Run(test.Name, func(t *testing.T) {
			got, err := VerifyToken(test.Token)

			if got["sub"] == test.Expected || err != nil && test.Expected == "" {
				return
			}

			if err != nil && test.Expected != "" {
				t.Error("Ошибка " + err.Error())
				return
			}

			if got["sub"] != test.Expected {
				t.Errorf("Ожидали %s, получили %s", test.Expected, got["sub"])
			}
		})
	}
}

func TestCreateToken(tt *testing.T) {
	tests := []struct {
		Name     string
		Email    string
	}{
		{"Test 1", "test@test.ru"},
		{"Test 2", "test1@test1.ru"},
		{"Test 3", "test2@test2.ru"},
		{"Test 4", "test3@test3.ru"},
	}

	for _, test := range tests {
		tt.Run(test.Name, func(t *testing.T) {
			got, err := CreateToken(test.Email)
			if err != nil {
				t.Error("Ошибка " + err.Error())
				return
			}

			res, err := VerifyToken(got)

			if res["sub"] != test.Email {
				t.Errorf("Ожидали %s, получили %s", test.Email, res["sub"])
			}
		})
	}
}
