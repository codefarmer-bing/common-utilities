package wxpay

import (
	"gopkg.in/chanxuehong/wechat.v2/mch/core"
)

// NewWechatPaySSLClient
func NewWechatPaySSLClient(appId,mchId,mchKey,sslCertPath,sslKeyPath string) (*core.Client,error){
	sslHttpClient,err := core.NewTLSHttpClient(sslCertPath,sslKeyPath)
	if err != nil {
		return nil, err
	}
	client := core.NewClient(appId,mchId,mchKey,sslHttpClient)
	return client, nil
}
