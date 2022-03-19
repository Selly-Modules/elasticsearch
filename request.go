package elasticsearch

// Request ...
type Request struct {
}

// ProductSearch ...
func (Request) ProductSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestProductSearch, toBytes(query))
}

// ProductUpsert ...
func (Request) ProductUpsert(query UpdateDataPayload) (*Response, error) {
	return requestNats(SubjectRequestProductUpsert, toBytes(query))
}

// UserSearch ...
func (Request) UserSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestUserSearch, toBytes(query))
}

// UserUpsert ...
func (Request) UserUpsert(query UpdateDataPayload) (*Response, error) {
	return requestNats(SubjectRequestUserUpsert, toBytes(query))
}

// OrderSearch ...
func (Request) OrderSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestOrderSearch, toBytes(query))
}

// OrderUpsert ...
func (Request) OrderUpsert(query UpdateDataPayload) (*Response, error) {
	return requestNats(SubjectRequestOrderUpsert, toBytes(query))
}
