// Package easyShip provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package easyShip

import (
	"time"
)

// AmazonOrderId defines model for AmazonOrderId.
type AmazonOrderId string

// Code defines model for Code.
type Code string

// List of Code
const (
	Code_InternalServerError             Code = "InternalServerError"
	Code_InvalidInput                    Code = "InvalidInput"
	Code_InvalidOrderState               Code = "InvalidOrderState"
	Code_InvalidTimeSlotId               Code = "InvalidTimeSlotId"
	Code_OrderNotEligibleForRescheduling Code = "OrderNotEligibleForRescheduling"
	Code_RegionNotSupported              Code = "RegionNotSupported"
	Code_ResourceNotFound                Code = "ResourceNotFound"
	Code_RetryableAfterGettingNewSlots   Code = "RetryableAfterGettingNewSlots"
	Code_ScheduleWindowExpired           Code = "ScheduleWindowExpired"
	Code_ScheduledPackageAlreadyExists   Code = "ScheduledPackageAlreadyExists"
	Code_TimeSlotNotAvailable            Code = "TimeSlotNotAvailable"
)

// CreateScheduledPackageRequest defines model for CreateScheduledPackageRequest.
type CreateScheduledPackageRequest struct {

	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId AmazonOrderId `json:"amazonOrderId"`

	// A string of up to 255 characters.
	MarketplaceId String `json:"marketplaceId"`

	// Package details. Includes `packageItems`, `packageTimeSlot`, and `packageIdentifier`.
	PackageDetails PackageDetails `json:"packageDetails"`
}

// CreateScheduledPackagesRequest defines model for CreateScheduledPackagesRequest.
type CreateScheduledPackagesRequest struct {

	// The file format in which the shipping label will be created.
	LabelFormat LabelFormat `json:"labelFormat"`

	// A string of up to 255 characters.
	MarketplaceId String `json:"marketplaceId"`

	// An array allowing users to specify orders to be scheduled.
	OrderScheduleDetailsList []OrderScheduleDetails `json:"orderScheduleDetailsList"`
}

// CreateScheduledPackagesResponse defines model for CreateScheduledPackagesResponse.
type CreateScheduledPackagesResponse struct {

	// A pre-signed URL for the zip document containing the shipping labels and the documents enabled for your marketplace.
	PrintableDocumentsUrl *URL `json:"printableDocumentsUrl,omitempty"`

	// A list of orders we couldn't scheduled on your behalf. Each element contains the reason and details on the error.
	RejectedOrders *[]RejectedOrder `json:"rejectedOrders,omitempty"`

	// A list of packages. Refer to the `Package` object.
	ScheduledPackages *[]Package `json:"scheduledPackages,omitempty"`
}

// DateTime defines model for DateTime.
type DateTime time.Time

// Dimension defines model for Dimension.
type Dimension float32

// Dimensions defines model for Dimensions.
type Dimensions struct {

	// The numerical value of the specified dimension.
	Height *Dimension `json:"height,omitempty"`

	// A string of up to 255 characters.
	Identifier *String `json:"identifier,omitempty"`

	// The numerical value of the specified dimension.
	Length *Dimension `json:"length,omitempty"`

	// The unit of measurement used to measure the length.
	Unit *UnitOfLength `json:"unit,omitempty"`

	// The numerical value of the specified dimension.
	Width *Dimension `json:"width,omitempty"`
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
type ErrorList struct {
	Errors []Error `json:"errors"`
}

// HandoverMethod defines model for HandoverMethod.
type HandoverMethod string

// List of HandoverMethod
const (
	HandoverMethod_Dropoff HandoverMethod = "Dropoff"
	HandoverMethod_Pickup  HandoverMethod = "Pickup"
)

// InvoiceData defines model for InvoiceData.
type InvoiceData struct {

	// A datetime value in ISO 8601 format.
	InvoiceDate *DateTime `json:"invoiceDate,omitempty"`

	// A string of up to 255 characters.
	InvoiceNumber String `json:"invoiceNumber"`
}

// Item defines model for Item.
type Item struct {

	// The Amazon-defined order item identifier.
	OrderItemId *OrderItemId `json:"orderItemId,omitempty"`

	// A list of serial numbers for the items associated with the `OrderItemId` value.
	OrderItemSerialNumbers *OrderItemSerialNumbers `json:"orderItemSerialNumbers,omitempty"`
}

// Items defines model for Items.
type Items []Item

// LabelFormat defines model for LabelFormat.
type LabelFormat string

// List of LabelFormat
const (
	LabelFormat_PDF LabelFormat = "PDF"
	LabelFormat_ZPL LabelFormat = "ZPL"
)

// ListHandoverSlotsRequest defines model for ListHandoverSlotsRequest.
type ListHandoverSlotsRequest struct {

	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId AmazonOrderId `json:"amazonOrderId"`

	// A string of up to 255 characters.
	MarketplaceId String `json:"marketplaceId"`

	// The dimensions of the scheduled package.
	PackageDimensions Dimensions `json:"packageDimensions"`

	// The weight of the scheduled package
	PackageWeight Weight `json:"packageWeight"`
}

// ListHandoverSlotsResponse defines model for ListHandoverSlotsResponse.
type ListHandoverSlotsResponse struct {

	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId AmazonOrderId `json:"amazonOrderId"`

	// A list of time slots.
	TimeSlots TimeSlots `json:"timeSlots"`
}

// OrderItemId defines model for OrderItemId.
type OrderItemId string

// OrderItemSerialNumber defines model for OrderItemSerialNumber.
type OrderItemSerialNumber string

// OrderItemSerialNumbers defines model for OrderItemSerialNumbers.
type OrderItemSerialNumbers []OrderItemSerialNumber

// OrderScheduleDetails defines model for OrderScheduleDetails.
type OrderScheduleDetails struct {

	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId AmazonOrderId `json:"amazonOrderId"`

	// Package details. Includes `packageItems`, `packageTimeSlot`, and `packageIdentifier`.
	PackageDetails *PackageDetails `json:"packageDetails,omitempty"`
}

// Package defines model for Package.
type Package struct {

	// Invoice number and date.
	Invoice *InvoiceData `json:"invoice,omitempty"`

	// The dimensions of the scheduled package.
	PackageDimensions Dimensions `json:"packageDimensions"`

	// Optional seller-created identifier that is printed on the shipping label to help the seller identify the package.
	PackageIdentifier *PackageIdentifier `json:"packageIdentifier,omitempty"`

	// A list of items contained in the package.
	PackageItems *Items `json:"packageItems,omitempty"`

	// The status of the package.
	PackageStatus *PackageStatus `json:"packageStatus,omitempty"`

	// A time window to hand over an Easy Ship package to Amazon Logistics.
	PackageTimeSlot TimeSlot `json:"packageTimeSlot"`

	// The weight of the scheduled package
	PackageWeight Weight `json:"packageWeight"`

	// Identifies the scheduled package to be updated.
	ScheduledPackageId ScheduledPackageId `json:"scheduledPackageId"`

	// Representation of tracking metadata.
	TrackingDetails *TrackingDetails `json:"trackingDetails,omitempty"`
}

// PackageDetails defines model for PackageDetails.
type PackageDetails struct {

	// Optional seller-created identifier that is printed on the shipping label to help the seller identify the package.
	PackageIdentifier *PackageIdentifier `json:"packageIdentifier,omitempty"`

	// A list of items contained in the package.
	PackageItems *Items `json:"packageItems,omitempty"`

	// A time window to hand over an Easy Ship package to Amazon Logistics.
	PackageTimeSlot TimeSlot `json:"packageTimeSlot"`
}

// PackageId defines model for PackageId.
type PackageId string

// PackageIdentifier defines model for PackageIdentifier.
type PackageIdentifier string

// PackageStatus defines model for PackageStatus.
type PackageStatus string

// List of PackageStatus
const (
	PackageStatus_AtDestinationFC  PackageStatus = "AtDestinationFC"
	PackageStatus_AtOriginFC       PackageStatus = "AtOriginFC"
	PackageStatus_DamagedInTransit PackageStatus = "DamagedInTransit"
	PackageStatus_Delivered        PackageStatus = "Delivered"
	PackageStatus_LabelCanceled    PackageStatus = "LabelCanceled"
	PackageStatus_LostInTransit    PackageStatus = "LostInTransit"
	PackageStatus_OutForDelivery   PackageStatus = "OutForDelivery"
	PackageStatus_PickedUp         PackageStatus = "PickedUp"
	PackageStatus_ReadyForPickup   PackageStatus = "ReadyForPickup"
	PackageStatus_Rejected         PackageStatus = "Rejected"
	PackageStatus_ReturnedToSeller PackageStatus = "ReturnedToSeller"
	PackageStatus_Undeliverable    PackageStatus = "Undeliverable"
)

// Packages defines model for Packages.
type Packages struct {
	Packages []Package `json:"packages"`
}

// RejectedOrder defines model for RejectedOrder.
type RejectedOrder struct {

	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId AmazonOrderId `json:"amazonOrderId"`

	// Error response returned when the request is unsuccessful.
	Error *Error `json:"error,omitempty"`
}

// ScheduledPackageId defines model for ScheduledPackageId.
type ScheduledPackageId struct {

	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId AmazonOrderId `json:"amazonOrderId"`

	// An Amazon-defined identifier for the scheduled package.
	PackageId *PackageId `json:"packageId,omitempty"`
}

// String defines model for String.
type String string

// TimeSlot defines model for TimeSlot.
type TimeSlot struct {

	// A datetime value in ISO 8601 format.
	EndTime *DateTime `json:"endTime,omitempty"`

	// Identifies the method by which a seller will hand a package over to Amazon Logistics.
	HandoverMethod *HandoverMethod `json:"handoverMethod,omitempty"`

	// A string of up to 255 characters.
	SlotId String `json:"slotId"`

	// A datetime value in ISO 8601 format.
	StartTime *DateTime `json:"startTime,omitempty"`
}

// TimeSlots defines model for TimeSlots.
type TimeSlots []TimeSlot

// TrackingDetails defines model for TrackingDetails.
type TrackingDetails struct {

	// A string of up to 255 characters.
	TrackingId *String `json:"trackingId,omitempty"`
}

// URL defines model for URL.
type URL string

// UnitOfLength defines model for UnitOfLength.
type UnitOfLength string

// List of UnitOfLength
const (
	UnitOfLength_Cm UnitOfLength = "Cm"
)

// UnitOfWeight defines model for UnitOfWeight.
type UnitOfWeight string

// List of UnitOfWeight
const (
	UnitOfWeight_G     UnitOfWeight = "G"
	UnitOfWeight_Grams UnitOfWeight = "Grams"
)

// UpdatePackageDetails defines model for UpdatePackageDetails.
type UpdatePackageDetails struct {

	// A time window to hand over an Easy Ship package to Amazon Logistics.
	PackageTimeSlot TimeSlot `json:"packageTimeSlot"`

	// Identifies the scheduled package to be updated.
	ScheduledPackageId ScheduledPackageId `json:"scheduledPackageId"`
}

// UpdatePackageDetailsList defines model for UpdatePackageDetailsList.
type UpdatePackageDetailsList []UpdatePackageDetails

// UpdateScheduledPackagesRequest defines model for UpdateScheduledPackagesRequest.
type UpdateScheduledPackagesRequest struct {

	// A string of up to 255 characters.
	MarketplaceId String `json:"marketplaceId"`

	// A list of package update details.
	UpdatePackageDetailsList UpdatePackageDetailsList `json:"updatePackageDetailsList"`
}

// Weight defines model for Weight.
type Weight struct {

	// The unit of measurement used to measure the weight.
	Unit *UnitOfWeight `json:"unit,omitempty"`

	// The weight of the package.
	Value *WeightValue `json:"value,omitempty"`
}

// WeightValue defines model for WeightValue.
type WeightValue float32

// GetScheduledPackageParams defines parameters for GetScheduledPackage.
type GetScheduledPackageParams struct {

	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId string `json:"amazonOrderId"`

	// An identifier for the marketplace in which the seller is selling.
	MarketplaceId string `json:"marketplaceId"`
}

// UpdateScheduledPackagesJSONBody defines parameters for UpdateScheduledPackages.
type UpdateScheduledPackagesJSONBody UpdateScheduledPackagesRequest

// CreateScheduledPackageJSONBody defines parameters for CreateScheduledPackage.
type CreateScheduledPackageJSONBody CreateScheduledPackageRequest

// CreateScheduledPackageBulkJSONBody defines parameters for CreateScheduledPackageBulk.
type CreateScheduledPackageBulkJSONBody CreateScheduledPackagesRequest

// ListHandoverSlotsJSONBody defines parameters for ListHandoverSlots.
type ListHandoverSlotsJSONBody ListHandoverSlotsRequest

// UpdateScheduledPackagesRequestBody defines body for UpdateScheduledPackages for application/json ContentType.
type UpdateScheduledPackagesJSONRequestBody UpdateScheduledPackagesJSONBody

// CreateScheduledPackageRequestBody defines body for CreateScheduledPackage for application/json ContentType.
type CreateScheduledPackageJSONRequestBody CreateScheduledPackageJSONBody

// CreateScheduledPackageBulkRequestBody defines body for CreateScheduledPackageBulk for application/json ContentType.
type CreateScheduledPackageBulkJSONRequestBody CreateScheduledPackageBulkJSONBody

// ListHandoverSlotsRequestBody defines body for ListHandoverSlots for application/json ContentType.
type ListHandoverSlotsJSONRequestBody ListHandoverSlotsJSONBody
