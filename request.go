package elasticsearch

// Request ...
type Request struct {
}

// ProductSearch ...
func (r Request) ProductSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestProductSearch, toBytes(query))
}

// ProductUpsert ...
func (r Request) ProductUpsert(query UpdateDataPayload) (*Response, error) {
	return requestNats(SubjectRequestProductUpsert, toBytes(query))
}

// UserSearch ...
func (r Request) UserSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestUserSearch, toBytes(query))
}

// UserUpsert ...
func (r Request) UserUpsert(query UpdateDataPayload) (*Response, error) {
	return requestNats(SubjectRequestUserUpsert, toBytes(query))
}

// OrderSearch ...
func (r Request) OrderSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestOrderSearch, toBytes(query))
}

// OrderUpsert ...
func (r Request) OrderUpsert(query UpdateDataPayload) (*Response, error) {
	return requestNats(SubjectRequestOrderUpsert, toBytes(query))
}
