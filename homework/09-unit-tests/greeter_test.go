package greeter

import (
	"testing"
)

// 11 tests are required to check the states of the function
func TestGreet(t *testing.T) {
	type args struct {
		name string
		hour int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"check <hour> in range [0;5]; value:0", args{"buddy", 0}, "Good night Buddy!"},
		{"check <hour> in range [0;5]; value:5", args{"buddy", 5}, "Good night Buddy!"},
		{"check <hour> in range [22;23]; value:22", args{"buddy", 22}, "Good night Buddy!"},
		{"check <hour> in range [22;23]; value:23", args{"buddy", 23}, "Good night Buddy!"},
		{"check <hour> in range [6;11]; value:6", args{"buddy", 6}, "Good morning Buddy!"},
		{"check <hour> in range [6;11]; value:11", args{"buddy", 11}, "Good morning Buddy!"},
		{"check <hour> in range [12;17]; value:12", args{"buddy", 12}, "Hello Buddy!"},
		{"check <hour> in range [12;17]; value:17", args{"buddy", 17}, "Hello Buddy!"},
		{"check <hour> in range [18;21]; value:18", args{"buddy", 18}, "Good evening Buddy!"},
		{"check <hour> in range [18;21]; value:21", args{"buddy", 21}, "Good evening Buddy!"},
		{"check <hour> out of range [0;5]; value:-1", args{"buddy", -1}, "<hour>: range error!"},
		{"check <hour> out of range [22;23]; value:24", args{"buddy", 24}, "<hour>: range error!"},
		//{"check <hour> in range [min;max]; boundary_value: max", args{"buddy", math.MaxInt}, "<hour>: range error!"},
		//{"check <hour> in range [min;max]; boundary_value: min", args{"buddy", math.MinInt}, "<hour>: range error!"},
		//{"testing <hour> out of range [min;max]; boundary value: max+1, args{"buddy", math.MaxInt + 1}, "Error!"},
		//{"testing <hour> out of range [min;max]; boundary value: min-1, args{"buddy", math.MinInt - 1}, "Error!"},
		{"check <name> for single word", args{"johann", 13}, "Hello Johann!"},
		{"check <name> for two-word", args{"johann bach", 13}, "Hello Johann!"},
		{"check <name> for three-word", args{"johann sebastian bach", 13}, "Hello Johann!"},
		{"check <name> for hyphen", args{"mustapha-ibrahim", 13}, "Hello Mustapha-Ibrahim!"},
		{"check <name> for apostrophe", args{"d'artagnan", 13}, "Hello D'Artagnan!"},
		{"check <name> for spaces", args{"  tom     ", 13}, "Hello Tom!"},
		{"check <name> for uppercase", args{"WOO", 13}, "Hello WOO!"},
		{"check <name> for russia", args{"андрей", 13}, "Hello Андрей!"},
		{"check <name> for china", args{"晓明", 13}, "Hello 晓明!"},
		{"check <name> for numbers", args{"2pac", 13}, "Hello 2pac!"},
		{"check <name> for special characters", args{"Δ", 13}, "Hello Δ!"},
		{"check <name> for long", args{
			"晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明" +
				"晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明" +
				"晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明" +
				"晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明" +
				"", 13}, "Hello " +
			"晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明" +
			"晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明" +
			"晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明" +
			"晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明!"},
		{"check <name> for empty", args{"", 13}, "<name>: empty!"},
		{"check <name> for too long", args{
			"晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明" +
				"晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明" +
				"晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明" +
				"晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明晓明" +
				"", 13}, "<name>: too much!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.args.name, tt.args.hour); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
