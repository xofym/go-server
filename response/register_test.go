// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xofym/go-server/response/encoder"
)

func TestRegister(t *testing.T) {

	assert.Equal(t, 0, len(encoders))

	provider := &encoder.MockEncoder{}

	provider.On("MimeType").Return("application/json")

	Register(provider)

	assert.Equal(t, 1, len(encoders))

	encoders = make(map[string]encoder.Encoder)
}
