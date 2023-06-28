package config

import (
	"fmt"
	fileUtil "github.com/flightx31/file"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestLocatePropertiesFile_no_config(t *testing.T) {
	fileName := "f.json"
	directory := "x"
	Fs = afero.NewMemMapFs()
	SetFs(Fs)
	currentDir, homeDir := fileUtil.CreateMockFileSystem(t, Fs)
	configFilePath, found, err := LocateConfigFile(fileName, directory)
	assert.True(t, len(configFilePath) == 0)
	assert.Nil(t, err)
	assert.False(t, found)
	fmt.Print(currentDir, homeDir)
}

func TestLocatePropertiesFile_config_current_directory(t *testing.T) {

	fileName := "f.json"
	directory := "x"
	Fs = afero.NewMemMapFs()
	SetFs(Fs)
	currentDir, homeDir := fileUtil.CreateMockFileSystem(t, Fs)
	configFilePath, found, err := LocateConfigFile(fileName, directory)
	assert.True(t, len(configFilePath) == 0)
	assert.Nil(t, err)
	assert.False(t, found)
	fmt.Print(currentDir, homeDir)

	// Create subdirectory in current directory, and put file there
	err = Fs.MkdirAll(filepath.Join(currentDir, directory), os.ModePerm)
	assert.Nil(t, err)
	createdPath := filepath.Join(currentDir, directory, fileName)
	_, err = Fs.Create(createdPath)
	assert.Nil(t, err)

	configFilePath, found, err = LocateConfigFile(fileName, directory)
	assert.True(t, createdPath == configFilePath)
	assert.Nil(t, err)
	assert.True(t, found)
	fmt.Print(currentDir, homeDir)
}

func TestLocatePropertiesFile_config_users_home_directory(t *testing.T) {

	fileName := "f.json"
	directory := "x"
	Fs = afero.NewMemMapFs()
	SetFs(Fs)
	currentDir, homeDir := fileUtil.CreateMockFileSystem(t, Fs)
	configFilePath, found, err := LocateConfigFile(fileName, directory)
	assert.True(t, len(configFilePath) == 0)
	assert.Nil(t, err)
	assert.False(t, found)
	fmt.Print(currentDir, homeDir)

	// Create subdirectory in users home directory that contains config file
	err = Fs.MkdirAll(filepath.Join(homeDir, directory), os.ModePerm)
	assert.Nil(t, err)
	createdPath := filepath.Join(homeDir, directory, fileName)
	_, err = Fs.Create(createdPath)
	assert.Nil(t, err)

	configFilePath, found, err = LocateConfigFile(fileName, directory)
	assert.True(t, createdPath == configFilePath)
	assert.Nil(t, err)
	assert.True(t, found)
	fmt.Print(currentDir, homeDir)
}
