package helper

import (
	"os"
	"strconv"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func Payment(Name, Email string, Total int) (snapResp *snap.Response) {
	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// generate order id
	orderIDGenerate := strconv.FormatInt(time.Now().Unix(), 10)

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderIDGenerate,
			GrossAmt: int64(Total),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: Name,
			Email: Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ = s.CreateTransaction(req)

	return snapResp
}
