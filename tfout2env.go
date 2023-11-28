package tfout2env

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Value struct {
	Datum     any    `json:"value"`
	Sensitive bool   `json:"sensitive"`
	Type      string `json:"type"`
}

type TFOut map[string]Value

func (v Value) String() string {
	return fmt.Sprintf("%v", v.Datum)
}

func New(r io.Reader) (*TFOut, error) {
	o := TFOut{}
	if err := json.NewDecoder(r).Decode(&o); err != nil {
		return nil, err
	}
	return &o, nil
}

// String returns contents in a shell env format.
func (t TFOut) String() string {
	// sort by keys
	keys := make([]string, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	for _, k := range keys {
		b.WriteString(strings.ToUpper(k))
		b.WriteString("=")
		b.WriteString(t[k].String())
		b.WriteString("\n")
	}
	return b.String()
}

// Write dumps the content into w.
func (t TFOut) Write(w io.Writer) (int, error) {
	return w.Write([]byte(t.String()))
}
