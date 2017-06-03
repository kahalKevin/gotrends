package gotrends

import (
	"log"
	"strings"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"strconv"
)

type ValueTrend struct {
	V 	interface{}		`json:"v"`
	F 	string 			`json:"f,omitempty"`
}
type ValuesTrends []ValueTrend

type ComponentTrend struct {
	AllValues 	ValuesTrends 	`json:"c"`
}
type ComponentsTrends []ComponentTrend

type Row struct {
	Row 	ComponentsTrends 	`json:"rows"`
}

type Whole struct {
	Table 	Row 	`json:"table"`
}

func SearchWithKeyword(keyword string) (string, int){

	googleTrendURL := "https://trends.google.com/trends/fetchComponent"
	cat := "18" //shopping
	date := "today 12-m" //12 months past
	geo := "ID" //Indonesia as region of search
	query := keyword //keyword in query
	cid := "TOP_QUERIES_0_0" //one op the option, RISING also available | also TIMESERIES_GRAPH_0
	export := "3" //to exported to 'almost' json like | can use 5 to get a HTML 5 based return

	req, _ := http.NewRequest("GET", googleTrendURL, nil)	
	q := req.URL.Query()
	q.Add("cat", cat)
	q.Add("date", date)
	q.Add("geo", geo)
	q.Add("q", query)
	q.Add("cid", cid)
	q.Add("export", export)
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
	log.Printf(req.URL.String())
	if err != nil {
		log.Printf("1 error: %v", err)
	}
	defer resp.Body.Close()
	resultBytes, _ := ioutil.ReadAll(resp.Body)
	resultString := string(resultBytes)
	// log.Printf(resultString)
	switch {
		case strings.Contains(resultString,`"status":"error"`): //check if error
			log.Printf("error in trend")
			return "", 0
		case strings.Contains(resultString,`"rows":[`): //check if success
			log.Printf("success in trend")
			break
		default : // will be considered error if none of above
			log.Printf("limit exceeded")
			return "", 0
	}

	resultString = resultString[62:] //remove unnecessary string from upfront
	resultString = resultString[:len(resultString)-2] //remove unnecessary string from behind

	var TableData Whole
	err = json.Unmarshal([]byte(resultString), &TableData)
	if(err != nil){
		log.Printf("error: %v", err)
	}

	keywordOutput := ""
	keywordScore := 0.0
	for no := range TableData.Table.Row {
		_keywordOutput := TableData.Table.Row[no].AllValues[0].V.(string)
		_keywordScore, _ :=  strconv.ParseFloat(TableData.Table.Row[no].AllValues[1].F, 64)

		if _keywordScore > keywordScore {
			keywordScore = _keywordScore
			keywordOutput = _keywordOutput
		}else if _keywordScore == keywordScore{
			if len(_keywordOutput) > len(keywordOutput){
				keywordScore = _keywordScore
				keywordOutput = _keywordOutput
			}
		}
	}

	return keywordOutput , int(keywordScore)
}
