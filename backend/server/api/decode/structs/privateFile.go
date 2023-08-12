package decodeStructs

type PrivateFile struct {
	FileName string   `json:"file_name"`
	FileType string   `json:"file_type"`
	CheckSum string   `json:"check_sum"`
	TxInfo   []TxData `json:"tx_info"`
}

type TxData struct {
	ChainID int    `json:"chain_ID"`
	TxHash  string `json:"tx_hash"`
}

type PublicInfo struct {
	ReturnBuffer []byte `json:"return_buffer"`
	FileName     string `json:"file_name"`
}
