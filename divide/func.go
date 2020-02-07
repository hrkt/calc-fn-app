package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	
	"github.com/cockroachdb/apd"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

type CalcForm struct {
	Left string `json:"left"`
	Right string `json:"right"`
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	form := &CalcForm{Left: "0", Right: "1"}
	json.NewDecoder(in).Decode(form)

	left, _, err := apd.NewFromString(form.Left)
	right, _, err := apd.NewFromString(form.Right)
	d := new(apd.Decimal)
	c := apd.BaseContext
	c.Precision = 28
	_, err = c.Quo(d, left, right)

	var resString string
	if err != nil {
        resString = fmt.Sprintf(", error: %s", err)
    } else {
        resString = fmt.Sprintf("%s", d)
    }

	calcResult := struct {
		Result string `json:"result"`
	}{
		Result: resString,
	}
	json.NewEncoder(out).Encode(&calcResult)
}
