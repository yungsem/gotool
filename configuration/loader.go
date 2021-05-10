package configuration

import "gopkg.in/yaml.v2"

type Loader interface {
	Load() map[string][]byte
}

// LoadConfig 解析配置文件的内容
func LoadConfig(loader Loader) *Conf {
	contentMap := loader.Load()

	e := env{}
	envContent := contentMap[ContentMapKeyEnv]
	err := yaml.Unmarshal(envContent, &e)
	if err != nil {
		panic("配置文件env加载失败")
	}

	var content []byte
	switch e.Env {
	case envLocal:
		content = contentMap[ContentMapKeyLocal]
	case envProd:
		content = contentMap[ContentMapKeyProd]
	}

	// 定义一个 conf
	c := Conf{}

	// 解析 yml
	err = yaml.Unmarshal(content, &c)
	if err != nil {
		panic("配置文件内容加载失败")
	}

	return &c
}
