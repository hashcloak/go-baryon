package pir

// IPIRClient : Interface for a generalized PIR Client
type IPIRClient interface {
	EncodeRequest(requestedBlockIndices []uint, numberOfQueries uint) (uint, error)

	SendRequest(requestID uint) error
}

type IPIRServer interface {
	// TODO: Eventually change this signature to better handle pools of clients
	HandleRequest(clients []IPIRClient) error
}
