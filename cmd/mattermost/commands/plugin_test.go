package commands

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mattermost/mattermost-server/api4"
	"github.com/mattermost/mattermost-server/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlugin(t *testing.T) {
	os.MkdirAll("./test-plugins", os.ModePerm)
	os.MkdirAll("./test-client-plugins", os.ModePerm)

	th := api4.Setup(t)
	defer th.TearDown()
	th.InitBasic().InitSystemAdmin()

	path, _ := utils.FindDir("tests")

	os.Chdir(filepath.Join("..", "..", ".."))

	CheckCommand(t, "--config", th.TempConfigPath, "plugin", "add", filepath.Join(path, "testplugin.tar.gz"))

	CheckCommand(t, "--config", th.TempConfigPath, "plugin", "enable", "testplugin")
	cfg, _, _, err := utils.LoadConfig(th.TempConfigPath)
	require.Nil(t, err)
	assert.Equal(t, cfg.PluginSettings.PluginStates["testplugin"].Enable, true)

	CheckCommand(t, "--config", th.TempConfigPath, "plugin", "disable", "testplugin")
	cfg, _, _, err = utils.LoadConfig(th.TempConfigPath)
	require.Nil(t, err)
	assert.Equal(t, cfg.PluginSettings.PluginStates["testplugin"].Enable, false)

	CheckCommand(t, "--config", th.TempConfigPath, "plugin", "list")

	CheckCommand(t, "--config", th.TempConfigPath, "plugin", "delete", "testplugin")

	os.Chdir(filepath.Join("cmd", "mattermost", "commands"))
}
