/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.9
 * Contact: nurmi@anchore.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// A subscription entry to add to the system
type SubscriptionRequest struct {

	SubscriptionKey string `json:"subscription_key,omitempty"`

	SubscriptionValue string `json:"subscription_value,omitempty"`

	SubscriptionType string `json:"subscription_type,omitempty"`
}
