package config

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/notaryproject/notation-go/dir"
)

const (
	signingKeysPath = "./testdata/signingkeys.json"
)

var sampleSigningKeysInfo = &SigningKeys{
	Default: "wabbit-networks",
	Keys: []KeySuite{
		{
			Name: "wabbit-networks",
			X509KeyPair: &X509KeyPair{
				KeyPath:         "/home/demo/.config/notation/localkeys/wabbit-networks.key",
				CertificatePath: "/home/demo/.config/notation/localkeys/wabbit-networks.crt",
			},
		},
		{
			Name: "import.acme-rockets",
			X509KeyPair: &X509KeyPair{
				KeyPath:         "/home/demo/.config/notation/localkeys/import.acme-rockets.key",
				CertificatePath: "/home/demo/.config/notation/localkeys/import.acme-rockets.crt",
			},
		},
		{
			Name: "external-key",
			ExternalKey: &ExternalKey{

				ID:         "id1",
				PluginName: "pluginX",
				PluginConfig: map[string]string{
					"key": "value",
				},
			},
		},
	},
}

func TestLoadSigningKeysInfo(t *testing.T) {
	t.Cleanup(func() {
		// restore path
		SigningKeysPath = dir.Path.SigningKeyConfig()
	})
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want *SigningKeys
	}{
		{
			name: "read signingkeys info",
			args: args{filePath: signingKeysPath},
			want: sampleSigningKeysInfo,
		},
		{
			name: "get default signingkeys info",
			args: args{filePath: nonexistentPath},
			want: NewSigningKeys(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SigningKeysPath = tt.args.filePath
			got, err := LoadSigningKeys()
			if err != nil {
				t.Errorf("LoadSigningKeysInfo() error = %v", err)
				return
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatal("singingKeysInfo test failed.")
			}
		})
	}
}

func TestSaveSigningKeys(t *testing.T) {
	t.Cleanup(func() {
		// restore path
		SigningKeysPath = dir.Path.SigningKeyConfig()
	})
	root := t.TempDir()
	SigningKeysPath = filepath.Join(root, "signingkeys.json")
	sampleSigningKeysInfo.Save()
	info, err := LoadSigningKeys()
	if err != nil {
		t.Fatal("Load signingkeys.json from temp dir failed.")
	}
	if !reflect.DeepEqual(sampleSigningKeysInfo, info) {
		t.Fatal("Save signingkeys.json failed.")
	}
}
