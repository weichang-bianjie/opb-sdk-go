package integration_test

import (
	"context"
	"fmt"
	sdk "github.com/bianjieai/irita-sdk-go"
	"github.com/bianjieai/irita-sdk-go/crypto"
	cryptoAmino "github.com/bianjieai/irita-sdk-go/crypto/codec"
	"github.com/bianjieai/irita-sdk-go/types"
	"github.com/bianjieai/irita-sdk-go/types/store"
	"github.com/bianjieai/irita-sdk-go/utils/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
	"time"
)

const (
	user     = "user"
	password = "password"

	nodeURI  = "https://opbt.bsngate.com:18602/api/b26b8b53e1b74aa6974ab207e25310f0/rpc"
	wsAddr   = "wss://opbt.bsngate.com:18602/api/b26b8b53e1b74aa6974ab207e25310f0/ws"
	grpcAddr = "opbt.bsngate.com:18603"
	chainID  = "wenchangchain"
	charset  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	mnemonic = "pull eagle capital skin tackle gather session stock drill left already bulk divide midnight general exhaust apple bind turkey evoke must isolate inner rose"
)

type IntegrationTestSuite struct {
	suite.Suite
	sdk.IRITAClient
	r      *rand.Rand
	baseTx types.BaseTx
}

func TestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	keyDao := store.NewMemory(nil)
	getkeyDaoByMnemonicKey(keyDao)
	options := []types.Option{
		types.KeyDAOOption(keyDao),
		types.TimeoutOption(20),
		types.CachedOption(true),
	}
	cfg, err := types.NewClientConfig(nodeURI, wsAddr, grpcAddr, chainID, options...)
	if err != nil {
		panic(err)
	}

	s.IRITAClient = sdk.NewIRITAClient(cfg)
	s.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	s.SetLogger(log.NewLogger(log.Config{
		Format: log.FormatText,
		Level:  log.DebugLevel,
	}))
	res, err := s.IRITAClient.Status(context.Background())
	require.NoError(s.T(), err)
	fmt.Println(res)

	fee, err := types.ParseDecCoins("300000uirita")
	require.NoError(s.T(), err)

	s.baseTx = types.BaseTx{
		From:     user,
		Gas:      200000,
		Memo:     "test",
		Mode:     types.Commit,
		Password: password,
		Fee:      fee,
	}
}

// RandStringOfLength return a random string
func (s *IntegrationTestSuite) RandStringOfLength(l int) string {
	var result []byte
	bytes := []byte(charset)
	for i := 0; i < l; i++ {
		result = append(result, bytes[s.r.Intn(len(bytes))])
	}
	return string(result)
}

func getkeyDaoByMnemonicKey(keyDao store.MemoryDAO) {

	km, err := crypto.NewMnemonicKeyManager(mnemonic, "sm2")
	if err != nil {
		panic("recover mnemonic fail" + err.Error())
	}
	_, priv := km.Generate()
	ki := store.KeyInfo{
		Name:         user,
		PubKey:       cryptoAmino.MarshalPubkey(km.ExportPubKey()),
		PrivKeyArmor: string(cryptoAmino.MarshalPrivKey(priv)),
		Algo:         "sm2",
	}
	if err := keyDao.Write(user, password, ki); err != nil {
		panic(err)
	}
	DefaultKmAddr := types.AccAddress(km.ExportPubKey().Address()).String()
	fmt.Println("Provider: " + DefaultKmAddr)
}
