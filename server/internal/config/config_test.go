package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigFromEnv(t *testing.T) {
	return
	type args struct {
		envs map[string]string
	}

	type wants struct {
		config *Config
	}

	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "Test OK",
			args: args{
				envs: map[string]string{
					"SERVER_ADDRESS": "http://localhost:8080",
				},
			},
			wants: wants{
				config: func() *Config {
					ret := &Config{}
					ret.Server.Address = "http://localhost:8080"
					return ret
				}(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnvVars(tt.args.envs)

			config := &Config{}
			err := config.Load()
			assert.NoError(t, err)

			assert.Equal(t, tt.wants.config, config)
		})
	}
}

func setEnvVars(envs map[string]string) {
	for key, value := range envs {
		os.Setenv(key, value)
	}
}
