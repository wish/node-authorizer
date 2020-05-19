module k8s.io/kops/node-authorizer

go 1.14

replace k8s.io/client-go => k8s.io/client-go v0.18.1

require (
	github.com/aws/aws-sdk-go v1.31.1
	github.com/fullsailor/pkcs7 v0.0.0-20190404230743-d7302db945fa
	github.com/gorilla/mux v1.7.4
	github.com/jpillora/backoff v1.0.0
	github.com/prometheus/client_golang v1.6.0
	github.com/stretchr/testify v1.5.1
	github.com/urfave/cli v1.22.4
	go.uber.org/zap v1.15.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20200414100711-2df71ebbae66 // indirect
)
