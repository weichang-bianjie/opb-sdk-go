package integration_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
)

func (s *IntegrationTestSuite) TestQueryToken() {
	token, err := s.IRITAClient.Token.QueryToken("uirita")
	require.NoError(s.T(), err)
	fmt.Println(token)
	fmt.Println(strings.ToLower(s.RandStringOfLength(10)))
}
func (s *IntegrationTestSuite) TestNFT() {
	//denomID := "psimpc"
	//denomName := "文昌链数字艺术品"
	//nftId := strings.ToLower(s.RandStringOfLength(10))
	//nftName := "《富春山居图剩山图卷》"

	//denomID := strings.ToLower(s.RandStringOfLength(10))
	//schema := "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"$id\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"Digital Art  Denom schema meta-schema\",\"type\":\"object\",\"properties\":{\"digital_art\":{\"type\":\"object\",\"properties\":{\"holder\":{\"type\":\"string\",\"description\":\"holder of digital art\"},\"img\":{\"type\":\"string\",\"description\":\"img of digital art\"},\"icon\":{\"type\":\"string\",\"description\":\"icon of digital art\"},\"status\":{\"type\":\"string\",\"enum\":[\"active\",\"sold\"]},\"description\":{\"type\":\"string\"},\"issuer\":{\"type\":\"array\",\"items\":{\"type\":\"string\"},\"description\":\"issuer of digital art\"},\"escrower\":{\"type\":\"array\",\"items\":{\"type\":\"string\"},\"description\":\"escrower of digital art\"},\"union_issue\":{\"type\":\"array\",\"items\":{\"type\":\"string\"},\"description\":\"union issuer of digital art\"}},\"required\":[\"holder\",\"img\",\"icon\",\"status\",\"description\",\"issuer\",\"escrower\",\"union_issue\"]}},\"definitions\":{}}"
	//issueReq := nft.IssueDenomRequest{
	//	ID:     denomID,
	//	Name:   denomName,
	//	Schema: schema,
	//}
	//
	//res, err := s.NFT.IssueDenom(issueReq, s.baseTx)
	//require.NoError(s.T(), err)
	//require.NotEmpty(s.T(), res.Hash)
	//fmt.Println(res)

	//mintNft := nft.MintNFTRequest{
	//	Denom: denomID,
	//	ID:    strings.ToLower(s.RandStringOfLength(10)),
	//	Name:  "《富春山居图剩山图卷》",
	//	//URI:   "",
	//	Data:  "{\"holder\":\"iaa19ahdy774264ghacku03r5tdmwghgssaxdh6zqx\",\"img\":\"https://wengchangchain.oss-cn-shanghai.aliyuncs.com/artnft/home/mhfk/fcsjt/fcsjt_img.jpeg\",\"icon\":\"https://wengchangchain.oss-cn-shanghai.aliyuncs.com/artnft/home/mhfk/fcsjt/fcsjt_icon.png\",\"status\":\"active\",\"issuer\":[\"海南国际文化艺术品交易中心\"],\"union_issue\":[\"中工美工艺美术文化有限公司\",\"华泰财产保险有限公司北京分公司\",\"火凤凰（北京）国际艺术品物流有限公司\",\"北京国博文物鉴定中心\",\"全国文物流通经营机构联盟\",\"中华文化促进会古玩城联盟\",\"中国博协非国有博物馆专业委员会\",\"易拍全球\",\"中国国际商会文化和旅游产业委员会\",\"中国拍卖行业协会\",\"北京画廊协会\",\"中国文物交流中心\",\"海南省工商业联合会\",\"中国博物馆协会\",\"中国民营文化产业商会\",\"中国建设银行\",\"海南省登记结算有限责任公司\"],\"escrower\":[\"海南国际文化艺术品交易中心\"],\"description\":\"富春山居图是元代画家黄公望于1350年创作的纸本水墨画，中国十大传世名画之一。黄公望为师弟郑樗（无用师）所绘，几经易手，并因“焚画殉葬”而身首两段。前半卷：剩山图，现收藏于浙江省博物馆； 后半卷：无用师卷，现藏台北故宫博物院。\\n \\n富春山居图以浙江富春江为背景，画面用墨淡雅，山和水的布置疏密得当，墨色浓淡干湿并用，极富于变化，被誉为“画中之兰亭”，属国宝级文物。\\n\\n《富春山居图》原画画在六张纸上，六张纸接裱而成一副约七百公分的长卷。而黄公望并没有一定按着每一张纸的大小长宽构思结构，而是任凭个人的自由创作悠然于山水之间，可远观可近看。这种浏览、移动、重叠的视点，或广角深远，或推近特写，浏览过程中，视觉观看的方式极其自由无拘，角度也非常千变万化。\\n\\n作者介绍\\n\\n黄公望（1269年9月12日—1354年11月10日），元代画家。自称浙东平阳人，元顺帝至正十四年（1354年）十月二十五日逝世，享年八十六岁。\\n黄公望工书法，通音律，善诗词散曲。尤擅画山水，名列“元四家”（黄公望、吴镇、倪瓒、王蒙）之首。传世画作有《富春山居图》《水阁清幽图》《天池石壁图》《九峰雪霁图》《富春大岭图》等。著有画论《写山水诀》。\"}",
	//}
	//res, err := s.IRITAClient.NFT.MintNFT(mintNft, s.baseTx)
	//require.NoError(s.T(), err)
	//fmt.Println(res)

	//es, err := s.IRITAClient.NFT.EditNFT(nft.EditNFTRequest{
	//	Denom: denomID,
	//	Name:  nftName,
	//	ID:    "rrqugrpbhr",
	//	//URI:   "http://wcchain.bianjie.ai/#/denoms",
	//	Data:  "{\"holder\":\"iaa19ahdy774264ghacku03r5tdmwghgssaxdh6zqx\",\"img\":\"https://wengchangchain.oss-cn-shanghai.aliyuncs.com/artnft/home/mhfk/fcsjt/fcsjt_img.jpeg\",\"icon\":\"https://wengchangchain.oss-cn-shanghai.aliyuncs.com/artnft/home/mhfk/fcsjt/fcsjt_icon.png\",\"status\":\"active\",\"issuer\":[\"海南国际文化艺术品交易中心\"],\"union_issue\":[\"中工美工艺美术文化有限公司\",\"华泰财产保险有限公司北京分公司\",\"火凤凰（北京）国际艺术品物流有限公司\",\"北京国博文物鉴定中心\",\"全国文物流通经营机构联盟\",\"中华文化促进会古玩城联盟\",\"中国博协非国有博物馆专业委员会\",\"易拍全球\",\"中国国际商会文化和旅游产业委员会\",\"中国拍卖行业协会\",\"北京画廊协会\",\"中国文物交流中心\",\"海南省工商业联合会\",\"中国博物馆协会\",\"中国民营文化产业商会\",\"中国建设银行\",\"海南省登记结算有限责任公司\"],\"escrower\":[\"海南国际文化艺术品交易中心\"],\"description\":\"富春山居图是元代画家黄公望于1350年创作的纸本水墨画，中国十大传世名画之一。黄公望为师弟郑樗（无用师）所绘，几经易手，并因“焚画殉葬”而身首两段。前半卷：剩山图，现收藏于浙江省博物馆； 后半卷：无用师卷，现藏台北故宫博物院。\\n \\n富春山居图以浙江富春江为背景，画面用墨淡雅，山和水的布置疏密得当，墨色浓淡干湿并用，极富于变化，被誉为“画中之兰亭”，属国宝级文物。\\n\\n《富春山居图》原画画在六张纸上，六张纸接裱而成一副约七百公分的长卷。而黄公望并没有一定按着每一张纸的大小长宽构思结构，而是任凭个人的自由创作悠然于山水之间，可远观可近看。这种浏览、移动、重叠的视点，或广角深远，或推近特写，浏览过程中，视觉观看的方式极其自由无拘，角度也非常千变万化。\\n\\n作者介绍\\n\\n黄公望（1269年9月12日—1354年11月10日），元代画家。自称浙东平阳人，元顺帝至正十四年（1354年）十月二十五日逝世，享年八十六岁。\\n黄公望工书法，通音律，善诗词散曲。尤擅画山水，名列“元四家”（黄公望、吴镇、倪瓒、王蒙）之首。传世画作有《富春山居图》《水阁清幽图》《天池石壁图》《九峰雪霁图》《富春大岭图》等。著有画论《写山水诀》。\"}",
	//	//Data:  strings.ToLower(s.RandStringOfLength(10)),
	//}, s.baseTx)
	//require.NoError(s.T(), err)
	//fmt.Println(es)

	//res, er := s.IRITAClient.NFT.BurnNFT(nft.BurnNFTRequest{
	//	Denom: denomID,
	//	ID:    "qcmjlb",
	//}, s.baseTx)
	//require.NoError(s.T(), er)
	//fmt.Println(res)
}
