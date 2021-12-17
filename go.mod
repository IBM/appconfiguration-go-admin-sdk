module github.com/IBM/appconfiguration-go-admin-sdk

go 1.16

replace github.com/gobuffalo/packr/v2 => github.com/gobuffalo/packr/v2 v2.3.2

require (
	github.com/IBM/go-sdk-core/v5 v5.7.0
	github.com/go-openapi/strfmt v0.20.2
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.5
	github.com/stretchr/testify v1.7.0
)

//Retract v1.x.x versions
retract [v1.0.0, v1.1.1]
