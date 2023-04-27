package define

type MessageStruct struct {
	Message      string `json:"message"`
	RoomIdentity string `json:"room_identity"` // 哪个房间的消息，由于用户id在头部，故不用user_identity字段
}

var MailPassword = "qtssdidxkuytbcah"
