package elasticsearch

// Request ...
type Request struct {
}

// ProductSearch ...
func (r Request) ProductSearch(query ESQuery) (*Response, error) {
	return GetClient().RequestNats(SubjectRequestProductSearch, toBytes(query))
}

// ProductUpsert ...
func (r Request) ProductUpsert(query UpdateDataPayload) (*Response, error) {
	return GetClient().RequestNats(SubjectRequestProductUpsert, toBytes(query))
}

// UserSearch ...
func (r Request) UserSearch(query ESQuery) (*Response, error) {
	return GetClient().RequestNats(SubjectRequestUserSearch, toBytes(query))
}

// UserUpsert ...
func (r Request) UserUpsert(query UpdateDataPayload) (*Response, error) {
	return GetClient().RequestNats(SubjectRequestUserUpsert, toBytes(query))
}

// OrderSearch ...
func (r Request) OrderSearch(query ESQuery) (*Response, error) {
	return GetClient().RequestNats(SubjectRequestOrderSearch, toBytes(query))
}

// OrderUpsert ...
func (r Request) OrderUpsert(query UpdateDataPayload) (*Response, error) {
	return GetClient().RequestNats(SubjectRequestOrderUpsert, toBytes(query))
}
