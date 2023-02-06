package main

import (
	"net/http"
	"os"

	"github.com/cbguder/revzilla/zilla"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	err := fetch(os.Args[1])
	if err != nil {
		panic(err)
	}
}

func fetch(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	p := zilla.NewParser(resp.Body)
	res, err := p.Parse()
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.Style().Format.Header = text.FormatDefault

	t.SetTitle(res.Products[0].Name)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Color / Size", "MSRP", "Retail", "Members", "In Stock?"})
	for _, sku := range res.SkuDetails {
		t.AppendRow(table.Row{
			sku.OptionsDescription,
			sku.Msrp,
			sku.Retail,
			sku.LoyaltyPrice,
			sku.InStock,
		})
	}
	t.Render()

	return nil
}
