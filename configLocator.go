package config

import (
	fileUtil "github.com/flightx31/file"
	"github.com/spf13/afero"
	"os"
	"os/user"
	"path/filepath"
)

var Fs afero.Fs

func SetFs(fs afero.Fs) {
	Fs = fs
	fileUtil.SetFs(fs)
}

func LocateConfigFile(configFileName string, folderName string) (configFilePath string, found bool, err error) {
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", false, err
	}

	// Try current directory first
	configFilePath = filepath.Join(currentDir, configFileName)
	configFileExists := fileUtil.ExistsAndIsFile(configFilePath)

	if configFileExists {
		return configFilePath, true, nil
	}

	// Try current directory in folder second
	configFilePath = filepath.Join(currentDir, folderName, configFileName)
	configFileExists = fileUtil.ExistsAndIsFile(configFilePath)

	if configFileExists {
		return configFilePath, true, nil
	}

	// Finally try users home directory
	usr, err := user.Current()
	if err != nil {
		return "", false, err
	}
	homeDirectory := usr.HomeDir
	configFilePath = filepath.Join(homeDirectory, folderName, configFileName)
	configFileExists = fileUtil.ExistsAndIsFile(configFilePath)

	if configFileExists {
		return configFilePath, true, nil
	}

	return "", false, err
}
