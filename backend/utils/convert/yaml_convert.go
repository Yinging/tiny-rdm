package convutil

import (
	"gopkg.in/yaml.v3"
	"log"
)

type YamlConvert struct{}

func (YamlConvert) Encode(str string) (string, bool) {
	return str, true
}

func (YamlConvert) Decode(str string) (string, bool) {
	var obj map[string]any
	err := yaml.Unmarshal([]byte(str), &obj)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(obj)
	}
	return str, err == nil
}