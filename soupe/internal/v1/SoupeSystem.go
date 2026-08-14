package soupev1

import (
	"fmt"
	"time"

	limbov1 "github.com/limboware/limbo"
	errnov1 "github.com/rejchev/errno"
	"github.com/untrustedmodders/s2sdkv213/s2sdk"
)

type SoupeSystem struct{}

func (x *SoupeSystem) Init() errnov1.Code {
	limbov1.Events().Subscribe("world.loaded", x.onAllLoaded)
	limbov1.Events().Subscribe("s2sdk.cmd", x.onCmdUse)
	return errnov1.OK
}

func (x *SoupeSystem) onCmdUse(_ string, data any) {
	if cmd, ok := data.(*UseCommand); ok && s2sdk.IsClientAuthorized(cmd.Caller) {
		steamId := s2sdk.GetClientSteamID64(cmd.Caller)

		ticket := Ticket{}
		if !Tickets().Iterator().First(&ticket, func(t Ticket) bool {
			return Tickets().Requester(t) == steamId
		}) {
			if !Tickets().Create(steamId, 3600, &ticket) {
				return
			}
		}

		exp := Tickets().CreatedAt(ticket) + Tickets().Duration(ticket) - time.Now().Unix()

		s2sdk.PrintToChat(cmd.Caller, fmt.Sprintf("Your ticket %s expired after %.2f", string(ticket[:]), time.Duration(exp).Minutes()))
	}
}

func (x *SoupeSystem) onAllLoaded(string, any) {

}
