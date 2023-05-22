package go_mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

var url string = "mongodb://root:L9S2J2BctwQor1b5iMkgomCm@voicemaker-test-sharding-01-01.mongodb.aliyun.voicemaker.media:3717,voicemaker-test-sharding-01-02.mongodb.aliyun.voicemaker.media:3717/admin"

func Test0201(t *testing.T) {
	iResult, err := client.Database("audit").Collection("user_audit_record").CountDocuments(context.Background(), bson.M{"uid": 900099, "audit_status": 2})
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("", iResult)
}
