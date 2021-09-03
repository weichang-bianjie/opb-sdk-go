package integration_test

import (
	"fmt"
	"github.com/bianjieai/irita-sdk-go/modules/nft"
	"github.com/stretchr/testify/require"
	"strings"
)

func (s *IntegrationTestSuite) TestQueryToken() {
	token, err := s.IRITAClient.Token.QueryToken("uirita")
	require.NoError(s.T(), err)
	fmt.Println(token)
}
func (s *IntegrationTestSuite) TestNFT() {
	denomID := "psimpc"
	nftId := strings.ToLower(s.RandStringOfLength(6))
	denomName := strings.ToLower(s.RandStringOfLength(4))
	//schema := strings.ToLower(s.RandStringOfLength(10))
	//issueReq := nft.IssueDenomRequest{
	//	ID:     denomID,
	//	Name:   denomName,
	//	Schema: schema,
	//}

	//res,err := s.IRITAClient.Status(context.Background())
	//res, err := s.NFT.IssueDenom(issueReq, s.baseTx)
	//require.NoError(s.T(), err)
	//require.NotEmpty(s.T(), res.Hash)
	//fmt.Println(res)

	res, err := s.IRITAClient.NFT.MintNFT(nft.MintNFTRequest{
		Denom: denomID,
		ID:    nftId,
		Name:  denomName,
		URI:   "http://wcchain.bianjie.ai/#/denoms",
		Data:  strings.ToLower(s.RandStringOfLength(10)),
	}, s.baseTx)
	require.NoError(s.T(), err)
	fmt.Println(res)

	es, err := s.IRITAClient.NFT.EditNFT(nft.EditNFTRequest{
		Denom: denomID,
		ID:    nftId,
		Name:  denomName,
		URI:   "http://wcchain.bianjie.ai/#/nftAsset",
		Data:  strings.ToLower(s.RandStringOfLength(10)),
	}, s.baseTx)
	require.NoError(s.T(), err)
	fmt.Println(es)

	res, er := s.IRITAClient.NFT.BurnNFT(nft.BurnNFTRequest{
		Denom: denomID,
		ID:    nftId,
	}, s.baseTx)
	require.NoError(s.T(), er)
	fmt.Println(res)
}
