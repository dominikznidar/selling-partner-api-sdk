// Package shipmentInvoice provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package shipmentInvoice

import (
	"time"
)

// Address defines model for Address.
type Address struct {

	// The street address.
	AddressLine1 *string `json:"AddressLine1,omitempty"`

	// Additional street address information, if required.
	AddressLine2 *string `json:"AddressLine2,omitempty"`

	// Additional street address information, if required.
	AddressLine3 *string `json:"AddressLine3,omitempty"`

	// The shipping address type.
	AddressType *AddressTypeEnum `json:"AddressType,omitempty"`

	// The city.
	City *string `json:"City,omitempty"`

	// The country code.
	CountryCode *string `json:"CountryCode,omitempty"`

	// The county.
	County *string `json:"County,omitempty"`

	// The district.
	District *string `json:"District,omitempty"`

	// The name.
	Name *string `json:"Name,omitempty"`

	// The phone number.
	Phone *string `json:"Phone,omitempty"`

	// The postal code.
	PostalCode *string `json:"PostalCode,omitempty"`

	// The state or region.
	StateOrRegion *string `json:"StateOrRegion,omitempty"`
}

// AddressTypeEnum defines model for AddressTypeEnum.
type AddressTypeEnum string

// List of AddressTypeEnum
const (
	AddressTypeEnum_Commercial  AddressTypeEnum = "Commercial"
	AddressTypeEnum_Residential AddressTypeEnum = "Residential"
)

// Blob defines model for Blob.
type Blob []byte

// BuyerTaxInfo defines model for BuyerTaxInfo.
type BuyerTaxInfo struct {

	// The legal name of the company.
	CompanyLegalName *string `json:"CompanyLegalName,omitempty"`

	// The list of tax classifications.
	TaxClassifications *TaxClassificationList `json:"TaxClassifications,omitempty"`

	// The country or region imposing the tax.
	TaxingRegion *string `json:"TaxingRegion,omitempty"`
}

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional details that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition.
	Message string `json:"message"`
}

// ErrorList defines model for ErrorList.
type ErrorList []Error

// GetInvoiceStatusResponse defines model for GetInvoiceStatusResponse.
type GetInvoiceStatusResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// The shipment invoice status response.
	Payload *ShipmentInvoiceStatusResponse `json:"payload,omitempty"`
}

// GetShipmentDetailsResponse defines model for GetShipmentDetailsResponse.
type GetShipmentDetailsResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// The information required by a selling partner to issue a shipment invoice.
	Payload *ShipmentDetail `json:"payload,omitempty"`
}

// MarketplaceTaxInfo defines model for MarketplaceTaxInfo.
type MarketplaceTaxInfo struct {

	// The legal name of the company.
	CompanyLegalName *string `json:"CompanyLegalName,omitempty"`

	// The list of tax classifications.
	TaxClassifications *TaxClassificationList `json:"TaxClassifications,omitempty"`

	// The country or region imposing the tax.
	TaxingRegion *string `json:"TaxingRegion,omitempty"`
}

// Money defines model for Money.
type Money struct {

	// The currency amount.
	Amount *string `json:"Amount,omitempty"`

	// Three-digit currency code in ISO 4217 format.
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
}

// PaymentMethodDetailItemList defines model for PaymentMethodDetailItemList.
type PaymentMethodDetailItemList []string

// SerialNumbersList defines model for SerialNumbersList.
type SerialNumbersList []string

// ShipmentDetail defines model for ShipmentDetail.
type ShipmentDetail struct {

	// The Amazon-defined identifier for the order.
	AmazonOrderId *string `json:"AmazonOrderId,omitempty"`

	// The Amazon-defined identifier for the shipment.
	AmazonShipmentId *string `json:"AmazonShipmentId,omitempty"`

	// The county of the buyer.
	BuyerCounty *string `json:"BuyerCounty,omitempty"`

	// The name of the buyer.
	BuyerName *string `json:"BuyerName,omitempty"`

	// Tax information about the buyer.
	BuyerTaxInfo *BuyerTaxInfo `json:"BuyerTaxInfo,omitempty"`

	// The identifier for the marketplace where the order was placed.
	MarketplaceId *string `json:"MarketplaceId,omitempty"`

	// Tax information about the marketplace.
	MarketplaceTaxInfo *MarketplaceTaxInfo `json:"MarketplaceTaxInfo,omitempty"`

	// The list of payment method details.
	PaymentMethodDetails *PaymentMethodDetailItemList `json:"PaymentMethodDetails,omitempty"`

	// The date and time when the order was created.
	PurchaseDate *time.Time `json:"PurchaseDate,omitempty"`

	// The seller’s friendly name registered in the marketplace.
	SellerDisplayName *string `json:"SellerDisplayName,omitempty"`

	// The seller identifier.
	SellerId *string `json:"SellerId,omitempty"`

	// A list of shipment items.
	ShipmentItems *ShipmentItems `json:"ShipmentItems,omitempty"`

	// The shipping address details of the shipment.
	ShippingAddress *Address `json:"ShippingAddress,omitempty"`

	// The Amazon-defined identifier for the warehouse.
	WarehouseId *string `json:"WarehouseId,omitempty"`
}

// ShipmentInvoiceStatus defines model for ShipmentInvoiceStatus.
type ShipmentInvoiceStatus string

// List of ShipmentInvoiceStatus
const (
	ShipmentInvoiceStatus_Accepted   ShipmentInvoiceStatus = "Accepted"
	ShipmentInvoiceStatus_Errored    ShipmentInvoiceStatus = "Errored"
	ShipmentInvoiceStatus_NotFound   ShipmentInvoiceStatus = "NotFound"
	ShipmentInvoiceStatus_Processing ShipmentInvoiceStatus = "Processing"
)

// ShipmentInvoiceStatusInfo defines model for ShipmentInvoiceStatusInfo.
type ShipmentInvoiceStatusInfo struct {

	// The Amazon-defined shipment identifier.
	AmazonShipmentId *string `json:"AmazonShipmentId,omitempty"`

	// The shipment invoice status.
	InvoiceStatus *ShipmentInvoiceStatus `json:"InvoiceStatus,omitempty"`
}

// ShipmentInvoiceStatusResponse defines model for ShipmentInvoiceStatusResponse.
type ShipmentInvoiceStatusResponse struct {

	// The shipment invoice status information.
	Shipments *ShipmentInvoiceStatusInfo `json:"Shipments,omitempty"`
}

// ShipmentItem defines model for ShipmentItem.
type ShipmentItem struct {

	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN *string `json:"ASIN,omitempty"`

	// The currency type and amount.
	GiftWrapPrice *Money `json:"GiftWrapPrice,omitempty"`

	// The currency type and amount.
	ItemPrice *Money `json:"ItemPrice,omitempty"`

	// The Amazon-defined identifier for the order item.
	OrderItemId *string `json:"OrderItemId,omitempty"`

	// The currency type and amount.
	PromotionDiscount *Money `json:"PromotionDiscount,omitempty"`

	// The number of items ordered.
	QuantityOrdered *float32 `json:"QuantityOrdered,omitempty"`

	// The seller SKU of the item.
	SellerSKU *string `json:"SellerSKU,omitempty"`

	// The list of serial numbers.
	SerialNumbers *SerialNumbersList `json:"SerialNumbers,omitempty"`

	// The currency type and amount.
	ShippingDiscount *Money `json:"ShippingDiscount,omitempty"`

	// The currency type and amount.
	ShippingPrice *Money `json:"ShippingPrice,omitempty"`

	// The name of the item.
	Title *string `json:"Title,omitempty"`
}

// ShipmentItems defines model for ShipmentItems.
type ShipmentItems []ShipmentItem

// SubmitInvoiceRequest defines model for SubmitInvoiceRequest.
type SubmitInvoiceRequest struct {

	// MD5 sum for validating the invoice data. For more information about calculating this value, see [Working with Content-MD5 Checksums](https://docs.developer.amazonservices.com/en_US/dev_guide/DG_MD5.html).
	ContentMD5Value string `json:"ContentMD5Value"`

	// Shipment invoice document content.
	InvoiceContent Blob `json:"InvoiceContent"`

	// An Amazon marketplace identifier.
	MarketplaceId *string `json:"MarketplaceId,omitempty"`
}

// SubmitInvoiceResponse defines model for SubmitInvoiceResponse.
type SubmitInvoiceResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`
}

// TaxClassification defines model for TaxClassification.
type TaxClassification struct {

	// The type of tax.
	Name *string `json:"Name,omitempty"`

	// The entity's tax identifier.
	Value *string `json:"Value,omitempty"`
}

// TaxClassificationList defines model for TaxClassificationList.
type TaxClassificationList []TaxClassification

// SubmitInvoiceJSONBody defines parameters for SubmitInvoice.
type SubmitInvoiceJSONBody SubmitInvoiceRequest

// SubmitInvoiceRequestBody defines body for SubmitInvoice for application/json ContentType.
type SubmitInvoiceJSONRequestBody SubmitInvoiceJSONBody
