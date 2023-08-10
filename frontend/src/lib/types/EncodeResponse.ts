
export interface EncodeResponse {
    Data: EncodeResponseData;
    Status: boolean;
}

export interface EncodeResponseData {
    Checksum: string;
    FileList: BlockList[];
}

export interface BlockList {
    file_name: string;
    file_size: number;
}
