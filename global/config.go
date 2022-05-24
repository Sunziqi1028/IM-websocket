package global

import (
	"IM/logic"
)

const (
	CHATROOM = "room"
	RADIO    = "radio"
	ORIENT   = "orient"
)

var PartnerMap = make(map[uint64]struct{})

var GlobalUsers = make(map[uint64]logic.User, 1024)

var MessageChan = make(chan logic.Message, 1024)
