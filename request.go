package elasticsearch

// Request ...
type Request struct {
}

// ProductSearch ...
func (Request) ProductSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestProductSearch, toBytes(query))
}

// ProductUpsert ...
func (Request) ProductUpsert(payload Payload) (*Response, error) {
	return requestNats(SubjectRequestProductUpsert, toBytes(payload))
}

// UserSearch ...
func (Request) UserSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestUserSearch, toBytes(query))
}

// UserUpsert ...
func (Request) UserUpsert(payload Payload) (*Response, error) {
	return requestNats(SubjectRequestUserUpsert, toBytes(payload))
}

// OrderSearch ...
func (Request) OrderSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestOrderSearch, toBytes(query))
}

// OrderUpsert ...
func (Request) OrderUpsert(payload Payload) (*Response, error) {
	return requestNats(SubjectRequestOrderUpsert, toBytes(payload))
}

// KeywordSearch ...
func (Request) KeywordSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestKeywordSearch, toBytes(query))
}

// KeywordUpsert ...
func (Request) KeywordUpsert(payload Payload) (*Response, error) {
	return requestNats(SubjectRequestKeywordUpsert, toBytes(payload))
}

// CreateIndex ...
func (Request) CreateIndex(payload Payload) (*Response, error) {
	return requestNats(SubjectRequestCreateIndex, toBytes(payload))
}
