package exchangeratesapiio

import (
	errors "fiatconv/pkg/exchanging"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type getterMonk struct {
	body string
}

func (g *getterMonk) Get(url string) (resp *http.Response, err error) {
	w := httptest.NewRecorder()
	io.WriteString(w, g.body)
	w.Result()
	return w.Result(), nil
}

func TestRateAPI_Rate(t *testing.T) {
	type fields struct {
		addr   string
		client getter
	}
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float32
		wantErr bool
		err     error
	}{
		{
			name: "Success responce from a monk endpoint.",
			fields: fields{
				addr:   "",
				client: &getterMonk{body: "{\"rates\":{\"RUB\":63.5200534759},\"base\":\"USD\",\"date\":\"2019-07-10\"}"},
			},
			args: args{
				src: "USD",
				dst: "RUB",
			},
			want:    63.5200534759,
			wantErr: false,
			err:     nil,
		},
		{
			name: "Error responce from a monk endpoint. A source currency is wrong.",
			fields: fields{
				addr:   "",
				client: &getterMonk{body: "{\"error\": \"Base 'USD1' is not supported.\"}"},
			},
			args: args{
				src: "USD1",
				dst: "RUB",
			},
			want:    0,
			wantErr: true,
			err:     errors.ErrSrcCurrensyNotFound,
		},
		{
			name: "Error responce from a monk endpoint. A destination currency is wrong.",
			fields: fields{
				addr:   "",
				client: &getterMonk{body: "{\"error\": \"Symbols 'usd1' are invalid for date 2019-07-10.\"}"},
			},
			args: args{
				src: "USD",
				dst: "RUB1",
			},
			want:    0,
			wantErr: true,
			err:     errors.ErrDstCurrensyNotFound,
		},
		{
			name: "Error responce when the service unavailable",
			fields: fields{
				addr:   addr,
				client: &http.Client{Timeout: time.Duration(1 * time.Microsecond)},
			},
			args: args{
				src: "USD",
				dst: "RUB1",
			},
			want:    0,
			wantErr: true,
			err:     errors.ErrRateUnavailable,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &RateAPI{
				addr:   tt.fields.addr,
				client: tt.fields.client,
			}
			got, err := api.Rate(tt.args.src, tt.args.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("RateAPI.Rate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && (err != tt.err) {
				t.Errorf("RateAPI.Rate() \nerror = %v, \nwantErr = %v", err, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("RateAPI.Rate() = %v, want %v", got, tt.want)
			}
		})
	}
}
