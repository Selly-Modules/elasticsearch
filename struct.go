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
	Success bool     `json:"success"`
	Data    []string `json:"data,omitempty"`
	Total   int64    `json:"total,omitempty"`
	Page    int64    `json:"page,omitempty"`
	Limit   int64    `json:"limit,omitempty"`
	Message string   `json:"message"`
}

// SyncData
// Payload for sync data to service es
type SyncData struct {
	Index string
	Data  []byte
}

// UpdateDataPayload
// Payload for insert or update document
type UpdateDataPayload struct {
	Index string
	Body  []byte
}

// ESQuery
// Query support to search document
type ESQuery struct {
	Index                 string // Index
	Page                  int64
	Limit                 int64
	Keyword               string
	ProvinceCode          int
	Active                string
	Categories            []string
	SubCategories         []string
	IgnoreIDs             []string
	Suppliers             []string
	SlugCites             []string
	Type                  string
	ServiceDelivery       string
	SourceDelivery        string
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
