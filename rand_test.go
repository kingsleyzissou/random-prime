package rand_test

import (
	"testing"

	rand "github.com/kingsleyzissou/rand-prime"
)

func TestFirst10PrimeNumbers(t *testing.T) {
	type fields struct {
		RandomInt int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test (2).IsPrime()",
			fields: fields{
				RandomInt: 2,
			},
			want: true,
		},
		{
			name: "Test (3).IsPrime()",
			fields: fields{
				RandomInt: 3,
			},
			want: true,
		},
		{
			name: "Test (7).IsPrime()",
			fields: fields{
				RandomInt: 7,
			},
			want: true,
		},
		{
			name: "Test (11).IsPrime()",
			fields: fields{
				RandomInt: 11,
			},
			want: true,
		},
		{
			name: "Test (13).IsPrime()",
			fields: fields{
				RandomInt: 13,
			},
			want: true,
		},
		{
			name: "Test (17).IsPrime()",
			fields: fields{
				RandomInt: 17,
			},
			want: true,
		},
		{
			name: "Test (19).IsPrime()",
			fields: fields{
				RandomInt: 19,
			},
			want: true,
		},
		{
			name: "Test (23).IsPrime()",
			fields: fields{
				RandomInt: 23,
			},
			want: true,
		},
		{
			name: "Test (29).IsPrime()",
			fields: fields{
				RandomInt: 29,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rand.RandomInt(tt.fields.RandomInt)
			if got := r.IsPrime(); got != tt.want {
				t.Errorf("RandomInt.IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
