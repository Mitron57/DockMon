package config

import (
    "gopkg.in/yaml.v3"
    "os"
)

type Config struct {
    App app `yaml:"app"`
}

type app struct {
    Host   string   `yaml:"host"`
    Port   string   `yaml:"port"`
    Period int      `yaml:"period"`
    Db     dbConfig `yaml:"db"`
}

type dbConfig struct {
    Host           string `yaml:"host"`
    Port           string `yaml:"port"`
    UserEnvKey     string `yaml:"userEnvKey"`
    PasswordEnvKey string `yaml:"passwordEnvKey"`
    Dbname         string `yaml:"dbname"`
    Sslmode        string `yaml:"sslmode"`
}

func Parse(path string) (*Config, error) {
    var cfg Config
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}
