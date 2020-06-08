package superlink

const (
	TypeAsset     = "Asset"
	TypeDigital   = "Digital"
	TypeService   = "Service"
	TypePayment   = "Payment"
	TypeRecurrent = "Recurrent"
)

type Link struct {
	Name                    string `json:"name,omitempty"`
	Description             string `json:"description,omitempty"`
	ShowDescription         bool   `json:"showDescription,omitempty"`
	Price                   int    `json:"price,omitempty"`
	ExpirationDate          string `json:"expirationDate,omitempty"`
	Weight                  int    `json:"weight,omitempty"`
	SoftDescriptor          string `json:"softDescriptor,omitempty"`
	MaxNumberOfInstallments string `json:"maxNumberOfInstallments,omitempty"`
	Type                    string `json:"type,omitempty"`

	Shipping   Shipping   `json:"shipping,omitempty"`
	Recurrency Recurrency `json:"recurrent,omitempty"`

	// Non documented fields
	Quantity int    `json:"quantity,omitempty"`
	Sku      string `json:"sku,omitempty"`
}

const (
	ShippingTypeCorreios              = "Correios"
	ShippingFixedPrice                = "FixedAmount"
	ShippingFree                      = "Free"
	ShippingTypeWithoutShippingPickUp = "WithoutShippingPickUp"
	ShippingTypeWithoutShipping       = "WithoutShipping"
)

type Shipping struct {
	Name          string `json:"name,omitempty"`
	Price         string `json:"price,omitempty"`
	OriginZipCode string `json:"originZipCode,omitempty"`
	Type          string `json:"type,omitempty"`
}

const (
	IntervalMonthly    = "Monthly"
	IntervalBimonthly  = "Bimonthly"
	IntervalQuarterly  = "Quarterly"
	IntervalSemiAnnual = "SemiAnnual"
	IntervalAnnual     = "Annual"
)

type Recurrency struct {
	Interval string `json:"interval,omitempty"`
	EndDate  string `json:"endDate,omitempty"`
}

type LinkCreated struct {
	Link
	ID       string `json:"id,omitempty"`
	ShortURL string `json:"shortUrl,omitempty"`
	Links    []struct {
		Method string `json:"method,omitempty"`
		Rel    string `json:"rel,omitempty"`
		HRef   string `json:"rel,omitempty"`
	} `json:"links,omitempty"`
}
