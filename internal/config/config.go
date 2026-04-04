package config

import (
"fmt"
"os"
"path/filepath"
"time"

"gopkg.in/yaml.v2"
"venera/internal/constants"
)

// Config represents the application configuration
type Config struct {
// Repository settings
Repository RepositoryConfig `yaml:"repository"`

// HTTP settings
HTTP HTTPConfig `yaml:"http"`

// Paths
Paths PathsConfig `yaml:"paths"`

// Security settings
Security SecurityConfig `yaml:"security"`

// Logging settings
Logging LoggingConfig `yaml:"logging"`
}

type RepositoryConfig struct {
URL         string `yaml:"url"`
SignatureURL string `yaml:"signature_url"`
}

type HTTPConfig struct {
DialTimeout           time.Duration `yaml:"dial_timeout"`
TLSHandshakeTimeout   time.Duration `yaml:"tls_handshake_timeout"`
ResponseHeaderTimeout time.Duration `yaml:"response_header_timeout"`
ExpectContinueTimeout time.Duration `yaml:"expect_continue_timeout"`
KeepAlive             time.Duration `yaml:"keep_alive"`
RequestTimeout        time.Duration `yaml:"request_timeout"`
}

type PathsConfig struct {
VeneraDir    string `yaml:"venera_dir"`
DatabaseFile string `yaml:"database_file"`
LogFile      string `yaml:"log_file"`
}

type SecurityConfig struct {
MaxFileReadSize      int64  `yaml:"max_file_read_size"`
ScriptDirPermissions uint32 `yaml:"script_dir_permissions"`
VeneraDirPermissions uint32 `yaml:"venera_dir_permissions"`
LogFilePermissions   uint32 `yaml:"log_file_permissions"`
}

type LoggingConfig struct {
Level      string `yaml:"level"`
Structured bool   `yaml:"structured"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
return &Config{
Repository: RepositoryConfig{
URL:         constants.DefaultRepoURL,
SignatureURL: constants.DefaultSignURL,
},
HTTP: HTTPConfig{
DialTimeout:           constants.HTTPDialTimeout,
TLSHandshakeTimeout:   constants.HTTPTLSHandshakeTimeout,
ResponseHeaderTimeout: constants.HTTPResponseHeaderTimeout,
ExpectContinueTimeout: constants.HTTPExpectContinueTimeout,
KeepAlive:             constants.HTTPKeepAlive,
RequestTimeout:        30 * time.Second,
},
Paths: PathsConfig{
VeneraDir:    constants.VeneraDirName,
DatabaseFile: constants.DatabaseFileName,
LogFile:      constants.LogFileName,
},
Security: SecurityConfig{
MaxFileReadSize:      constants.MaxLuaFileReadSize,
ScriptDirPermissions: uint32(constants.ScriptDirPermissions),
VeneraDirPermissions: uint32(constants.VeneraDirPermissions),
LogFilePermissions:   uint32(constants.LogFilePermissions),
},
Logging: LoggingConfig{
Level:      "info",
Structured: true,
},
}
}

// LoadConfig loads configuration from a YAML file
// If the file doesn't exist, returns default configuration
func LoadConfig(configPath string) (*Config, error) {
// If config file doesn't exist, return defaults
if _, err := os.Stat(configPath); os.IsNotExist(err) {
return DefaultConfig(), nil
}

data, err := os.ReadFile(configPath)
if err != nil {
return nil, fmt.Errorf("failed to read config file: %w", err)
}

cfg := DefaultConfig()
if err := yaml.Unmarshal(data, cfg); err != nil {
return nil, fmt.Errorf("failed to parse config file: %w", err)
}

return cfg, nil
}

// SaveConfig saves the configuration to a YAML file
func SaveConfig(cfg *Config, configPath string) error {
// Ensure directory exists
dir := filepath.Dir(configPath)
if err := os.MkdirAll(dir, 0750); err != nil {
return fmt.Errorf("failed to create config directory: %w", err)
}

data, err := yaml.Marshal(cfg)
if err != nil {
return fmt.Errorf("failed to marshal config: %w", err)
}

if err := os.WriteFile(configPath, data, 0640); err != nil {
return fmt.Errorf("failed to write config file: %w", err)
}

return nil
}

// GetConfigPath returns the default config file path
func GetConfigPath(homeDir string) string {
return filepath.Join(homeDir, constants.VeneraDirName, "config.yaml")
}
