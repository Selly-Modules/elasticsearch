package elasticsearch

// Pull ...
type Pull struct {
}

// ProductUpsert ...
func (Pull) ProductUpsert(payload Payload) (bool, error) {
	return publishWithJetStream(JetStreamSearchService, SubjectPullProductUpsert, toBytes(payload))
}

// UserUpsert ...
func (Pull) UserUpsert(payload Payload) (bool, error) {
	return publishWithJetStream(JetStreamSearchService, SubjectPullUserUpsert, toBytes(payload))
}

// OrderUpsert ...
func (Pull) OrderUpsert(payload Payload) (bool, error) {
	return publishWithJetStream(JetStreamSearchService, SubjectPullOrderUpsert, toBytes(payload))
}

// KeywordUpsert ...
func (Pull) KeywordUpsert(payload Payload) (bool, error) {
	return publishWithJetStream(JetStreamSearchService, SubjectPullKeywordUpsert, toBytes(payload))
}

// ProductSkuUpsert ...
func (Pull) ProductSkuUpsert(payload Payload) (bool, error) {
	return publishWithJetStream(JetStreamSearchService, SubjectPullProductSkuUpsert, toBytes(payload))
}
