package rate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeEcbHttpClient(t *testing.T) {
	host := "demo_string"
	client := MakeCustomsGovtNZHttpClient(host)
	assert.NotEmpty(t, client.HOST)
	assert.Equal(t, host, client.HOST)
}
