// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package commands

import (
	"flag"
	"os"
	"testing"

	"github.com/mattermost/mattermost-server/api4"
	"github.com/mattermost/mattermost-server/mlog"
	"github.com/mattermost/mattermost-server/store/storetest"
	"github.com/mattermost/mattermost-server/utils"
)

func TestMain(m *testing.M) {
	// Setup a global logger to catch tests logging outside of app context
	// The global logger will be stomped by apps initalizing but that's fine for testing. Ideally this won't happen.
	mlog.InitGlobalLogger(mlog.NewLogger(&mlog.LoggerConfiguration{
		EnableConsole: true,
		ConsoleJson:   true,
		ConsoleLevel:  "error",
		EnableFile:    false,
	}))

	utils.TranslationsPreInit()

	// In the case where a dev just wants to run a single test, it's faster to just use the default
	// store.
	if filter := flag.Lookup("test.run").Value.String(); filter != "" && filter != "." {
		mlog.Info("-test.run used, not creating temporary containers")
		os.Exit(m.Run())
	}

	status := 0

	api4.UseTestStore(storetest.MySQLSettings())

	defer func() {
		api4.StopTestStore()
		os.Exit(status)
	}()

	status = m.Run()
}
