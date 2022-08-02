package domain_test

import (
	"encoder/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}
