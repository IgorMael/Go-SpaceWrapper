/*
Package spacewrapper implements
an wrapper to access ocr.space OCR service
*/
package spacewrapper

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_buildForm(t *testing.T) {
	type args struct {
		paramTexts Params
		file       File
	}
	tests := []struct {
		name    string
		args    args
		want    *bytes.Buffer
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := buildForm(tt.args.paramTexts, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildForm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildForm() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("buildForm() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
