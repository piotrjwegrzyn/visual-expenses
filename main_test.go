package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	log.SetFlags(0)

	tests := []struct {
		args    []string
		want    string
		comment string
	}{
		{
			args:    []string{"visual-expenses", "-v"},
			want:    "Current version:",
			comment: "Printing version",
		},
	}

	for _, tc := range tests {
		var out bytes.Buffer
		log.SetOutput(&out)
		os.Args = tc.args

		main()

		require.Equal(t, tc.want, strings.TrimSpace(out.String()))
	}
}
