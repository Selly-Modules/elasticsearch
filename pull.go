package elasticsearch

// Pull ...
type Pull struct {
}

// ProductUpsert ...
func (Pull) ProductUpsert(query UpdateDataPayload) (bool, error) {
	return publishWithJetStream(JetStreamSearchService, SubjectPullProductUpsert, toBytes(query))
}

// UserUpsert ...
func (Pull) UserUpsert(query UpdateDataPayload) (bool, error) {
	return publishWithJetStream(JetStreamSearchService, SubjectPullUserUpsert, toBytes(query))
}

// OrderUpsert ...
func (Pull) OrderUpsert(query UpdateDataPayload) (bool, error) {
	return publishWithJetStream(JetStreamSearchService, SubjectPullOrderUpsert, toBytes(query))
}
