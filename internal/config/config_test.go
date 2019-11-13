package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	cases := []struct {
		name     string
		path     string
		wantData *Configuration
		wantErr  bool
	}{
		{
			name:    "Fail on non-existing file",
			path:    "notExists",
			wantErr: true,
		},
		{
			name:    "Fail on wrong file format",
			path:    "testdata/config.invalid.yaml",
			wantErr: true,
		},
		{
			name: "Success",
			path: "testdata/config.testdata.yaml",
			wantData: &Configuration{
				Server: &Server{
					Port: ":8989",
				},
				Database: &Database{
					Type: "mysql",
					PSN:  "psn://stringurl",
				},
			},
		},
		{
			name: "Mysql as default type",
			path: "testdata/config.test.default.database.yaml",
			wantData: &Configuration{
				Server: &Server{
					Port: ":8989",
				},
				Database: &Database{
					Type: "mysql",
					PSN:  "psn://stringurl",
				},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := Load(tt.path)
			assert.Equal(t, tt.wantData, cfg)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
