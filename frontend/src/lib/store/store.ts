import type { Writable } from 'svelte/store';
import { writable } from 'svelte/store';
import type { BLCKFile } from '$lib/types/BLCKFile';
import type { TxInfoContainer } from '$lib/types/SmartContractBridgeData';

// For storing the overall progress
export const Progress: Writable<number> = writable(0);

// For storing the file in .blcks
export const ProcessedInput: Writable<BLCKFile> = writable();

// For storing the txs that contain each .blck file
export const TxList: Writable<TxInfoContainer[]> = writable();

// For storing the selected file's checksum 
export const FileSHA256Checksum: Writable<string> = writable();