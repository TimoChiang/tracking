package tracking

import (
	"testing"
)

func TestTracking_SetCompany(t *testing.T) {
	tests := []struct {
		name string
		company string
		isErr bool
	}{
		{"Set valid company", string(Yamato), false},
		{"Set invalid company", "noExist", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tk := New()
			err := tk.SetCompany(test.company)
			var isErr bool
			if err != nil {
				isErr = true
			}
			if test.isErr != isErr {
				t.Errorf("expected isErr = %v, got %v", test.isErr, isErr)
			}
		})
	}
}

func TestTracking_SetNumber(t *testing.T) {
	tests := []struct {
		name string
		number string
	}{
		{"Set valid number", "123456"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tk := New()
			tk.SetNumber(test.number)
			if tk.Number != test.number {
				t.Errorf("expected Number is %v, got %v", test.number, tk.Number)
			}
		})
	}
}

func TestTracking_Request(t *testing.T) {
	tests := []struct {
		name string
		company Company
		number string
	}{
		{"Yamato", Yamato , "467472294192"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tk := New()
			tk.SetCompany(string(test.company))
			tk.SetNumber(test.number)
			if err := tk.Request(); err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if len(tk.StatusList) == 0 {
				t.Errorf("expected status more then 0, got 0")
			}
		})
	}
}