package core

import (
	"errors"
	"math/big"
	"sync"

	"github.com/thetatoken/theta/common"
	ts "github.com/thetatoken/theta/store"
	"github.com/thetatoken/theta/store/database"
	"github.com/thetatoken/theta/store/kvstore"
)

// ------------------------------------ Transaction Cache ----------------------------------------------

type Transactions struct {
	Index         *big.Int
	TxHash        [32]byte
	ChainId       *big.Int
	SignedTxData  []byte
	Height        *big.Int
	Validator     common.Address
	RpcUrl        string
	ProxyContract common.Address
	Value         *big.Int
	Sender        common.Address
	Target        common.Address
}

var (
	ErrTransactionNotFound      = errors.New("TransactionNotFound")
	ErrTransactionExisted       = errors.New("TransactionExisted")
	ErrTransactionPersistFailed = errors.New("TransactionPersistFailed")
)

// TransactionIndexKey constructs the DB key for the given transaction hash.
func TransactionIndexKey(index *big.Int) common.Bytes {
	return common.Bytes("tx/" + index.String())
}

type TransactionCache struct {
	mutex *sync.Mutex // mutex for concurrency protection
	db    database.Database
}

// NewTransactionCache creates a new transaction cache instance.
func NewTransactionCache(db database.Database) *TransactionCache {
	cache := &TransactionCache{
		mutex: &sync.Mutex{},
		db:    db,
	}
	return cache
}

func (c *TransactionCache) InsertList(tx []*Transactions) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	for _, tx := range tx {
		err := store.Put(TransactionIndexKey(tx.Index), tx)
		if err != nil {
			return err
		}
	}
	return nil
}

// a function that changes already inserted transaction
func (c *TransactionCache) Update(tx *Transactions) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Put(TransactionIndexKey(tx.Index), tx)
	return err // the caller should handle the error
}

func (c *TransactionCache) Insert(tx *Transactions) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Put(TransactionIndexKey(tx.Index), tx)
	logger.Infof("Inserting transaction: ", tx.Index, " ", tx.TxHash, " ", tx.ChainId, " ", tx.SignedTxData, " ", tx.Height, " ", tx.Validator, " ", tx.RpcUrl, " ", tx.ProxyContract, " ", tx.Value, " ", tx.Sender, " ", tx.Target)
	return err // the caller should handle the error
}

func (c *TransactionCache) Delete(Index *big.Int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	store := kvstore.NewKVStore(c.db)
	err := store.Delete(TransactionIndexKey(Index))
	return err // the caller should handle the error
}

func (c *TransactionCache) Get(Index *big.Int) (*Transactions, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	tx := Transactions{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(TransactionIndexKey(Index), &tx)
	return &tx, err // the caller should handle the error
}

func (c *TransactionCache) Exists(Index *big.Int) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	tx := Transactions{}
	store := kvstore.NewKVStore(c.db)
	err := store.Get(TransactionIndexKey(Index), &tx)
	if err == nil {
		return true, nil
	}

	if err == ts.ErrKeyNotFound {
		return false, nil
	}

	return false, err // the caller should handle the error
}
