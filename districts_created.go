/* 
 * Clever API
 *
 * The Clever API
 *
 * OpenAPI spec version: 1.2.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type DistrictsCreated struct {

	Created string `json:"created,omitempty"`

	Id string `json:"id,omitempty"`

	Type_ string `json:"type"`

	Data DistrictObject `json:"data,omitempty"`
}
