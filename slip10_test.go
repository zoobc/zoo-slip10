package slip10

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

var (
	mnemonic = "stand cheap entire summer claw subject victory supreme top divide tooth park change excite " +
		"legend category motor text zebra bottom mystery off garage energy"
)

func TestDeriveForPath(t *testing.T) {
	type args struct {
		path string
		seed []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Key
		wantErr bool
	}{
		{
			name: "wantSuccess",
			args: args{
				path: StellarPrimaryAccountPath,
				seed: NewSeed(mnemonic, DefaultPassword),
			},
			want: &Key{
				Key: []byte{176, 59, 0, 195, 59, 16, 248, 192, 96, 249, 116, 127, 137, 93, 227, 187, 175, 197, 112, 240,
					216, 57, 74, 203, 15, 161, 107, 3, 81, 213, 152, 159},
				ChainCode: []byte{44, 245, 19, 193, 34, 48, 73, 135, 55, 28, 150, 71, 184, 42, 167, 54, 27, 112, 145, 246,
					29, 59, 247, 160, 54, 9, 23, 73, 11, 144, 240, 206},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeriveForPath(tt.args.path, tt.args.seed)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeriveForPath() error = \n%v, wantErr \n%v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeriveForPath() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestKey_PublicKey(t *testing.T) {
	type fields struct {
		Key       []byte
		ChainCode []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "wantSuccess",
			fields: fields{
				Key: []byte{176, 59, 0, 195, 59, 16, 248, 192, 96, 249, 116, 127, 137, 93, 227, 187, 175, 197, 112, 240,
					216, 57, 74, 203, 15, 161, 107, 3, 81, 213, 152, 159},
				ChainCode: []byte{44, 245, 19, 193, 34, 48, 73, 135, 55, 28, 150, 71, 184, 42, 167, 54, 27, 112, 145, 246,
					29, 59, 247, 160, 54, 9, 23, 73, 11, 144, 240, 206},
			},
			want: []byte{137, 34, 109, 220, 127, 48, 20, 236, 229, 88, 171, 50, 253, 74, 4, 88, 95, 206, 140, 94, 168, 202,
				97, 204, 170, 216, 114, 80, 192, 202, 75, 44},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Key{
				Key:       tt.fields.Key,
				ChainCode: tt.fields.ChainCode,
			}
			got, err := k.PublicKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("Key.PublicKey() error = \n%v, wantErr \n%v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Key.PublicKey() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestNewMasterKey(t *testing.T) {
	type args struct {
		seed []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Key
		wantErr bool
	}{
		{
			name: "wantSuccess",
			args: args{
				seed: NewSeed(mnemonic, DefaultPassword),
			},
			want: &Key{
				Key: []byte{34, 39, 248, 143, 145, 132, 25, 71, 137, 113, 163, 220, 35, 182, 180, 224, 139, 14, 89, 203,
					244, 104, 209, 113, 199, 254, 132, 141, 182, 158, 180, 54},
				ChainCode: []byte{97, 61, 54, 96, 78, 109, 3, 236, 143, 229, 120, 78, 212, 82, 89, 206, 204, 124, 130, 71,
					229, 4, 103, 63, 4, 125, 142, 214, 5, 199, 35, 255},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMasterKey(tt.args.seed)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMasterKey() error = \n%v, wantErr \n%v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMasterKey() got = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestNewSeed(t *testing.T) {
	type args struct {
		mnemonic string
		password string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "wantSuccess",
			args: args{
				mnemonic: mnemonic,
				password: DefaultPassword,
			},
			want: []byte{192, 166, 232, 226, 38, 129, 226, 64, 251, 106, 248, 138, 3, 237, 155, 156, 250, 183, 211, 81, 69,
				245, 156, 225, 229, 120, 210, 20, 232, 99, 130, 14, 68, 176, 193, 46, 54, 136, 18, 159, 114, 53, 234, 105,
				114, 177, 239, 111, 56, 23, 86, 81, 124, 103, 3, 128, 46, 116, 71, 154, 30, 165, 231, 246},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSeed(tt.args.mnemonic, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_Serialize(t *testing.T) {
	type fields struct {
		Key       []byte
		ChainCode []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "wantSuccess",
			fields: fields{
				Key: []byte{176, 59, 0, 195, 59, 16, 248, 192, 96, 249, 116, 127, 137, 93, 227, 187, 175, 197, 112, 240,
					216, 57, 74, 203, 15, 161, 107, 3, 81, 213, 152, 159},
				ChainCode: []byte{44, 245, 19, 193, 34, 48, 73, 135, 55, 28, 150, 71, 184, 42, 167, 54, 27, 112, 145, 246,
					29, 59, 247, 160, 54, 9, 23, 73, 11, 144, 240, 206},
			},
			want: "iSJt3H8wFOzlWKsy_UoEWF_OjF6oymHMqthyUMDKSyxb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Key{
				Key:       tt.fields.Key,
				ChainCode: tt.fields.ChainCode,
			}
			if got := k.Serialize(); got != tt.want {
				t.Errorf("Key.Serialize() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestKey_Sign(t *testing.T) {
	type args struct {
		Seed    []byte
		Payload []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "wantSuccess",
			args: args{
				Seed:    NewSeed(mnemonic, DefaultPassword),
				Payload: []byte{1, 1, 1, 1, 1},
			},
			want: []byte{0, 0, 0, 0, 110, 66, 70, 93, 92, 101, 4, 194, 85, 144, 52, 236, 241, 61, 16, 8, 182, 241, 85, 95, 155, 148, 34, 146, 109, 71, 83, 8, 218, 224, 74, 163, 171, 97, 1, 114,
				162, 241, 120, 90, 99, 133, 171, 201, 212, 234, 130, 199, 223, 82, 1, 241, 156, 219, 212, 204, 5, 179, 161, 18, 238, 181, 212, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			masterKey, _ := DeriveForPath("m/44'/148'/0'", tt.args.Seed)
			fmt.Println("")
			log.Println("masterkey derivedPath Result = ", masterKey.Key)
			k := &Key{
				Key: masterKey.Key,
			}
			if got := k.Sign(tt.args.Payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sign() = %v, \n want %v", got, tt.want)
			}
		})
	}
}
