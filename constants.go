package elasticsearch

const (
	SubjectRequestProductUpsert = "elasticsearch/selly.request.product.upsert"
	SubjectQueueProductUpsert   = "elasticsearch/selly.queue.product.upsert"
	SubjectPullProductUpsert    = "elasticsearch/selly.pull.product.upsert"
	SubjectRequestProductSearch = "elasticsearch/selly.request.product.search"

	SubjectRequestOrderUpsert = "elasticsearch/selly.request.order.upsert"
	SubjectQueueOrderUpsert   = "elasticsearch/selly.queue.order.upsert"
	SubjectPullOrderUpsert    = "elasticsearch/selly.pull.order.upsert"
	SubjectRequestOrderSearch = "elasticsearch/selly.request.order.search"

	SubjectRequestUserUpsert = "elasticsearch/selly.request.order.upsert"
	SubjectQueueUserUpsert   = "elasticsearch/selly.queue.order.upsert"
	SubjectPushUserUpsert    = "elasticsearch/selly.push.order.upsert"
	SubjectPullUserUpsert    = "elasticsearch/selly.pull.order.upsert"
	SubjectRequestUserSearch = "elasticsearch/selly.request.user.search"
)

const (
	JetStreamSearchService = "Service_Search"
)
