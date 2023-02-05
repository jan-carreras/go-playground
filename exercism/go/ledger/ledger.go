package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var ErrUnsupportedCurrency = errors.New("unsupported currency")
var ErrUnsupportedLocale = errors.New("unsupported locale")
var ErrInvalidDate = errors.New("invalid date")

const (
	localeNlNL = "nl-NL"
	localeEnUS = "en-US"
)

const (
	currencyEUR = "EUR"
	currencyUSD = "USD"
)

var supportedLocales = map[string]bool{
	localeNlNL: true,
	localeEnUS: true,
}

var currencyMapping = map[string]string{
	currencyEUR: "€",
	currencyUSD: "$",
}

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type money struct {
	amount   int // in cents
	currency string
	negative bool
}

func newMoney(cents int, currency string) money {
	negative := cents < 0
	if negative {

		cents *= -1
	}

	return money{
		amount:   cents,
		currency: currency,
		negative: negative,
	}
}

func (m money) Format(locale string) (string, error) {
	switch locale {
	case localeNlNL:
		return m.stringNL(), nil
	case localeEnUS:
		return m.stringEN(), nil
	}

	return "", ErrUnsupportedLocale
}

func (m money) stringEN() string {
	amount := m.addDelimiters(m.getAmount(), ",")

	s := fmt.Sprintf("%s%s.%02d", m.currencySymbol(), amount, m.amount%100)
	if m.negative {
		s = fmt.Sprintf("(%s)", s)
	} else {
		s = fmt.Sprintf("%s ", s)
	}
	return s
}

func (m money) getAmount() string {
	return strconv.Itoa(m.amount / 100)
}

func (m money) addDelimiters(in, delimiter string) string {
	out := make([]string, 0)
	for i := len(in); i > 0; i -= 3 {
		start := max(i-3, 0)
		out = append([]string{in[start:i]}, out...)
	}

	return strings.Join(out, delimiter)
}

func (m money) stringNL() string {
	amount := m.addDelimiters(m.getAmount(), ".")

	s := fmt.Sprintf("%s %s,%02d", m.currencySymbol(), amount, m.amount%100)
	if m.negative {
		s = fmt.Sprintf("%s-", s)
	} else {
		s = fmt.Sprintf("%s ", s)
	}
	return s
}

func (m money) currencySymbol() string {
	return currencyMapping[m.currency]
}

func FormatLedger(currency string, locale string, originalEntries []Entry) (string, error) {
	// We don't want to modify the original entries
	entries := make([]Entry, len(originalEntries))
	copy(entries, originalEntries)

	return printTable(currency, locale, entries)
}

func printTable(currency string, locale string, entries []Entry) (string, error) {
	// Sanity checks
	if _, ok := currencyMapping[currency]; !ok {
		return "", ErrUnsupportedCurrency
	} else if _, ok := supportedLocales[locale]; !ok {
		return "", ErrUnsupportedLocale
	}

	output := &strings.Builder{}

	if err := printHeader(locale, output); err != nil {
		return "", err
	}

	for _, entry := range sortEntries(entries) {
		if err := printEntry(currency, locale, entry, output); err != nil {
			return "", err
		}
	}

	return output.String(), nil
}

func printHeader(locale string, output *strings.Builder) error {
	header, err := getLocalisedHeader(locale)
	if err != nil {
		return err
	}
	output.WriteString(header)
	return nil
}

func printEntry(currency string, locale string, entry Entry, output *strings.Builder) error {
	// Sanity check
	if _, ok := currencyMapping[currency]; !ok {
		return ErrUnsupportedCurrency
	}

	date, err := parseDate(entry.Date, locale)
	if err != nil {
		return err
	}

	formattedAmount, err := newMoney(entry.Change, currency).Format(locale)
	if err != nil {
		return err
	}

	row := fmt.Sprintf("%-10s | %-25s | %13s\n", date, formatDescription(entry.Description), formattedAmount)

	output.WriteString(row)
	
	return nil
}

func formatDescription(description string) string {
	if len(description) > 25 {
		return description[:22] + "..."
	}
	return description
}

func sortEntries(entries []Entry) []Entry {
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Date == entries[j].Date {
			return entries[i].Change < entries[j].Change
		}
		return entries[i].Date < entries[j].Date
	})

	return entries
}

func parseDate(date string, locale string) (string, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", ErrInvalidDate
	}

	switch locale {
	case localeNlNL:
		return t.Format("02-01-2006"), nil
	case localeEnUS:
		return t.Format("01/02/2006"), nil
	}

	return "", ErrUnsupportedLocale
}

func getLocalisedHeader(locale string) (string, error) {
	tplColumns := "%-10s | %-25s | %s\n"
	switch locale {
	case localeNlNL:
		return fmt.Sprintf(tplColumns, "Datum", "Omschrijving", "Verandering"), nil
	case localeEnUS:
		return fmt.Sprintf(tplColumns, "Date", "Description", "Change"), nil
	}

	return "", ErrUnsupportedLocale
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
