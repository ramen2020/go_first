package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

// ConfigList dbの接続情報
type ConfigList struct {
	DBMS     string
	USER     string
	PASS     string
	PROTOCOL string
	DBNAME   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		DBMS:     cfg.Section("db").Key("DBMS").String(),
		USER:     cfg.Section("db").Key("USER").String(),
		PASS:     cfg.Section("db").Key("PASS").String(),
		PROTOCOL: cfg.Section("db").Key("PROTOCOL").String(),
		DBNAME:   cfg.Section("db").Key("DBNAME").String(),
	}
}
