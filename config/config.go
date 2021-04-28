package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

var cnfg Configuration

// Configuration modelo que contendrá un mapa de
// string/interface para leer y almacenar la data
// del archivo de configuración tipo json
// Código base obtenido de https://github.com/alexyslozada/config-go/
type Configuration struct {
	data map[string]interface{}
}

func New(fullpath string) (*Configuration, error) {
	b, err := loadFile(fullpath)
	if err != nil {
		return nil, err
	}

	if strings.Contains(fullpath, ".json") {
		err = loadBytes(b, &cnfg)
	} else if strings.Contains(fullpath, ".yml") || strings.Contains(fullpath, ".yaml") {
		err = loadBytesYaml(b, &cnfg)
	}

	if err != nil {
		return nil, err
	}

	return &cnfg, nil
}

func loadFile(fullpath string) ([]byte, error) {
	f, err := ioutil.ReadFile(fullpath)
	if err != nil {
		return f, err
	}

	return f, nil
}

func loadBytes(d []byte, c *Configuration) error {
	err := json.Unmarshal(d, &c.data)
	if err != nil {
		return err
	}

	return nil
}

func loadBytesYaml(d []byte, c *Configuration) error {
	fmt.Println(string(d))
	err := yaml.Unmarshal(d, &c.data)
	if err != nil {
		return err
	}

	fmt.Println("Data", c.data)
	return nil
}

func (c *Configuration) Validate(names ...string) error {
	for _, v := range names {
		_, ok := c.data[v]
		if !ok {
			return errors.New(fmt.Sprintf("no existe el campo %s", v))
		}
	}

	return nil
}

// Get devuelve el valor del campo si existe, tipo string
func (c *Configuration) Get(name string) (string, error) {
	v, ok := c.data[name].(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("no existe el campo %s o no se puede convertir en string", name))
	}

	return v, nil
}

// GetInt devuelve el valor del campo si existe, tipo int
func (c *Configuration) GetInt(name string) (int, error) {
	v, ok := c.data[name].(float64)
	if !ok {
		return 0, errors.New(fmt.Sprintf("no existe el campo %s o no se puede convertir en int", name))
	}

	return int(v), nil
}

// GetFloat devuelve el valor del campo si existe, tipo float64
func (c *Configuration) GetFloat(name string) (float64, error) {
	v, ok := c.data[name].(float64)
	if !ok {
		return 0, errors.New(fmt.Sprintf("no existe el campo %s o no se puede convertir en float64", name))
	}

	return v, nil
}

// GetBool devuelve el valor del campo si existe, tipo bool
func (c *Configuration) GetBool(name string) (bool, error) {
	v, ok := c.data[name].(bool)
	if !ok {
		return false, errors.New(fmt.Sprintf("no existe el campo %s o no se puede convertir en bool", name))
	}

	return v, nil
}

// GetSliceString devuelve el valor del campo si existe, tipo []string
func (c *Configuration) GetSliceString(name string) ([]string, error) {
	ss := make([]string, 0)

	vs, ok := c.data[name].([]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("no existe el campo %s", name))
	}

	for _, v := range vs {
		s, ok := v.(string)
		if !ok {
			return nil, errors.New(fmt.Sprintf("el campo %s no se puede convertir en []string", name))
		}

		ss = append(ss, s)
	}

	return ss, nil
}
