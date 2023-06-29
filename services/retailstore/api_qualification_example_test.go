// Copyright 2021 Tencent Inc. All rights reserved.
//
// 营销加价购对外API
//
// 指定服务商可通过该接口报名加价购活动、查询某个区域内的加价购活动列表、锁定加价活动购资格以及解锁加价购活动资格。
//
// API version: 1.3.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package retailstore_test

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/retailstore"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func ExampleQualificationApiService_LockQualification() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := retailstore.QualificationApiService{Client: client}
	resp, result, err := svc.LockQualification(ctx,
		retailstore.LockQualificationRequest{
			OrderInformation: &retailstore.OrderInfo{
				PayerOpenid:     core.String("oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"),
				OutTradeNo:      core.String("1217752501201407033233368018"),
				TotalFee:        core.Int64(100),
				StoreId:         core.String("123"),
				StoreMerchantId: core.String("1230000109"),
			},
			QualificationIds: []string{"8495134018"},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call LockQualification err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleQualificationApiService_UnlockQualification() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := retailstore.QualificationApiService{Client: client}
	resp, result, err := svc.UnlockQualification(ctx,
		retailstore.UnlockQualificationRequest{
			OrderInformation: &retailstore.OrderInfo{
				PayerOpenid:     core.String("oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"),
				OutTradeNo:      core.String("1217752501201407033233368018"),
				TotalFee:        core.Int64(100),
				StoreId:         core.String("123"),
				StoreMerchantId: core.String("1230000109"),
			},
			QualificationIds: []string{"8495134018"},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call UnlockQualification err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
