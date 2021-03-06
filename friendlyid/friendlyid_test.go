package friendlyid

import (
	"testing"
)

func Test_Encode(t *testing.T) {
	u, _ := Encode("796e28ae-f497-0143-f797-52169b36be94")

	if u != "3h8Pgh03y0pa2W6ltuUfZ6" {
		t.Error("Not valid")
	}
}

func Test_Decode(t *testing.T) {
	u, _ := Decode("3h8Pgh03y0pa2W6ltuUfZ6")

	if u != "796e28ae-f497-0143-f797-52169b36be94" {
		t.Error("Not valid")
	}
}

func Test_Wrong_ToUUID(t *testing.T) {
	u, _ := Decode("3h8Pgh0")

	if u != "00000000-0000-0000-0000-0030e021f32a" {
		t.Error("Not valid")
	}
}

func Test_ManyZeroes_UUID(t *testing.T) {
	ref := "00000000-0000-4000-8000-000000000000"
	u, err := Encode(ref)
	if u != "1VgEh72lXvTXkG" {
		t.Error("Not valid")
	}

	d, err := Decode(u)
	if err != nil || d != ref {
		t.Error("Not valid")
	}
}

func Test_convertDown(t *testing.T) {
	type args struct {
		oldNumber    string
		baseAlphabet string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertDown(tt.args.oldNumber, tt.args.baseAlphabet)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertDown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertDown() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertUp(t *testing.T) {
	type args struct {
		oldNumber    string
		baseAlphabet string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertUp(tt.args.oldNumber, tt.args.baseAlphabet)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertUp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Decode1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encode1(t *testing.T) {
	type args struct {
		shex string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.args.shex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insert(t *testing.T) {
	type args struct {
		s   string
		sep string
		i   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.s, tt.args.sep, tt.args.i); got != tt.want {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}