package soupev1

import (
	limbov1 "github.com/limboware/limbo"
	pbsoupev1 "github.com/limboware/pkg/proto/soupe/v1"
	"google.golang.org/protobuf/proto"
)

func SendNetMessage(typ limbov1.NetMessageType, payload proto.Message, data []byte, recipe limbov1.Entity) bool {
	if netsys := limbov1.SystemPtr[NetSystem](); netsys != nil {
		return netsys.SendMessage(typ, payload, data, recipe)
	}

	return false
}

func SendNetMessageB(typ limbov1.NetMessageType, data []byte, recipe limbov1.Entity) bool {
	return SendNetMessage(typ, nil, data, recipe)
}

func SendNetMessageC(typ limbov1.NetMessageType, recipe limbov1.Entity) bool {
	return SendNetMessageB(typ, nil, recipe)
}

func SendNetMessageD(typ limbov1.NetMessageType, payload proto.Message, recipe limbov1.Entity) bool {
	return SendNetMessage(typ, payload, nil, recipe)
}

func SendPing(e limbov1.Entity) bool {
	return SendNetMessageC(limbov1.NetMessageType(pbsoupev1.MsgType_Ping), e)
}

func SendNetMessageListed(typ limbov1.NetMessageType, payload proto.Message, data []byte, recipe ...limbov1.Entity) bool {
	if netsys := limbov1.SystemPtr[NetSystem](); netsys != nil {
		return netsys.SendMessageListed(typ, payload, data, recipe)
	}

	return false
}

func NetDisconnect(e limbov1.Entity, reason string) bool {
	dptr := limbov1.ComponentPtr[Destroy](e, DestroyType)

	if dptr == nil {
		if dptr = limbov1.NewComponentB[Destroy](e, DestroyType); dptr == nil {
			return false
		}
	}

	if reason != "" {
		dptr.Reason = reason
	}

	return true
}
