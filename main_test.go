package main

import (
	"bytes"
	"html/template"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashPassword(tt.args.password); got != tt.want {
				t.Errorf("HashPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswordHash(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateBTCa(t *testing.T) {
	type args struct {
		tx string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateBTCa(tt.args.tx); got != tt.want {
				t.Errorf("validateBTCa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashthis(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashthis(tt.args.password); got != tt.want {
				t.Errorf("hashthis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNewAdressBTC(t *testing.T) {
	type args struct {
		label string
		uuid  string
		pwd   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNewAdressBTC(tt.args.label, tt.args.uuid, tt.args.pwd); got != tt.want {
				t.Errorf("getNewAdressBTC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBalanceBTC(t *testing.T) {
	type args struct {
		address string
		UID     string
		pwd     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBalanceBTC(tt.args.address, tt.args.UID, tt.args.pwd); got != tt.want {
				t.Errorf("getBalanceBTC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_payAdressBTC(t *testing.T) {
	type args struct {
		UID    string
		pwd    string
		outAdr string
		amount string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := payAdressBTC(tt.args.UID, tt.args.pwd, tt.args.outAdr, tt.args.amount); got != tt.want {
				t.Errorf("payAdressBTC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplate_Render(t *testing.T) {
	type fields struct {
		templates *template.Template
	}
	type args struct {
		name string
		data interface{}
		c    echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Template{
				templates: tt.fields.templates,
			}
			w := &bytes.Buffer{}
			if err := tr.Render(w, tt.args.name, tt.args.data, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Template.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Template.Render() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestInDex(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InDex(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("InDex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenIndex(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenIndex(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GenIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenFinal(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenFinal(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GenFinal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_transPanel(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := transPanel(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("transPanel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userbPanel(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := userbPanel(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("userbPanel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userbPanelGen(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := userbPanelGen(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("userbPanelGen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userbPanelRelease(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := userbPanelRelease(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("userbPanelRelease() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userbPanelReleaseFinal(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := userbPanelReleaseFinal(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("userbPanelReleaseFinal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_adminuserbPanelReleaseFinal(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := adminuserbPanelReleaseFinal(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("adminuserbPanelReleaseFinal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_adminPanel(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := adminPanel(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("adminPanel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_adminPanelSettings(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := adminPanelSettings(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("adminPanelSettings() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_adminPanelSettingsChange(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := adminPanelSettingsChange(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("adminPanelSettingsChange() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
