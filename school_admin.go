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

type SchoolAdmin struct {

	Credentials Credentials `json:"credentials,omitempty"`

	District string `json:"district,omitempty"`

	Email string `json:"email,omitempty"`

	Id string `json:"id,omitempty"`

	Name Name `json:"name,omitempty"`

	Schools []string `json:"schools,omitempty"`

	StaffId string `json:"staff_id,omitempty"`

	Title string `json:"title,omitempty"`
}
