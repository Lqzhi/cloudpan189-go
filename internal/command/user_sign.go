package command

import (
	"fmt"
	"github.com/tickstep/cloudpan189-go/cloudpan"
)

func RunUserSign() {
	activeUser := GetActiveUser()
	result, err := activeUser.PanClient().AppUserSign(&activeUser.AppToken)
	if err != nil {
		fmt.Printf("签到失败: %s\n", err)
		return
	}
	if result.Status == cloudpan.AppUserSignStatusSuccess {
		fmt.Printf("签到成功，%s", result.Tip)
	} else if result.Status == cloudpan.AppUserSignStatusHasSign {
		fmt.Printf("今日已签到，%s", result.Tip)
	} else {
		fmt.Printf("签到失败，%s", result.Tip)
	}
}
