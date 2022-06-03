package src

import (
	"fmt"
	"path"
)

func setup() (EnvConfig, SyncConfig, Machine) {
	envConfig := parseEnvConfigFile("./environment.yaml")
	syncConfig := parseSyncConfigFile(envConfig.ConfigFilePath)

	var current Machine
	for _, machine := range syncConfig.Machines {
		if machine.Name == envConfig.Machine {
			current = machine
		}
	}
	fmt.Printf("machine name: %s\n", current.Name)
	return envConfig, syncConfig, current
}

func Push() {
	envConfig, _, current := setup()

	for _, m := range current.Maps {
		// pwd, _ := os.Getwd()
		// fmt.Println(envConfig.FileDirPath)
		src := path.Join(envConfig.FileDirPath, m.Repo)
		fmt.Printf("%s -> %s\n", src, m.Dest)

		_, err := Copy(src, m.Dest)
		if err != nil {
			panic(err)
		}
	}
}

func Pull() {
	envConfig, _, current := setup()

	for _, m := range current.Maps {
		// pwd, _ := os.Getwd()
		// fmt.Println(envConfig.FileDirPath)
		src := path.Join(envConfig.FileDirPath, m.Repo)
		fmt.Printf("%s -> %s\n", m.Dest, src)

		_, err := Copy(m.Dest, src)
		if err != nil {
			fmt.Println(err)
		}
	}
}
