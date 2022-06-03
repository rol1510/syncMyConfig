package src

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type SyncConfig struct {
	Machines []Machine `yaml:"machines"`
}

type Map struct {
	Repo string `yaml:"repo"`
	Dest string `yaml:"dest"`
}

type Machine struct {
	Name string `yaml:"name"`
	Maps []Map  `yaml:"maps,omitempty"`
}

type EnvConfig struct {
	Machine        string `yaml:"machine"`
	ConfigFilePath string `yaml:"configFile"`
	FileDirPath    string `yaml:"fileDir"`
}

func readEntireFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func parseEnvConfigFile(path string) EnvConfig {
	data := readEntireFile(path)
	envConf := EnvConfig{}

	err := yaml.Unmarshal(data, &envConf)
	if err != nil {
		panic(err)
	}

	return envConf
}

func parseSyncConfigFile(path string) SyncConfig {
	data := readEntireFile(path)
	syncConf := SyncConfig{}

	err := yaml.Unmarshal(data, &syncConf)
	if err != nil {
		panic(err)
	}

	return syncConf
}

// func GetSyncConfigFile() SyncConfig {
// 	env := parseEnvConfigFile("./environment.yaml")
// 	return parseSyncConfigFile(env.ConfigFilePath)
// }

func Do() {
	envConfig := parseEnvConfigFile("./environment.yaml")
	syncConfig := parseSyncConfigFile(envConfig.ConfigFilePath)

	fmt.Println(envConfig.ConfigFilePath)
	fmt.Println(envConfig.Machine)
	fmt.Println("---")

	fmt.Println(syncConfig.Machines)
	fmt.Println(syncConfig.Machines[0].Name)
	fmt.Println(syncConfig.Machines[0].Maps)
	fmt.Println(len(syncConfig.Machines))
}
