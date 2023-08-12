export interface TxInfoContainer {
    tx_hash: string,
    chain_ID: number
}

export interface TxInfoFileJson {
    file_name: string | null,
    file_type: string | null,
    check_sum: string | null,
    tx_info: TxInfoContainer[] | null,
}