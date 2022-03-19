package elasticsearch

// Queue ...
type Queue struct {
}

// ProductUpsert ...
func (r Queue) ProductUpsert(query UpdateDataPayload) (bool, error) {
	return GetClient().PublishWithJetStream(JetStreamSearchService, SubjectQueueProductUpsert, toBytes(query))
}

// UserUpsert ...
func (r Queue) UserUpsert(query UpdateDataPayload) (bool, error) {
	return GetClient().PublishWithJetStream(JetStreamSearchService, SubjectQueueUserUpsert, toBytes(query))
}

// OrderUpsert ...
func (r Queue) OrderUpsert(query UpdateDataPayload) (bool, error) {
	return GetClient().PublishWithJetStream(JetStreamSearchService, SubjectQueueOrderUpsert, toBytes(query))
}
