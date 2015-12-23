package config

/*
test.conf
USERNAME = jiazhoulvke
PORT     = 1984
VERSION  = 1.1
HOST     = localhost
*/

import (
	"testing"
)

const (
	TestFile string = "test.conf"
)

type Cfg struct {
	Host     string  `cfgname:"HOST"`
	DBName   string  `default:"go"`
	Username string  `cfgname:"USERNAME" default:"jiazhoulvke"`
	Port     int     `cfgname:"PORT" default:"80"`
	Version  float64 `cfgname:"VERSION" default:"1.0"`
}

func Test_GetInt(t *testing.T) {
	cfg, err := Parse(TestFile)
	if err != nil {
		t.Error(err)
	}
	port, err := cfg.GetInt("PORT")
	if err != nil {
		t.Error(err)
	}
	if port != 1984 {
		t.Error("GetInt出错")
	} else {
		t.Log("GetInt测试成功")
	}
}

func Test_IntDefault(t *testing.T) {
	cfg, err := Parse(TestFile)
	if err != nil {
		t.Error(err)
	}
	dv := cfg.IntDefault("NOKEY", 1234)
	if dv != 1234 {
		t.Error("IntDefault出错")
	} else {
		t.Log("IntDefault测试成功")
	}
}

func Test_GetInt64(t *testing.T) {
	cfg, err := Parse(TestFile)
	if err != nil {
		t.Error(err)
	}
	port, err := cfg.GetInt64("PORT")
	if err != nil {
		t.Error(err)
	}
	if port != 1984 {
		t.Error("GetInt64出错")
	} else {
		t.Log("GetInt64测试成功")
	}
}

func Test_Int64Default(t *testing.T) {
	cfg, err := Parse(TestFile)
	if err != nil {
		t.Error(err)
	}
	dv := cfg.Int64Default("NOKEY", 1234)
	if dv != 1234 {
		t.Error("IntDefault出错")
	} else {
		t.Log("IntDefault测试成功")
	}
}

func Test_GetFloat(t *testing.T) {
	cfg, err := Parse(TestFile)
	if err != nil {
		t.Error(err)
	}
	port, err := cfg.GetFloat("VERSION")
	if err != nil {
		t.Error(err)
	}
	if port != 1.1 {
		t.Error("GetFloat出错")
	} else {
		t.Log("GetFloat测试成功")
	}
}

func Test_FloatDefault(t *testing.T) {
	cfg, err := Parse(TestFile)
	if err != nil {
		t.Error(err)
	}
	dv := cfg.FloatDefault("NOKEY", 12.34)
	if dv != 12.34 {
		t.Error("IntDefault出错")
	} else {
		t.Log("IntDefault测试成功")
	}
}

func Test_GetString(t *testing.T) {
	cfg, err := Parse(TestFile)
	if err != nil {
		t.Error(err)
	}
	username, err := cfg.GetString("USERNAME")
	if err != nil {
		t.Error(err)
	}
	if username != "jiazhoulvke" {
		t.Error("GetString出错")
	} else {
		t.Log("GetFloat测试成功")
	}
}

func Test_StringDefault(t *testing.T) {
	cfg, err := Parse(TestFile)
	if err != nil {
		t.Error(err)
	}
	dv := cfg.StringDefault("NOKEY", "hello,world!")
	if dv != "hello,world!" {
		t.Error("StringDefault出错")
	} else {
		t.Log("StringDefault测试成功")
	}
}

func Test_Init(t *testing.T) {
	var cfg Cfg
	err := Init(TestFile, &cfg)
	if err != nil {
		t.Error(err)
	}
	if cfg.DBName != "go" || cfg.Username != "jiazhoulvke" || cfg.Port != 1984 || cfg.Version != 1.1 {
		t.Error("Init测试失败")
	} else {
		t.Log("Init测试成功")
	}
}