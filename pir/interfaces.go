package pir

type IPIRClient interface {
	SetParams()

	EncodeRequest()

	GetResult()

	SendRequest()

	ReceiveReplies()

	ProcessReplies()
}

type IPIRServer interface {
	SetParams()

	HandleRequest()
}
