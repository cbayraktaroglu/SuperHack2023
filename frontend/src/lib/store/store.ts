import type { Writable } from 'svelte/store';
import { writable } from 'svelte/store';
import type { BlockList } from '$lib/types/EncodeResponse';

export const Progress: Writable<number> = writable(0);
export const TheBlockList: Writable<BlockList[]> = writable();
