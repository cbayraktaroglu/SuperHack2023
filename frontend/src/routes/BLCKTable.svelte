<script lang="ts">
	import type { BlockList } from '$lib/types/EncodeResponse';
	import { Progress, TheBlockList } from '$lib/store/store';

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
	let blockList: BlockList[];
	TheBlockList.subscribe((value) => {
		blockList = value;
	});

	let txHasList: string[] = [];

	//TODO duplicated code make it componenent and store wallet address in store as well
	async function sendBlockData(toAddres: string, valueToSend: string): Promise<void> {
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
			value: valueToSend
		};

		try {
			const x = (await metaMaskEth.request({
				method: 'eth_sendTransaction',
				params: [transactionParams] as any
			})) as string;
			txHasList.push(x);
			console.log('TxHash added to list!', x);
		} catch (error) {
			console.log(error);
		}
	}

	// Function for getting individual .blck file data
	function processBLCK(blck: BlockList): any {
		return function (event: Event): void {
			console.log(blck.file_name);

			fetch('http://localhost:3000/data/' + blck.file_name, {
				method: 'GET'
			})
				.then((response) => {
					if (response.ok) {
						return response.json();
					} else {
						throw new Error('Error uploading file: ' + response.statusText);
					}
				})
				.then((serverResponse: any) => {
					console.log(serverResponse);
					console.log(serverResponse.Data);
					//Send each block data to some address
					//You can change address )
					sendBlockData('0x0', serverResponse.Data);
				})
				.catch((error) => {
					console.error('Error uploading file :', error);
				});
		};
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
		{#each blockList as item}
			<TableBodyRow>
				<TableBodyCell>{item.file_name}</TableBodyCell>
				<TableBodyCell>{item.file_size}</TableBodyCell>
				<TableBodyCell>
					<Button class="!p-2" on:click={processBLCK(item)}
						><Icon name="arrow-right-outline" class="w-5 h-5" /></Button
					>
				</TableBodyCell>
			</TableBodyRow>
		{/each}
	</TableBody>
</Table>
