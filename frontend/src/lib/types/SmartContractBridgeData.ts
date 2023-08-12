export interface TxInfoContainer {
    txHash: string,
    txChainID: number
}

export interface TxInfoFileJson {
    file_name: string | null,
    file_type: string | null,
    check_sum: string | null,
    txInfo: TxInfoContainer[] | null,
}