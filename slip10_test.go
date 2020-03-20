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
				Seed: NewSeed(mnemonic, DefaultPassword),
				// Payload: []byte{2, 2, 0, 0, 1, 16, 1, 115, 94, 0, 0, 0, 0, 44, 0, 0, 0, 105, 83, 74, 116, 51, 72, 56, 119, 70, 79, 122, 108, 87, 75, 115, 121, 95, 85, 111, 69, 87, 70, 95, 79, 106, 70, 54, 111, 121, 109, 72, 77, 113, 116, 104, 121, 85, 77, 68, 75, 83, 121, 120, 98, 44, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 66, 15, 0, 0, 0, 0, 0, 32, 0, 0, 0, 91, 36, 228, 70, 101, 94, 186, 246, 186, 4, 78, 142, 173, 162, 187, 173, 202, 81, 243, 92, 141, 120, 148, 220, 41, 160, 208, 94, 174, 166, 62, 207, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				Payload: []byte{1, 1, 1, 1, 1},
			},
			want: []byte{0, 0, 0, 0, 226, 197, 43, 149, 140, 220, 160, 213, 179, 103, 200, 25, 104, 114, 186, 66, 171, 190, 168, 95, 185, 207, 31, 217, 190, 213, 73, 206, 120, 165, 216, 156, 22, 17, 232, 47, 233, 10, 19, 60, 229,
				255, 158, 59, 56, 154, 99, 27, 117, 9, 175, 233, 217, 61, 213, 16, 67, 172, 204, 122, 39, 43, 80, 1}},
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
