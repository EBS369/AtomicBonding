// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"main/AtomicBond"
	"main/IERC20"
	"main/JoePair"
	"main/TimeBondDepository"
	"math"
	"math/big"
	"os"
	"os/signal"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-co-op/gocron"
)

type Account struct {
	PrivateKey     *ecdsa.PrivateKey
	PublicKeyECDSA *ecdsa.PublicKey
	FromAddress    common.Address
}

const (
	Min_ROI = int64(6_90) // 6.90%
)

var (
	AtomicBondAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")              // Redacted
	CallerAccount     = GetAccount("4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d") // Example Private Key
	FunderAddress     = common.HexToAddress("0x0000000000000000000000000000000000000000")              // Redacted
	MinterAddress     = common.HexToAddress("0x0000000000000000000000000000000000000000")              // Redacted
)

const (
	RPC_HTTPS = "https://api.avax.network/ext/bc/C/rpc"
	RPC_WSS   = "wss://api.avax.network/ext/bc/C/ws"
)

var (
	ChainId                      = big.NewInt(0xa86a)
	MimTimeBondDepositoryAddress = common.HexToAddress("0x694738E0A438d90487b4a549b201142c1a97B556")
	MemoAddress                  = common.HexToAddress("0x136acd46c134e8269052c62a67042d6bdedde3c9")
	MimTimeJoePairAddress        = common.HexToAddress("0x113f413371fC4CC4C9d6416cf1DE9dFd7BF747Df")
	MimTimeSwapPath              = []common.Address{
		common.HexToAddress("0xb54f16fB19478766A268F172C9480f8da1a7c9C3"),
		common.HexToAddress("0x130966628846BFd36ff31a822705796e8cb8C18D"),
	}
)

var (
	int0     = big.NewInt(0)
	int997   = big.NewInt(997)
	int1000  = big.NewInt(1000)
	int10000 = big.NewInt(10000)
	int10e16 = big.NewInt(10_000_000_000_000_000)
	Gwei10   = big.NewInt(10_000_000_000)
)

// --- --- ---

var (
	balance       *big.Int
	balanceUint64 uint64
	balanceFloat  *big.Float
	balanceString string
)

// --- --- ---

func HTTPSClient() *ethclient.Client {
	client, err := ethclient.Dial(RPC_HTTPS)
	if err != nil {
		log.Println(err)
		return nil
	}
	return client
}

func WSSClient() *ethclient.Client {
	client, err := ethclient.Dial(RPC_WSS)
	if err != nil {
		log.Println(err)
		return nil
	}
	return client
}

// --- --- ---

func GetAccount(pk string) *Account {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Println(err)
		return nil
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Account{
		PrivateKey:     privateKey,
		PublicKeyECDSA: publicKeyECDSA,
		FromAddress:    fromAddress,
	}
}

// --- --- ---

// https://ethereum.stackexchange.com/a/83702
func EstimateMinAmountOut(srcReserve *big.Int, destReserve *big.Int, srcAmount *big.Int) *big.Int {
	var Y, X, minAmountOut big.Int
	Y.Mul(destReserve, srcAmount)
	X.Add(srcReserve, srcAmount)
	minAmountOut.Div(&Y, &X)
	// 0.3% Platform / LP Fees
	return minAmountOut.Div(minAmountOut.Mul(&minAmountOut, int997), int1000)
}

// --- --- ---

func RunAtomicBond(
	gasPriceOffset *big.Int,
	funder common.Address,
	minter common.Address,
	bondIn *big.Int,
	swapPath []common.Address,
	minSwapAmount *big.Int,
	maxBondPrice *big.Int,
	deadlineOffset int64,
) bool {
	client := HTTPSClient()
	if client == nil {
		return false
	}
	defer client.Close()

	nonce, err := client.PendingNonceAt(context.Background(), CallerAccount.FromAddress)
	if err != nil {
		log.Println(err)
		return false
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	gasPrice.Add(gasPrice, gasPriceOffset)
	if err != nil {
		log.Println(err)
		return false
	}

	auth, err := bind.NewKeyedTransactorWithChainID(CallerAccount.PrivateKey, ChainId)
	if err != nil {
		log.Println(err)
		return false
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = int0
	auth.GasLimit = uint64(512000)
	auth.GasPrice = gasPrice

	instance, err := AtomicBond.NewAtomicBond(AtomicBondAddress, client)
	if err != nil {
		log.Println(err)
		return false
	}

	tx, err := instance.AtomicBondMemoMim(auth, funder, minter, bondIn, swapPath, minSwapAmount, maxBondPrice, big.NewInt(time.Now().UnixMilli()+deadlineOffset))
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println(
		bondIn,
		minSwapAmount,
		maxBondPrice,
	)

	log.Println("[TX] " + tx.Hash().Hex())

	return true
}

func RunRedeem(
	gasPriceOffset *big.Int,
	minter common.Address,
) bool {
	client := HTTPSClient()
	if client == nil {
		return false
	}
	defer client.Close()

	nonce, err := client.PendingNonceAt(context.Background(), CallerAccount.FromAddress)
	if err != nil {
		log.Println(err)
		return false
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	gasPrice.Add(gasPrice, gasPriceOffset)
	if err != nil {
		log.Println(err)
		return false
	}

	auth, err := bind.NewKeyedTransactorWithChainID(CallerAccount.PrivateKey, ChainId)
	if err != nil {
		log.Println(err)
		return false
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = int0
	auth.GasLimit = uint64(320000)
	auth.GasPrice = gasPrice

	instance, err := AtomicBond.NewAtomicBond(AtomicBondAddress, client)
	if err != nil {
		log.Println(err)
		return false
	}

	tx, err := instance.Redeem(auth, minter)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("[TX] " + tx.Hash().Hex())

	return true
}

// --- --- ---

func renewBalance(client *ethclient.Client) {
	log.Println(" --- renew balance --- ")

	memoInstance, err := IERC20.NewIERC20(MemoAddress, client)
	if err != nil {
		log.Println(err)
		return
	}

	balance, err = memoInstance.BalanceOf(&bind.CallOpts{}, FunderAddress)
	if err != nil {
		log.Println(err)
		return
	}
	balanceUint64 = balance.Uint64()

	var ok bool
	balanceFloat, ok = new(big.Float).SetString(balance.String())
	if !ok {
		log.Println("cannot cast balance of type *big.Int as type *big.Float")
		return
	}
	balanceFloat.Quo(balanceFloat, big.NewFloat(math.Pow10(int(9))))
	balanceString = balanceFloat.String()
}

func redeem() {
	RunRedeem(
		Gwei10,
		MinterAddress,
	)
}

func main() {
	go func() {
		var err error

		client := HTTPSClient()
		defer func() {
			if client != nil {
				client.Close()
			}
		}()

		s := gocron.NewScheduler(time.Local)
		_, err = s.Cron("56 21 * * *").Do(redeem)
		if err != nil {
			log.Println(err)
		}
		_, err = s.Cron("56 13 * * *").Do(redeem)
		if err != nil {
			log.Println(err)
		}
		_, err = s.Cron("56 5 * * *").Do(redeem)
		if err != nil {
			log.Println(err)
		}
		s.Every("15s").Do(func() { renewBalance(client) })
		s.StartAsync()
	}()

	go func() {
		client := HTTPSClient()
		defer func() {
			if client != nil {
				client.Close()
			}
		}()

		joePairInstance, err := JoePair.NewJoePair(MimTimeJoePairAddress, client)
		if err != nil {
			log.Println(err)
			return
		}
		timeBondDepositoryInstance, err := TimeBondDepository.NewTimeBondDepository(MimTimeBondDepositoryAddress, client)
		if err != nil {
			log.Println(err)
			return
		}

		wssClient := WSSClient()
		defer func() {
			if wssClient != nil {
				defer wssClient.Close()
			}
		}()

		headers := make(chan *types.Header)
		defer close(headers)

		sub, err := wssClient.SubscribeNewHead(context.Background(), headers)
		if err != nil {
			log.Println(err)
			return
		}
		defer sub.Unsubscribe()

		var minSwapAmount *big.Int
		var priceFloat big.Float
		var price, roi big.Int
		var roiInt64 int64
		for {
			select {
			case err := <-sub.Err():
				log.Println(err)
				return
			case <-headers:
				if balanceUint64 == 0 {
					log.Println("balance is zero")
					continue
				}

				reserves, err := joePairInstance.GetReserves(&bind.CallOpts{})
				if err != nil {
					log.Println(err)
					continue
				}

				bondPriceInUSD, err := timeBondDepositoryInstance.BondPriceInUSD(&bind.CallOpts{})
				if err != nil {
					log.Println(err)
					continue
				}

				minSwapAmount = EstimateMinAmountOut(reserves.Reserve1, reserves.Reserve0, balance)

				minSwapAmountFloat, ok := new(big.Float).SetString(minSwapAmount.String())
				if !ok {
					log.Println("cannot cast EstimateMinAmountOut() output of type *big.Int as type *big.Float")
					continue
				}

				priceFloat.Quo(minSwapAmountFloat, balanceFloat).Int(&price)
				roi.Sub(&price, bondPriceInUSD)
				roi.Mul(&roi, int10000)
				roi.Div(&roi, bondPriceInUSD)
				roiInt64 = roi.Int64()

				log.Println(balanceString, "$TIME | $MIM", roi.String(), "/", Min_ROI, "|", bondPriceInUSD.String(), "/", price.String())

				if roiInt64 >= Min_ROI && roiInt64 < 5000 {
					log.Println("Bond!")

					RunAtomicBond(
						Gwei10,
						FunderAddress,
						MinterAddress,
						balance,
						MimTimeSwapPath,
						minSwapAmount,
						bondPriceInUSD.Div(bondPriceInUSD, int10e16),
						15,
					)

					time.Sleep(15 * time.Second)

					renewBalance(client)

					return
				}
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
