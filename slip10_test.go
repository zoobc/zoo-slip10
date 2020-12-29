// ZooBC zoo-slip10
//
// Copyright © 2020 Quasisoft Limited - Hong Kong
//
// ZooBC is architected by Roberto Capodieci & Barton Johnston
//             contact us at roberto.capodieci[at]blockchainzoo.com
//             and barton.johnston[at]blockchainzoo.com
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package slip10

import (
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
				path: ZoobcPrimaryAccountPath,
				seed: NewSeed(mnemonic, DefaultPassword),
			},
			want: &Key{
				Key: []byte{1, 31, 185, 144, 13, 153, 44, 94, 226, 37, 5, 72, 123, 69, 118, 251, 77, 47, 125, 171, 240,
					9, 27, 82, 96, 223, 0, 209, 103, 177, 204, 8},
				ChainCode: []byte{214, 57, 48, 7, 176, 194, 128, 90, 181, 245, 31, 121, 191, 238, 64, 189, 31, 224, 247,
					225, 240, 213, 206, 52, 211, 36, 198, 243, 207, 98, 251, 58},
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
				Key: []byte{2, 198, 12, 49, 184, 230, 187, 32, 18, 64, 25, 217, 29, 173, 221, 228, 139, 149, 111,
					120, 255, 66, 46, 139, 32, 86, 37, 108, 129, 168, 207, 189},
				ChainCode: []byte{124, 114, 210, 214, 229, 31, 51, 217, 212, 112, 13, 208, 243, 26, 124, 58, 83,
					71, 186, 43, 206, 123, 46, 183, 8, 250, 114, 126, 64, 66, 157, 102},
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
			want: []byte{173, 197, 58, 226, 61, 78, 249, 72, 169, 74, 85, 145, 56, 101, 223, 135, 130, 71, 151, 107,
				42, 137, 208, 201, 33, 226, 189, 201, 186, 7, 125, 229, 106, 177, 38, 122, 13, 107, 38, 191, 244, 178,
				136, 15, 169, 244, 235, 72, 154, 75, 95, 135, 10, 214, 228, 206, 35, 89, 122, 13, 8, 8, 91, 16},
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
			want: []byte{0, 0, 0, 0, 53, 5, 176, 196, 177, 144, 20, 15, 20, 181, 199, 155, 139, 120, 180, 250, 62, 253, 98,
				26, 203, 206, 47, 160, 239, 48, 58, 189, 183, 48, 16, 127, 237, 107, 43, 12, 192, 25, 106, 111, 172, 213,
				155, 57, 40, 210, 130, 46, 120, 207, 203, 175, 43, 40, 205, 156, 19, 4, 122, 161, 109, 142, 252, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			masterKey, _ := DeriveForPath("m/44'/883'/0'", tt.args.Seed)
			k := &Key{
				Key: masterKey.Key,
			}
			if got := k.Sign(tt.args.Payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sign() = %v, \n want %v", got, tt.want)
			}
		})
	}
}
