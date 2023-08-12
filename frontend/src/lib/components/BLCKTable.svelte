<script lang="ts">
	import type { BLCKFile, BLCK } from '$lib/types/BLCKFile';
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
	import type { TxInfoContainer } from '$lib/types/SmartContractBridgeData';

	// Parsed file
	let blckFile: BLCKFile;
	ProcessedInput.subscribe((value) => {
		blckFile = value;
	});

	let txHasList: TxInfoContainer[] = [];

	// To keep track of the processed .blck files
	let isProcessed: { [key: string]: boolean } = {};
	let numberOfProcessed: number = 0;

	// Listener for updating the progress
	$: if (numberOfProcessed) {
		console.log(numberOfProcessed);

		if (numberOfProcessed == blckFile.files.length) {
			Progress.set(2);
		}
	}

	//TODO duplicated code make it componenent and store wallet address in store as well
	async function sendBlockData(
		toAddres: string,
		valueToSend: string,
		fileNum: number
	): Promise<void> {
		const blockchainProvider = await detectEthereumProvider();
		if (!blockchainProvider) {
			console.log('MetaMask extension not found');
			return;
		}
		//get accounts
		const accounts = (await blockchainProvider.request({
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
			const hash = (await blockchainProvider.request({
				method: 'eth_sendTransaction',
				params: [transactionParams] as any
			})) as string;

			const chainID = await blockchainProvider.request({
				method: 'eth_chainId'
			});

			// Add the tx hash to the correct place
			txHasList[fileNum] = { txHash: hash, txChainID: parseInt(chainID, 16) };

			// Update the global variable store
			TxList.set(txHasList);

			console.log(fileNum + '.blck has been processed');

			isProcessed[fileNum + '.blck'] = true;
			numberOfProcessed = numberOfProcessed + 1;

			console.log('TxHash added to list!', hash);
		} catch (error) {
			console.log(error);
		}
	}

	// Function for getting individual .blck file data
	function processBLCK(blck: BLCK): any {
		return function (event: Event): void {
			console.log(blck.file_name);

			let blckNumber: number = Number(blck.file_name.split('.', 1)[0]);
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
					<Button
						disabled={isProcessed[item.file_name]}
						color="dark"
						class="!p-2"
						on:click={processBLCK(item)}><Icon name="arrow-right-outline" class="w-5 h-5" /></Button
					>
				</TableBodyCell>
			</TableBodyRow>
		{/each}
	</TableBody>
</Table>
