package elasticsearch

const (
	SubjectRequestProductSkuUpsert      = "selly.request.product_sku.upsert"
	SubjectPullProductSkuUpsert         = "selly.pull.product_sku.upsert"
	SubjectRequestProductSkuSearch      = "selly.request.product_sku.search"
	SubjectRequestProductSkuCreateIndex = "selly.request.product_sku.create_index"

	SubjectRequestShopProductUpsert      = "selly.request.shop_product.upsert"
	SubjectPullShopProductUpsert         = "selly.pull.shop_product.upsert"
	SubjectRequestShopProductSearch      = "selly.request.shop_product.search"
	SubjectRequestShopProductCreateIndex = "selly.request.shop_product.create_index"

	SubjectRequestProductUpsert      = "selly.request.product.upsert"
	SubjectPullProductUpsert         = "selly.pull.product.upsert"
	SubjectRequestProductSearch      = "selly.request.product.search"
	SubjectRequestProductCreateIndex = "selly.request.product.create_index"

	SubjectRequestOrderUpsert      = "selly.request.order.upsert"
	SubjectPullOrderUpsert         = "selly.pull.order.upsert"
	SubjectRequestOrderSearch      = "selly.request.order.search"
	SubjectRequestOrderCreateIndex = "selly.request.order.create_index"

	SubjectRequestKeywordUpsert      = "selly.request.keyword.upsert"
	SubjectPullKeywordUpsert         = "selly.pull.keyword.upsert"
	SubjectRequestKeywordSearch      = "selly.request.keyword.search"
	SubjectRequestKeywordCreateIndex = "selly.request.keyword.create_index"

	SubjectRequestUserUpsert      = "selly.request.user.upsert"
	SubjectPullUserUpsert         = "selly.pull.user.upsert"
	SubjectRequestUserSearch      = "selly.request.user.search"
	SubjectRequestUserCreateIndex = "selly.request.user.create_index"

	SubjectRequestDeleteMultipleIndex = "selly.request.delete.multiple_index"
)

const (
	JetStreamSearchService = "Service_Search"
)
