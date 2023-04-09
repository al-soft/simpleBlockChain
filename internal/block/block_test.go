package block

import (
	"fmt"
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
		actual := CalculateHash(tt.testBlock)
		assert.Equal(t, tt.expectedHash, actual)
	}
}

func Test_generateBlock(t *testing.T) {
	tests := []struct {
		testOldBlock  Block
		testBPM       int
		expected      Block
		expectedError error
	}{
		{
			testOldBlock: Block{
				Index:     1,
				Timestamp: "2023-04-09 16:47:08.213643 +0300 MSK m=+0.000662831",
				BPM:       120,
				Hash:      "81270bf9b03e9c245eeef66601426bcec676b513d4a635b84d664617d70d4712",
				PrevHash:  "81270bf9b03e9c245eeef66601426bcec676b513d4a635b84d664617d70d4710",
			},
			testBPM: 120,
			expected: Block{
				Index:    2,
				PrevHash: "81270bf9b03e9c245eeef66601426bcec676b513d4a635b84d664617d70d4712",
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		actual, err := GenerateBlock(tt.testOldBlock, tt.testBPM)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%+v\n", actual)

		assert.Equal(t, tt.expectedError, err)
		assert.Equal(t, tt.expected.Index, actual.Index)
		assert.Equal(t, tt.expected.PrevHash, tt.testOldBlock.Hash)
	}
}

func Test_isBlockValid(t *testing.T) {
	tests := []struct {
		oldBlock Block
		newBlock Block
		expected bool
	}{
		{
			oldBlock: Block{
				Index: 1,
				Hash:  "81270bf9b03e9c245eeef66601426bcec676b513d4a635b84d664617d70d4712",
			},
			newBlock: Block{
				Index:    2,
				PrevHash: "81270bf9b03e9c245eeef66601426bcec676b513d4a635b84d664617d70d4712",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		tt.newBlock.Hash = CalculateHash(tt.newBlock)

		actual := IsBlockValid(tt.newBlock, tt.oldBlock)
		assert.Equal(t, tt.expected, actual)
	}
}

func Test_replaceChain(t *testing.T) {
	tests := []struct {
		testNewBlocks  []Block
		expectedBlocks []Block
	}{
		{
			testNewBlocks: []Block{
				{},
				{},
			},
			expectedBlocks: []Block{
				{},
				{},
			},
		},
	}

	for _, tt := range tests {
		ReplaceChain(tt.testNewBlocks)
		actual := Blockchain
		assert.Equal(t, tt.expectedBlocks, actual)
	}
}
