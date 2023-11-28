package tfout2env_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/llonchj/tfout2env"
	"github.com/llonchj/tfout2env/internal/helpers"
)

func TestNewTFOut(t *testing.T) {
	for name, tt := range map[string]struct {
		Input    []byte
		Expected *tfout2env.TFOut
		Err      error
	}{
		"empty": {Err: io.EOF},
		"basic": {
			Input: []byte(`{"a":{
				"sensitive": false,
				"type": "string",
				"value": "my_value"},
				"web_server_count": {
					"sensitive": false,
					"type": "number",
					"value": 4
				  }}`),
			Expected: &tfout2env.TFOut{
				"a":                {Datum: "my_value", Type: "string"},
				"web_server_count": {Datum: float64(4), Type: "number"},
			},
		},
	} {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			Output, err := tfout2env.New(bytes.NewBuffer(tt.Input))
			if err != nil || tt.Err != nil {
				if diff := cmp.Diff(tt.Err, err, helpers.EquateErrorContent()); diff != "" {
					t.Errorf("error do not match: (-expected +got)\n%s", diff)
				}
				return
			}

			if diff := cmp.Diff(tt.Expected, Output); diff != "" {
				t.Errorf("do not match: (-expected +got)\n%s", diff)
			}
		})
	}
}

func TestTFOutStringer(t *testing.T) {
	for name, tt := range map[string]struct {
		Input    *tfout2env.TFOut
		Expected string
	}{
		"basic": {
			Input: &tfout2env.TFOut{
				"a":                {Datum: "my_value", Type: "string"},
				"b_1":              {Datum: "data", Type: "string"},
				"web_server_count": {Datum: float64(4), Type: "number"},
			},
			Expected: "A=my_value\nB_1=data\nWEB_SERVER_COUNT=4\n",
		},
	} {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			Output := tt.Input.String()
			if diff := cmp.Diff(tt.Expected, Output); diff != "" {
				t.Errorf("do not match: (-expected +got)\n%s", diff)
			}
		})
	}
}
