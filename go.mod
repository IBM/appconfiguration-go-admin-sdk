module github.com/IBM/appconfiguration-go-admin-sdk

go 1.16

require (
	github.com/IBM/go-sdk-core/v5 v5.6.4
	github.com/go-openapi/strfmt v0.20.1
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.4
	github.com/stretchr/testify v1.6.1
)

//Retract v1.x.x versions
retract [v1.0.0, v1.1.1]
