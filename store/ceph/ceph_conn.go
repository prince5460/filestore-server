package ceph

import (
	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
)

var cephConn *s3.S3

//GetCephConnection:获取ceph连接
func GetCephConnection() *s3.S3 {
	if cephConn != nil {
		return cephConn
	}
	//1.初始化ceph的一些信息
	auth := aws.Auth{
		AccessKey: "Y7E7AFFR4MU2SJD1YK4T",
		SecretKey: "BtkIV4EZNv7zvIC7JjO9vhmuq2ieEGm5k2xmQbqo",
	}

	curRegion := aws.Region{
		Name:                 "default",
		EC2Endpoint:          "http://127.0.0.1:9080",
		S3Endpoint:           "http://127.0.0.1:9080",
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
		Sign:                 aws.SignV2,
	}

	//2.创建s3类型的连接
	return s3.New(auth, curRegion)
}

//GetCephBucket:获取指定bucket对象
func GetCephBucket(bucket string) *s3.Bucket {
	conn := GetCephConnection()
	return conn.Bucket(bucket)
}
