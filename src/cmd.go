package src

import (
	"fmt"
	"path"
)

func Push() {
	envConfig := parseEnvConfigFile("./environment.yaml")
	syncConfig := parseSyncConfigFile(envConfig.ConfigFilePath)

	var current Machine
	for _, machine := range syncConfig.Machines {
		if machine.Name == envConfig.Machine {
			current = machine
		}
	}

	// fmt.Println(current.Name)
	// fmt.Println(current.Maps)

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
	fmt.Println("")

	// for i, map := range current.Maps {
	// fmt.Sprintf("%s -> %s", map.repo, map.dest)
	// Copy(map.repo, map.dest)
	// }

}
