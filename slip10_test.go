package slip10

import (
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
				Seed: []byte{192, 166, 232, 226, 38, 129, 226, 64, 251, 106, 248, 138, 3, 237, 155, 156, 250, 183, 211, 81, 69, 245, 156, 225, 229, 120, 210, 20, 232, 99, 130, 14, 6,
					8, 176, 193, 46, 54, 136, 18, 159, 114, 53, 234, 105, 114, 177, 239, 111, 56, 23, 86, 81, 124, 103, 3, 128, 46, 116, 71, 154, 30, 165, 231, 246},
				Payload: []byte{2, 2, 0, 0, 1, 204, 54, 116, 94, 0, 0, 0, 0, 44, 0, 0, 0, 105, 83, 74, 116, 51, 72, 56, 119, 70, 79, 122, 108, 87, 75, 115, 121, 9,
					5, 85, 111, 69, 87, 70, 95, 79, 106, 70, 54, 111, 121, 109, 72, 77, 113, 116, 104, 121, 85, 77, 68, 75, 83, 121, 120, 98, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 44, 0, 0, 0, 87,
					121, 84, 107, 82, 109, 86, 101, 117, 118, 97, 54, 66, 69, 54, 79, 114, 97, 75, 55, 114, 99, 112, 82, 56, 49, 121, 78, 101, 74, 84, 99, 75, 97, 68, 81, 88, 113, 54, 10,
					9, 80, 115, 56, 61, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: []byte{0, 0, 0, 0, 59, 124, 226, 161, 217, 229, 41, 151, 237, 243, 135, 100, 219, 120, 91, 78, 136, 179, 93, 238, 41, 122, 139, 98, 154, 134, 124, 216, 222, 161, 187, 190, 221, 205, 128, 37, 207, 103, 49, 114, 188, 85, 59, 106, 11, 44, 200, 9, 130, 31, 126, 79, 0, 49, 37, 225, 203, 145, 127, 184, 249, 52, 55, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			masterKey, _ := DeriveForPath("m/44'/148'/0'", tt.args.Seed)
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
