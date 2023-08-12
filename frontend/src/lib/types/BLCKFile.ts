
export interface BLCKFile {
    checksum: string
    files: BLCK[]
}

export interface BLCK {
    file_name: string;
    file_size: number;
    data: Uint8Array;
}
