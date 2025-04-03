package db

import (
	"github.com/go-pg/pg"
	"time"
)

type Transaction struct {
	
    account_id 							int				`json:"id"` // filler for now but this is how to do it for json...still need to figure out exporting
    plaid_transaction_id    string		`json:"trans_id"`
    //transaction_date DATE
    //post_date DATE
    //merchant_name VARCHAR(255)
    //amount DECIMAL(19, 4) 
    //currency_code CHAR(3) 
    //transaction_type VARCHAR(50)
    //is_pending BOOLEAN 
    //description TEXT,
    //original_description TEXT,
    //is_manual BOOLEAN 
    //created_at TIMESTAMP 
    //updated_at TIMESTAMP 
}

func GetTransactions(db *pg.DB) ([]Transaction, error) {
	var transactions []Transaction
	err := db.Model(&transactions).Select()
	return transactions, err
}
