package go_utils

import (
	"testing"
)

// 更新到最新版本
func TestUpdateScan4allVersionToLatest(t *testing.T) {
	err := UpdateScan4allVersionToLatest(true, "hktalent", "scan4all", "./")
	if err != nil {
		t.Error("fail TestupdateNucleiVersionToLatest")
	}
}
