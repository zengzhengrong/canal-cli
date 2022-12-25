package canal

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	pbe "github.com/withlin/canal-go/protocol/entry"
)

// Watch is Listen Canal https://github.com/withlin/canal-go/blob/master/samples/main.go
func Watch() {
	for {

		message, err := canalClient.Get(100, nil, nil)
		if err != nil {
			logrus.Error(err)
		}
		if message.Id == -1 || len(message.Entries) <= 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		for _, entry := range message.Entries {
			if entry.GetEntryType() == pbe.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == pbe.EntryType_TRANSACTIONEND {
				continue
			}
			rowChange := new(pbe.RowChange)

			if err := proto.Unmarshal(entry.GetStoreValue(), rowChange); err != nil {
				logrus.Error(err)
			}

			header := entry.GetHeader()
			for _, rowData := range rowChange.GetRowDatas() {
				data := Format(rowData, header.GetSchemaName(), header.GetTableName(), header.GetEventType())
				logrus.Info(string(data))
			}

		}
	}
}
