package go_openapi_spec

// https://swagger.io/specification/
type Spec struct {
	OpenApiVersion string `json:"openapi,omitempty"`
	Info *Info `json:"info,omitempty"`
	Servers Servers `json:"servers,omitempty"`
	Paths Paths `json:"paths,omitempty"`
	Components *Components `json:"components,omitempty"`
	Security []Security `json:"security"`
	Tags []Tag `json:"tags"`
	ExternalDocs ExternalDocs `json:"external_docs"`
}

type Info struct {
	Title string `json:"title"`
	Description string `json:"description"`
	TermsOfService string `json:"terms_of_service"`
	Contact struct{
		Name string `json:"name"`
		Url string `json:"url"`
		Email string `json:"email"`
	} `json:"contact"`
	License struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"license"`
	ApiVersion string `json:"api_version"`
}

type Servers []*Server
type Server struct {
	Url string `json:"url"`
	Description string `json:"description"`
	Variables map[string]struct{
		Enum []string `json:"enum"`
		Default string `json:"default"`
		Description string `json:"description"`
	} `json:"variables"`
}

type Components struct {
	Schemas map[string]*Schema `json:"schemas"`
	Responses map[string]*Response `json:"responses"`
	Parameters map[string]*Parameter `json:"parameters"`
	Examples map[string]*Example `json:"examples"`
	RequestBodies map[string]*RequestBody `json:"requestBodies"`
	Headers map[string]*Header `json:"headers"`
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes"`
	Links map[string]*Link `json:"links"`
	Callback map[string]*Callback `json:"callback"`
}

type Schema struct {

}

type Response struct {

}

const ParameterInQuery = "query"
const ParameterInHeader = "header"
const ParameterInPath = "path"
const ParameterInCookie = "cookie"

type Parameter struct {
	Name string `json:"name"`
	In string `json:"in"`
	Description string `json:"description"`
	Required bool `json:"required"`
	Deprecated bool `json:"deprecated"`
	AllowEmptyValue bool `json:"allowEmptyValue"`
	Style string `json:"style"`
	Explode bool `json:"explode"`
	AllowReserved bool `json:"allowReserved"`
	Schema interface{} `json:"schema"` // Schema or Reference
	Example interface{} `json:"example"`
	Examples map[string]interface{} `json:"examples"` // Example or Reference
}

type Example struct {

}

type RequestBody struct {

}

type Header struct {

}

type SecurityScheme struct {

}

type Link struct {

}

type Callback struct {

}

type Paths map[string]*Path
type Path struct {
	Ref string `json:"$ref"`
	Summary string `json:"summary"`
	Description string `json:"description"`
	Get *Operation `json:"get"`
	Put *Operation `json:"put"`
	Post *Operation `json:"post"`
	Delete *Operation `json:"delete"`
	Options *Operation `json:"options"`
	Head *Operation `json:"head"`
	Path *Operation `json:"path"`
	Trace *Operation `json:"trace"`
	Servers Servers `json:"servers"`
	Parameter []Parameter `json:"parameter"`
}

//https://swagger.io/specification/#operationObject
type Operation struct {
	Tags []string `json:"tags"`
	Summary string `json:"summary"`
	Description string `json:"description"`
	ExternalDocs *ExternalDocs `json:"external_docs"`
	OperationsId string `json:"operations_id"`
	Parameters []interface{} `json:"parameters"` // Parameter or Reference
	RequestBody interface{} `json:"request_body"` // RequestBody or Reference
	Response Response `json:"response"`
	Callbacks map[string]interface{} `json:"callbacks"` // Callback Object Or Reference Obj
	Deprecated bool `json:"deprecated"`
	Security []*Security `json:"security"`
	Servers Servers `json:"servers"`
}

type Security struct {

}
//https://swagger.io/specification/#referenceObject
type Reference struct {
	Ref string `json:"$ref"`
}

//https://swagger.io/specification/#tagObject
type Tag struct {

}

//https://swagger.io/specification/#externalDocumentationObject
type ExternalDocs struct {
	Description string `json:"description"`
	Url string `json:"url"`
}