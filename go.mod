module github.com/tianniu-rd/storage

go 1.17

replace (
	github.com/NetEase-Object-Storage/nos-golang-sdk => github.com/karuppiah7890/nos-golang-sdk v0.0.0-20191116042345-0792ba35abcc
	go.etcd.io/etcd => github.com/eddycjy/etcd v0.5.0-alpha.5.0.20200218102753-4258cdd2efdf
	google.golang.org/grpc => google.golang.org/grpc v1.29.1
)

require (
	cloud.google.com/go/storage v1.16.1
	github.com/Azure/azure-sdk-for-go v57.1.0+incompatible
	github.com/NetEase-Object-Storage/nos-golang-sdk v0.0.0-00010101000000-000000000000
	github.com/aliyun/aliyun-oss-go-sdk v2.1.10+incompatible
	github.com/aws/aws-sdk-go v1.40.37
	github.com/baidubce/bce-sdk-go v0.9.86
	github.com/gophercloud/gophercloud v0.20.0
	github.com/oracle/oci-go-sdk v24.3.0+incompatible
	github.com/stretchr/testify v1.7.0
	github.com/tencentyun/cos-go-sdk-v5 v0.7.30
	go.etcd.io/etcd v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210903162142-ad29c8ab022f
	google.golang.org/api v0.56.0
)

require (
	github.com/Azure/go-autorest/autorest v0.11.20 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/dnaeon/go-vcr v1.2.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.16
)
