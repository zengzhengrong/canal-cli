package canal

import (
	"encoding/json"

	pbe "github.com/withlin/canal-go/protocol/entry"
)

func Format(row *pbe.RowData, databaseName string, tabelName string, eventType pbe.EventType) []byte {
	msg := make(map[string]interface{})
	msg["database"] = databaseName
	msg["table"] = tabelName
	msg["type"] = eventType
	msg["data"] = make(map[string]interface{})
	data := msg["data"].(map[string]interface{})
	if eventType == pbe.EventType_UPDATE {
		for _, colum := range row.GetAfterColumns() {
			if colum.Name == "id" {
				// default add id
				data[colum.Name] = colum.Value
			}
			if colum.GetUpdated() {
				// update
				data[colum.Name] = colum.Value
			}

		}
	} else if eventType == pbe.EventType_DELETE {

		for _, colum := range row.GetBeforeColumns() {
			// delete or insert
			data[colum.Name] = colum.Value

		}

	} else if eventType == pbe.EventType_INSERT {
		for _, colum := range row.GetAfterColumns() {
			// delete or insert
			data[colum.Name] = colum.Value

		}
	} else {
		for _, colum := range row.GetAfterColumns() {
			data[colum.Name] = colum.Value

		}
	}

	msgByte, _ := json.Marshal(msg)
	return msgByte
}
