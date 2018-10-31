// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api4

import (
	"testing"
)

func TestGetSamlMetadata(t *testing.T) {
	th := Setup(t)
	defer th.TearDown()
	th.InitBasic().InitSystemAdmin()

	Client := th.Client

	_, resp := Client.GetSamlMetadata()
	CheckNotImplementedStatus(t, resp)

	// Rest is tested by enterprise tests
}
