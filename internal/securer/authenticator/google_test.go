package authenticator

import (
	"reflect"
	"testing"
)

func Test_getGoogleUserInfo(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "22", args: args{
			t: "eyJhbGciOiJSUzI1NiIsImtpZCI6ImY0MTk2YWVlMTE5ZmUyMTU5M2Q0OGJmY2ZiNWJmMDAxNzdkZDRhNGQiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiNjM0ODk0MjIzNDczLTN0Z21ubzA3YzF0cDVjdWNuMTJtNWNmZnVnNmRpYjk1LmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwiYXVkIjoiNjM0ODk0MjIzNDczLTN0Z21ubzA3YzF0cDVjdWNuMTJtNWNmZnVnNmRpYjk1LmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwic3ViIjoiMTA0NzAyOTUzNTAyNzYyNzc4MzIzIiwiZW1haWwiOiJhcnlrYWxpbkBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiYXRfaGFzaCI6IldBc1pud3BON2ZnTE5TanZHZXBBYnciLCJuYW1lIjoiQWxleGFuZGVyIFJ5a2FsaW4iLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EtL0FPaDE0R2d3TzlVTkp5N1BGYnh3aDk4WHBFTEtRRks1d29xV2tLX2U3TklwZEdFPXM5Ni1jIiwiZ2l2ZW5fbmFtZSI6IkFsZXhhbmRlciIsImZhbWlseV9uYW1lIjoiUnlrYWxpbiIsImxvY2FsZSI6InJ1IiwiaWF0IjoxNjMzNzA0MTA4LCJleHAiOjE2MzM3MDc3MDgsImp0aSI6IjhjNDJiNTM3ZGUzNTE3N2Y2ZWVlMTVjMGU0YjMzMWEzZGQ5YjhlNjkifQ.nQNcYN65ykgcub6aI6XvdMuTKhT1QggFzHu9hlk0RkaMB6PE7Igr5WIkMzmoCCtUBf7htKUZKUy0CUVoVXu7_intmx5GzLXIwWqRpfjgYYd-FLZq4vErEkmDavz_G45zwIZ1jIw9n6zWya7mD09l9M9fte67oqlyFe0DHAWHt04rzSo3RY8vNXRKHbV94A3QIGUBxV4G8ZcHXOk-XtxDE9T9uJMmK7J4AUE7UeZi4CV9jTr95XYoebKzWF1em1PsERfv9rdWK8IJEi9yL7acnOr0ISE1_iuLdX-0c2ALtE11j-bCJtTq9DtQ-Mf_wDLDGs7rc_fGdYNSADe6V7jkbA",
		}, want: "arykalin@gmail.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGoogleUserInfo(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("getGoogleUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got\n: %+v", got)
			if !reflect.DeepEqual(got.Email, tt.want) {
				t.Errorf("getGoogleUserInfo() got = %v, want %v", got.Email, tt.want)
			}
		})
	}
}

func TestValidateGoogleJWT(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "nnn", args: args{
			tokenString: "eyJhbGciOiJSUzI1NiIsImtpZCI6ImY0MTk2YWVlMTE5ZmUyMTU5M2Q0OGJmY2ZiNWJmMDAxNzdkZDRhNGQiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiNjM0ODk0MjIzNDczLTN0Z21ubzA3YzF0cDVjdWNuMTJtNWNmZnVnNmRpYjk1LmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwiYXVkIjoiNjM0ODk0MjIzNDczLTN0Z21ubzA3YzF0cDVjdWNuMTJtNWNmZnVnNmRpYjk1LmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwic3ViIjoiMTA0NzAyOTUzNTAyNzYyNzc4MzIzIiwiZW1haWwiOiJhcnlrYWxpbkBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiYXRfaGFzaCI6IldBc1pud3BON2ZnTE5TanZHZXBBYnciLCJuYW1lIjoiQWxleGFuZGVyIFJ5a2FsaW4iLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EtL0FPaDE0R2d3TzlVTkp5N1BGYnh3aDk4WHBFTEtRRks1d29xV2tLX2U3TklwZEdFPXM5Ni1jIiwiZ2l2ZW5fbmFtZSI6IkFsZXhhbmRlciIsImZhbWlseV9uYW1lIjoiUnlrYWxpbiIsImxvY2FsZSI6InJ1IiwiaWF0IjoxNjMzNzA0MTA4LCJleHAiOjE2MzM3MDc3MDgsImp0aSI6IjhjNDJiNTM3ZGUzNTE3N2Y2ZWVlMTVjMGU0YjMzMWEzZGQ5YjhlNjkifQ.nQNcYN65ykgcub6aI6XvdMuTKhT1QggFzHu9hlk0RkaMB6PE7Igr5WIkMzmoCCtUBf7htKUZKUy0CUVoVXu7_intmx5GzLXIwWqRpfjgYYd-FLZq4vErEkmDavz_G45zwIZ1jIw9n6zWya7mD09l9M9fte67oqlyFe0DHAWHt04rzSo3RY8vNXRKHbV94A3QIGUBxV4G8ZcHXOk-XtxDE9T9uJMmK7J4AUE7UeZi4CV9jTr95XYoebKzWF1em1PsERfv9rdWK8IJEi9yL7acnOr0ISE1_iuLdX-0c2ALtE11j-bCJtTq9DtQ-Mf_wDLDGs7rc_fGdYNSADe6V7jkbA",
		}, want: "arykalin@gmail.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateGoogleJWT(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateGoogleJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateGoogleJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}
