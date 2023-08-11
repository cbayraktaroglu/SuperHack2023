import type { Writable } from 'svelte/store';
import { writable } from 'svelte/store';
import type { BLCKFile } from '$lib/types/EncodeResponse';

export const Progress: Writable<number> = writable(0);
export const ProcessedInput: Writable<BLCKFile> = writable();
export const TxList: Writable<string[]> = writable();
