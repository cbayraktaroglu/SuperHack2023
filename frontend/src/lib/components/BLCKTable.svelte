<script lang="ts">
	import type { BLCKFile, BlockList } from '$lib/types/EncodeResponse';
	import { Progress, ProcessedInput, TxList } from '$lib/store/store';

	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import { Button } from 'flowbite-svelte';
	import { Icon } from 'flowbite-svelte-icons';

	import detectEthereumProvider from '@metamask/detect-provider';

	// Encode endpoint response
	let blckFile: BLCKFile;
	ProcessedInput.subscribe((value) => {
		blckFile = value;
	});

	let txHasList: string[] = [];

	//TODO duplicated code make it componenent and store wallet address in store as well
	async function sendBlockData(
		toAddres: string,
		valueToSend: string,
		fileNum: number
	): Promise<void> {
		const metaMaskEth = await detectEthereumProvider();
		if (!metaMaskEth) {
			console.log('MetaMask extension not found');
			return;
		}
		//get accounts
		const accounts = (await metaMaskEth.request({
			method: 'eth_accounts',
			params: []
		})) as string[];

		//user first address
		const transactionParams = {
			from: accounts[0],
			to: toAddres, // Replace with wallet a toddress
			value: 0,
			data: valueToSend
		};

		try {
			const hash = (await metaMaskEth.request({
				method: 'eth_sendTransaction',
				params: [transactionParams] as any
			})) as string;

			// Add the tx hash to the correct place
			txHasList[fileNum] = hash;

			// Update the global variable store
			TxList.set(txHasList);

			console.log('TxHash added to list!', hash);
		} catch (error) {
			console.log(error);
		}
	}

	// Function for getting individual .blck file data
	function processBLCK(blck: BlockList): any {
		return function (event: Event): void {
			console.log(blck.file_name);

			let blckNumber: number = Number(blck.file_name.charAt(0));
			let hex: string = arrayBufferToHex(blck.data);

			sendBlockData('0x0000000000000000000000000000000000000000', hex, blckNumber);
		};
	}

	function arrayBufferToHex(buffer: Uint8Array): string {
		let hex = '';
		for (let i = 0; i < buffer.length; i++) {
			hex += buffer[i].toString(16).padStart(2, '0');
		}
		return hex;
	}
</script>

<Table hoverable={true}>
	<TableHead>
		<TableHeadCell>Part</TableHeadCell>
		<TableHeadCell>Size</TableHeadCell>
		<TableHeadCell>
			<span class="sr-only">Upload</span>
		</TableHeadCell>
	</TableHead>
	<TableBody tableBodyClass="divide-y">
		{#each blckFile.files as item}
			<TableBodyRow>
				<TableBodyCell>{item.file_name}</TableBodyCell>
				<TableBodyCell>{item.file_size}</TableBodyCell>
				<TableBodyCell>
					<Button color="dark" class="!p-2" on:click={processBLCK(item)}
						><Icon name="arrow-right-outline" class="w-5 h-5" /></Button
					>
				</TableBodyCell>
			</TableBodyRow>
		{/each}
	</TableBody>
</Table>
