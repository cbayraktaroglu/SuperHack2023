<script lang="ts">
	import { Label, NumberInput, Input } from 'flowbite-svelte';
	import { Progress, TxHashContainingTheFileCreation } from '$lib/store/store';
	import { KeepItFileFactory__factory } from '$lib/types/contracts';
	import type { KeepItFileFactory } from '$lib/types/contracts';

	import { KeepItFile__factory } from '$lib/types/contracts';
	import type { KeepItFile } from '$lib/types/contracts';

	import { onMount } from 'svelte';
	import { ethers } from 'ethers';

	import WorldIdWidget from './WorldIdWidget';
	// Subscribe to the values
	let targetTX: string;
	TxHashContainingTheFileCreation.subscribe((value) => {
		targetTX = value;
	});

	// Blockchain variables
	let provider: ethers.BrowserProvider;
	let fileFactory: KeepItFileFactory;
	let signer: ethers.JsonRpcSigner;

	let keepItFile: KeepItFile;

	let enteredUid: string;
	// // run the function onMount to set everything
	// onMount(async () => {
	// 	await initializeTheContract();
	// });

	async function verifyFileWithOrg(): Promise<void> {
		if (typeof window.ethereum !== 'undefined') {
			provider = new ethers.BrowserProvider(window.ethereum);
			try {
				// Request account access
				await window.ethereum.request({ method: 'eth_requestAccounts' });

				const topicAddres: string =
					'0x31a209c0c6a6a197b980906ee1598da098dcdcb051bf49abbc4a935702c0a603';

				const transactionReceipt = await provider.getTransactionReceipt(targetTX);

				if (!transactionReceipt) {
					console.log('transactionReceipt not found');
					return;
				}
				console.log('transactionReceipt', transactionReceipt);

				const found: any = transactionReceipt['logs'].find((log) => {
					return log['topics'].find((topic) => {
						return topic === topicAddres;
					});
				});

				console.log('found', found);
				const foundContractAdress: string = found['topics'][0];
				console.log('foundContractAdress', foundContractAdress);

				// Get the contract
				keepItFile = KeepItFile__factory.connect(foundContractAdress, signer);

				// // get event id from the transaction and call the verify function

				// // Call the verify function with enteredUid 
				await keepItFile.orgVerify(enteredUid);
			} catch (error) {
				console.error('User rejected the request:', error);
			}
		} else {
			console.error('Wallet is not installed!');
		}
	}
</script>

<div>
	<h1>Verify File</h1>

	<div class="p-12 grid grid-cols-2 gap-x-16 content-center">
		<div class="col-span-1">
			<react:WorldIdWidget />
		</div>
		<div class="col-span-1">
			<button class="verify-org" on:click={verifyFileWithOrg}> Verify with Org !</button>
		</div>
		<div>
			<Input bind:enteredUid />
		</div>
	</div>

	<Label>
		Now you need to get the emmited event from {targetTX}, extract the newly created File contract
		address. Then call the verify with users input
	</Label>
</div>

<style>
	.verify-org {
		padding: 10px 20px;
		border: none;
		cursor: pointer;
		background-color: #000;
		transition: background-color 0.3s ease-in-out;
		color: white;
		border-radius: 15px;
	}
</style>
