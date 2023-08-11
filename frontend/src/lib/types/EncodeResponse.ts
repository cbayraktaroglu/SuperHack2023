
export interface BLCKFile {
    checksum: string
    files: BlockList[]
}

export interface BlockList {
    file_name: string;
    file_size: number;
    data: Uint8Array;
}
