package recurly

import ()

type AccountCreate struct {
	Params `json:"-"`

	// The unique identifier of the account. This cannot be changed once the account is created.
	Code *string `json:"code,omitempty"`

	Acquisition *AccountAcquisitionUpdatable `json:"acquisition,omitempty"`

	ShippingAddresses []ShippingAddressCreate `json:"shipping_addresses,omitempty"`

	// A secondary value for the account.
	Username *string `json:"username,omitempty"`

	// The email address used for communicating with this customer. The customer will also use this email address to log into your hosted account management pages. This value does not need to be unique.
	Email *string `json:"email,omitempty"`

	// Used to determine the language and locale of emails sent on behalf of the merchant to the customer. The list of locales is restricted to those the merchant has enabled on the site.
	PreferredLocale *string `json:"preferred_locale,omitempty"`

	// Additional email address that should receive account correspondence. These should be separated only by commas. These CC emails will receive all emails that the `email` field also receives.
	CcEmails *string `json:"cc_emails,omitempty"`

	FirstName *string `json:"first_name,omitempty"`

	LastName *string `json:"last_name,omitempty"`

	Company *string `json:"company,omitempty"`

	// The VAT number of the account (to avoid having the VAT applied). This is only used for manually collected invoices.
	VatNumber *string `json:"vat_number,omitempty"`

	// The tax status of the account. `true` exempts tax on the account, `false` applies tax on the account.
	TaxExempt *bool `json:"tax_exempt,omitempty"`

	// The tax exemption certificate number for the account. If the merchant has an integration for the Vertex tax provider, this optional value will be sent in any tax calculation requests for the account.
	ExemptionCertificate *string `json:"exemption_certificate,omitempty"`

	// The account code of the parent account to be associated with this account. Passing an empty value removes any existing parent association from this account. If both `parent_account_code` and `parent_account_id` are passed, the non-blank value in `parent_account_id` will be used. Only one level of parent child relationship is allowed. You cannot assign a parent account that itself has a parent account.
	ParentAccountCode *string `json:"parent_account_code,omitempty"`

	// The UUID of the parent account to be associated with this account. Passing an empty value removes any existing parent association from this account. If both `parent_account_code` and `parent_account_id` are passed, the non-blank value in `parent_account_id` will be used. Only one level of parent child relationship is allowed. You cannot assign a parent account that itself has a parent account.
	ParentAccountId *string `json:"parent_account_id,omitempty"`

	// An enumerable describing the billing behavior of the account, specifically whether the account is self-paying or will rely on the parent account to pay.
	BillTo *string `json:"bill_to,omitempty"`

	// An optional type designation for the payment gateway transaction created by this request. Supports 'moto' value, which is the acronym for mail order and telephone transactions.
	TransactionType *string `json:"transaction_type,omitempty"`

	Address *AddressCreate `json:"address,omitempty"`

	BillingInfo *BillingInfoCreate `json:"billing_info,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomFieldCreate `json:"custom_fields,omitempty"`
}

func (attr *AccountCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
