module github.com/IBM/appconfiguration-go-admin-sdk

go 1.16

replace github.com/gobuffalo/packr/v2 => github.com/gobuffalo/packr/v2 v2.3.2

require (
	github.com/IBM/go-sdk-core/v5 v5.13.1
	github.com/go-openapi/strfmt v0.21.7
	github.com/go-playground/validator/v10 v10.12.0 // indirect
	github.com/leodido/go-urn v1.2.3 // indirect
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.27.1
	github.com/stretchr/testify v1.8.2
	go.mongodb.org/mongo-driver v1.11.4 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)

//Retract v1.x.x versions
retract [v1.0.0, v1.1.1]
