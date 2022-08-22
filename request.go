package elasticsearch

// Request ...
type Request struct {
}

// ProductShopSearch ...
func (Request) ProductShopSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestShopProductSearch, toBytes(query))
}

// ProductShopUpsert ...
func (Request) ProductShopUpsert(payload Payload) (*Response, error) {
	return requestNats(SubjectRequestShopProductUpsert, toBytes(payload))
}

// ProductShopCreateIndex ...
func (Request) ProductShopCreateIndex() (*Response, error) {
	return requestNats(SubjectRequestShopProductCreateIndex, toBytes(Payload{}))
}

// ProductSkuSearch ...
func (Request) ProductSkuSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestProductSkuSearch, toBytes(query))
}

// ProductSkuUpsert ...
func (Request) ProductSkuUpsert(payload Payload) (*Response, error) {
	return requestNats(SubjectRequestProductSkuUpsert, toBytes(payload))
}

// ProductSkuCreateIndex ...
func (Request) ProductSkuCreateIndex() (*Response, error) {
	return requestNats(SubjectRequestProductSkuCreateIndex, toBytes(Payload{}))
}

// ProductSearch ...
func (Request) ProductSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestProductSearch, toBytes(query))
}

// ProductUpsert ...
func (Request) ProductUpsert(payload Payload) (*Response, error) {
	return requestNats(SubjectRequestProductUpsert, toBytes(payload))
}

// ProductCreateIndex ...
func (Request) ProductCreateIndex() (*Response, error) {
	return requestNats(SubjectRequestProductCreateIndex, toBytes(Payload{}))
}

// UserSearch ...
func (Request) UserSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestUserSearch, toBytes(query))
}

// UserUpsert ...
func (Request) UserUpsert(payload Payload) (*Response, error) {
	return requestNats(SubjectRequestUserUpsert, toBytes(payload))
}

// UserCreateIndex ...
func (Request) UserCreateIndex() (*Response, error) {
	return requestNats(SubjectRequestUserCreateIndex, toBytes(Payload{}))
}

// OrderSearch ...
func (Request) OrderSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestOrderSearch, toBytes(query))
}

// OrderUpsert ...
func (Request) OrderUpsert(payload Payload) (*Response, error) {
	return requestNats(SubjectRequestOrderUpsert, toBytes(payload))
}

// OrderCreateIndex ...
func (Request) OrderCreateIndex() (*Response, error) {
	return requestNats(SubjectRequestOrderCreateIndex, toBytes(Payload{}))
}

// KeywordSearch ...
func (Request) KeywordSearch(query ESQuery) (*Response, error) {
	return requestNats(SubjectRequestKeywordSearch, toBytes(query))
}

// KeywordUpsert ...
func (Request) KeywordUpsert(payload Payload) (*Response, error) {
	return requestNats(SubjectRequestKeywordUpsert, toBytes(payload))
}

// KeywordCreateIndex ...
func (Request) KeywordCreateIndex() (*Response, error) {
	return requestNats(SubjectRequestKeywordCreateIndex, toBytes(Payload{}))
}

// DeleteMultipleIndex ...
func (Request) DeleteMultipleIndex(indexes []string) (*Response, error) {
	return requestNats(SubjectRequestDeleteMultipleIndex, toBytes(indexes))
}
