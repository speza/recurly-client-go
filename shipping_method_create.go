package recurly

import ()

type ShippingMethodCreate struct {
	Params `json:"-"`

	// The internal name used identify the shipping method.
	Code *string `json:"code,omitempty"`

	// The name of the shipping method displayed to customers.
	Name *string `json:"name,omitempty"`

	// Accounting code for shipping method.
	AccountingCode *string `json:"accounting_code,omitempty"`

	// Used by Avalara, Vertex, and Recurly’s built-in tax feature. The tax
	// code values are specific to each tax system. If you are using Recurly’s
	// built-in taxes the values are:
	// - `FR` – Common Carrier FOB Destination
	// - `FR022000` – Common Carrier FOB Origin
	// - `FR020400` – Non Common Carrier FOB Destination
	// - `FR020500` – Non Common Carrier FOB Origin
	// - `FR010100` – Delivery by Company Vehicle Before Passage of Title
	// - `FR010200` – Delivery by Company Vehicle After Passage of Title
	// - `NT` – Non-Taxable
	TaxCode *string `json:"tax_code,omitempty"`
}

func (attr *ShippingMethodCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
