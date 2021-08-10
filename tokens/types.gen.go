// Package tokens provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package tokens

// CreateRestrictedDataTokenRequest defines model for CreateRestrictedDataTokenRequest.
type CreateRestrictedDataTokenRequest struct {

	// A list of restricted resources.
	// Maximum: 50
	RestrictedResources []RestrictedResource `json:"restrictedResources"`

	// The application ID for the target application to which access is being delegated.
	TargetApplication *string `json:"targetApplication,omitempty"`
}

// CreateRestrictedDataTokenResponse defines model for CreateRestrictedDataTokenResponse.
type CreateRestrictedDataTokenResponse struct {

	// The lifetime of the Restricted Data Token, in seconds.
	ExpiresIn *int `json:"expiresIn,omitempty"`

	// A Restricted Data Token (RDT). This is a short-lived access token that authorizes you to access the restricted resources that you specified. Pass this value with the x-amz-access-token header when making subsequent calls to these restricted resources.
	RestrictedDataToken *string `json:"restrictedDataToken,omitempty"`
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
	Errors *[]Error `json:"errors,omitempty"`
}

// RestrictedResource defines model for RestrictedResource.
type RestrictedResource struct {

	// Indicates the type of Personally Identifiable Information requested. This parameter is required only when getting an RDT for use with the getOrder, getOrders, or getOrderItems operation of the Orders API. For more information, see the [Tokens API Use Case Guide](https://github.com/amzn/selling-partner-api-docs/blob/main/references/tokens-api/tokens_2021-03-01.md). Possible values include:
	// - **buyerInfo**. On the order level this includes general identifying information about the buyer and tax-related information. On the order item level this includes gift wrap information and custom order information, if available.
	// - **shippingAddress**. This includes information for fulfilling orders.
	DataElements *[]string `json:"dataElements,omitempty"`

	// The HTTP method in the restricted resource.
	Method string `json:"method"`

	// The path in the restricted resource. Here are some path examples:
	// - ```/orders/v0/orders```. For getting an RDT for the getOrders operation of the Orders API. For bulk orders.
	// - ```/orders/v0/orders/123-1234567-1234567```. For getting an RDT for the getOrder operation of the Orders API. For a specific order.
	// - ```/orders/v0/orders/123-1234567-1234567/orderItems```. For getting an RDT for the getOrderItems operation of the Orders API. For the order items in a specific order.
	// - ```/mfn/v0/shipments/FBA1234ABC5D```. For getting an RDT for the getShipment operation of the Shipping API. For a specific shipment.
	// - ```/mfn/v0/shipments/{shipmentId}```. For getting an RDT for the getShipment operation of the Shipping API. For any of a selling partner's shipments that you specify when you call the getShipment operation.
	Path string `json:"path"`
}

// CreateRestrictedDataTokenJSONBody defines parameters for CreateRestrictedDataToken.
type CreateRestrictedDataTokenJSONBody CreateRestrictedDataTokenRequest

// CreateRestrictedDataTokenRequestBody defines body for CreateRestrictedDataToken for application/json ContentType.
type CreateRestrictedDataTokenJSONRequestBody CreateRestrictedDataTokenJSONBody
