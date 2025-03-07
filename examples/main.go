package main

import (
	"fmt"
	"github.com/bianjieai/irita-sdk-go/types"
	opb "github.com/bianjieai/opb-sdk-go/pkg/app/sdk"
	"github.com/bianjieai/opb-sdk-go/pkg/app/sdk/model"
	"time"
)

func main() {

	// 初始化 SDK 配置
	cfg, err := types.NewClientConfig("http://localhost:26657", "ws://localhost:26657", "localhost:9090", "testing")
	if err != nil {
		panic(err)
	}

	// 初始化 OPB 网关账号
	authToken := model.NewAuthToken("TestProjectID", "TestProjectKey", "TestChainAccountAddress")

	// 创建 OPB 客户端
	client := opb.NewClient(cfg, &authToken)

	// 导入私钥
	client.Key.Recover("test_key_name", "test_password", "supreme zero ladder chaos blur lake dinner warm rely voyage scan dilemma future spin victory glance legend faculty join man mansion water mansion exotic")

	// 初始化 Tx 基础参数
	baseTx := types.BaseTx{
		From:     "test_key_name", // 对应上面导入的私钥名称
		Password: "test_password", // 对应上面导入的私钥密码
		Gas:      200000,		   // 单 Tx 消耗的 Gas 上限
		Memo:     "",			   // Tx 备注
		Mode:     types.Commit,    // Tx 广播模式
	}

	// 使用 Client 选择对应的功能模块，构造、签名并发送交易；例：发行 NFT 类别
	result, err := client.Bank.Send("iaa1lxvmp9h0v0dhzetmhstrmw3ecpplp5tljnr35f", types.NewDecCoins(types.NewDecCoin("uirita", types.NewInt(100))), baseTx)
	if err != nil {
		fmt.Errorf("转账失败: %s", err.Error())
	} else {
		fmt.Println("转账成功：", result.Hash)
	}

	// 使用 Client 选择对应的功能模块，查询链上状态；例：查询账户信息
	acc, _ := client.Bank.QueryAccount("iaa1lxvmp9h0v0dhzetmhstrmw3ecpplp5tljnr35f")
	fmt.Println("账户信息查询成功：", acc)

	// 使用 Client 订阅事件通知，例：订阅区块
	subs, err := client.SubscribeNewBlock(types.NewEventQueryBuilder(), func(block types.EventDataNewBlock) {
		fmt.Println(block)
	})

	if err != nil {
		fmt.Errorf("区块订阅失败: %s", err.Error())
	} else {
		fmt.Println("区块订阅成功：", subs.ID)
	}
	time.Sleep(time.Second * 20)
}
