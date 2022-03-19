package elasticsearch

// Pull ...
type Pull struct {
}

// ProductUpsert ...
func (r Pull) ProductUpsert(query UpdateDataPayload) (bool, error) {
	return publishWithJetStream(JetStreamSearchService, SubjectPullProductUpsert, toBytes(query))
}

// UserUpsert ...
func (r Pull) UserUpsert(query UpdateDataPayload) (bool, error) {
	return publishWithJetStream(JetStreamSearchService, SubjectPullUserUpsert, toBytes(query))
}

// OrderUpsert ...
func (r Pull) OrderUpsert(query UpdateDataPayload) (bool, error) {
	return publishWithJetStream(JetStreamSearchService, SubjectPullOrderUpsert, toBytes(query))
}
