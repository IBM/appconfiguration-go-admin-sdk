module github.com/IBM/appconfiguration-go-admin-sdk

go 1.16

replace github.com/gobuffalo/packr/v2 => github.com/gobuffalo/packr/v2 v2.3.2

require (
	github.com/IBM/go-sdk-core/v5 v5.15.0
	github.com/go-openapi/strfmt v0.21.7
	github.com/kr/pretty v0.3.0 // indirect
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.27.6
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/stretchr/testify v1.8.2
	go.mongodb.org/mongo-driver v1.11.4 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

//Retract v1.x.x versions
retract [v1.0.0, v1.1.1]
