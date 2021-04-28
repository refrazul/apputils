package config

import (
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New("config.json.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}
}

func TestConfiguration_Validate(t *testing.T) {
	c, err := New("config.json.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}

	err = c.Validate("db_server", "db_port", "db_user", "db_password", "is_secure")
	if err != nil {
		t.Error(err)
	}
}

func TestConfiguration_Get(t *testing.T) {
	c, err := New("config.json.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}

	_, err = c.Get("db_server")
	if err != nil {
		t.Error(err)
	}
}

func TestConfiguration_GetInt(t *testing.T) {
	c, err := New("config.json.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}

	_, err = c.GetInt("db_port")
	if err != nil {
		t.Error(err)
	}
}

func TestConfiguration_GetFloat(t *testing.T) {
	c, err := New("config.json.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}

	_, err = c.GetFloat("db_port")
	if err != nil {
		t.Error(err)
	}
}

func TestConfiguration_GetBool(t *testing.T) {
	c, err := New("config.json.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}

	_, err = c.GetBool("is_secure")
	if err != nil {
		t.Error(err)
	}
}

func TestConfiguration_Validate_Yaml(t *testing.T) {
	c, err := New("config.yml.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}

	err = c.Validate("db_server", "db_port", "db_user", "db_password", "is_secure")
	if err != nil {
		t.Error(err)
	}
}
