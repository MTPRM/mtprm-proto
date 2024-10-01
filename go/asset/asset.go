package asset

import (
	_ "embed"
)

//go:embed index.html
var bytesIndexHtml []byte

//go:embed js/load-openapi.js
var bytesLoadOpenapiJs []byte

//go:embed openapi/merged.swagger.json
var bytesOpenapiMergedSwaggerJson []byte

var Router = map[string][]byte{
	"/":                            bytesIndexHtml,
	"/js/load-openapi.js":          bytesLoadOpenapiJs,
	"/openapi/merged.swagger.json": bytesOpenapiMergedSwaggerJson,
}
