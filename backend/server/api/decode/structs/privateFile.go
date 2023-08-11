package decodeStructs

type PrivateFile struct {
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	CheckSum string `json:"check_sum"`
	TxInfo   []struct {
		ChainID int    `json:"chain_ID"`
		TxHash  string `json:"tx_hash"`
	} `json:"tx_info"`
}
