package zilla

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

var skuDetailsRe = regexp.MustCompile(`^\{"\d+.*options_description.*}`)

func init() {
	skuDetailsRe.Longest()
}

type tag struct {
	name  string
	attrs map[string]string
}

type Parser struct {
	tokenizer *html.Tokenizer
	tags      []tag

	products   []Product
	skuDetails []SkuDetails
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		tokenizer: html.NewTokenizer(r),
	}
}

func (p *Parser) Parse() (ParseResult, error) {
	for {
		tt := p.tokenizer.Next()
		switch tt {
		case html.StartTagToken:
			p.handleStartTag()
		case html.EndTagToken:
			p.handleEndTag()
		case html.TextToken:
			p.handleText()
		case html.ErrorToken:
			err := p.tokenizer.Err()
			if !errors.Is(err, io.EOF) {
				return ParseResult{}, err
			}

			return ParseResult{
				Products:   p.products,
				SkuDetails: p.skuDetails,
			}, nil
		}
	}
}

func (p *Parser) handleStartTag() {
	tn, hasAttr := p.tokenizer.TagName()

	tg := tag{
		name:  string(tn),
		attrs: map[string]string{},
	}

	if hasAttr {
		for {
			key, val, moreAttr := p.tokenizer.TagAttr()
			skey := string(key)
			sval := string(val)

			tg.attrs[skey] = sval
			if !moreAttr {
				break
			}
		}
	}

	p.tags = append(p.tags, tg)
}

func (p *Parser) handleEndTag() {
	p.tags = p.tags[:len(p.tags)-1]
}

func (p *Parser) handleText() {
	if len(p.tags) == 0 {
		return
	}

	lastTag := p.tags[len(p.tags)-1]

	if lastTag.name != "script" {
		return
	}

	scriptType := lastTag.attrs["type"]

	if scriptType == "application/ld+json" {
		p.handleJsonLd()
	} else if strings.HasSuffix(scriptType, "text/javascript") {
		p.handleJavascript()
	}
}

func (p *Parser) handleJsonLd() {
	var products []Product
	err := json.Unmarshal(p.tokenizer.Text(), &products)
	if err == nil {
		p.products = products
		return
	}
}

func (p *Parser) handleJavascript() {
	lines := bytes.Split(p.tokenizer.Text(), []byte{'\n'})
	for _, line := range lines {
		if match := skuDetailsRe.Find(line); match != nil {
			var skuDetails map[string]SkuDetails
			err := json.Unmarshal(match, &skuDetails)
			if err == nil {
				for _, skuDetail := range skuDetails {
					p.skuDetails = append(p.skuDetails, skuDetail)
				}

				sort.Slice(p.skuDetails, func(i, j int) bool {
					return p.skuDetails[i].Id < p.skuDetails[j].Id
				})

				return
			}
		}
	}
}
