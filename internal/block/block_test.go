package block

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_calculateHash(t *testing.T) {
	tests := []struct {
		testBlock    Block
		expectedHash string
	}{
		{
			testBlock: Block{
				Index:     1,
				Timestamp: "2023-04-09 16:47:08.213643 +0300 MSK m=+0.000662831",
				BPM:       120,
				Hash:      "81270bf9b03e9c245eeef66601426bcec676b513d4a635b84d664617d70d4712",
				PrevHash:  "81270bf9b03e9c245eeef66601426bcec676b513d4a635b84d664617d70d4710",
			},
			expectedHash: "e068b1f15599c57c7612909b1edce6c95bb0b11c367cc2f6491ecb80002a8f4c",
		},
	}

	for _, tt := range tests {
		actual := calculateHash(tt.testBlock)
		assert.Equal(t, tt.expectedHash, actual)
	}
}
