package elasticsearch

import "time"

// RequestBody ...
type RequestBody struct {
	ApiKey string `json:"apiKey"`
	Body   []byte `json:"body"`
}

// Response
// response to service es
type Response struct {
	Success bool   `json:"success"`
	Data    []byte `json:"data,omitempty"`
	Total   int64  `json:"total,omitempty"`
	Page    int64  `json:"page,omitempty"`
	Limit   int64  `json:"limit,omitempty"`
	Message string `json:"message"`
}

// Payload ...
// payload for sync data to service es
type Payload struct {
	Index string
	Data  []byte
}

// DeleteDataPayload
// Payload for delete document
type DeleteDataPayload struct {
	Index string
	ID    string
}

// ESQuery
// Query support to search document
type ESQuery struct {
	IsMatch               bool   // Search with match or prefix
	Index                 string // Index
	Page                  int64
	Limit                 int64
	Keyword               string
	ProvinceCode          int
	Active                string
	IsOutOfStock          string
	CanIssueInvoice       string
	PendingInactive       string
	Categories            []string
	SubCategories         []string
	IgnoreIDs             []string
	Suppliers             []string
	SlugCites             []string
	Type                  string
	ServiceDelivery       string
	SourceDelivery        string
	Brands                []string
	NoBrand               string
	Banned                string
	ListUser              []string
	ListNotUser           []string
	PaymentMethod         string
	Source                string
	FromNewActiveSeller   string
	FromNewActiveBuyer    string
	EmailStatus           string
	MerchantStatus        string
	IsCalled              string
	IsAutoApproved        string
	ProcessStatus         string
	OutboundRequestStatus string
	IsWholesaleBonus      string
	IsPreorder            string
	IsDeleted             string
	Tags                  []string
	Sorts                 []ESSort
	ListStatus            []string
	ListDeliveryStatus    []string
	FromAt                time.Time
	ToAt                  time.Time
	ApprovedFrom          time.Time
	ApprovedTo            time.Time
	DeliveredFrom         time.Time
	DeliveredTo           time.Time
	CashbackFrom          time.Time
	CashbackTo            time.Time
	FromPrice             float64
	ToPrice               float64
	Inventories           []string
	NotInventories        []string
	ReferralCode          string
	MembershipLevel       int
	Invitee               string
	Segments              []string
}

// ESSort
// ES sort with field
// ... filed is sort
// ... ascending [true is asc] [false is desc]
type ESSort struct {
	Field     string // Filed sort
	Ascending bool
}
