<script lang="ts">
	import type { TxInfoContainer } from '$lib/types/SmartContractBridgeData';
	import type { Factory } from '$lib/types/contracts';
	import { Label } from 'flowbite-svelte';
	import { Progress, FileSHA256Checksum, TxList } from '$lib/store/store';
	import { onMount } from 'svelte';
	import { Factory__factory } from '$lib/types/contracts';
	import { Button } from 'flowbite-svelte';
	import { ethers } from 'ethers';

	// Subscribe to changes
	let fileChecksum: string;
	FileSHA256Checksum.subscribe((value) => {
		fileChecksum = value;
	});

	let txList: TxInfoContainer[];
	TxList.subscribe((value) => {
		txList = value;
	});

	// Blockchain variables
	let provider: ethers.BrowserProvider;
	let fileFactory: Factory;
	let signer: ethers.JsonRpcSigner;

	// run the function onMount to set everything
	onMount(async () => {
		await initializeTheContract();
	});

	async function initializeTheContract(): Promise<void> {
		if (typeof window.ethereum !== 'undefined') {
			provider = new ethers.BrowserProvider(window.ethereum);
			try {
				// Request account access
				await window.ethereum.request({ method: 'eth_requestAccounts' });

				signer = await provider.getSigner();

				// Get the contract
				fileFactory = Factory__factory.connect(
					'0xcf58de8eccc3a54cc1c58f5932dadba5e1483fdb',
					signer
				);
			} catch (error) {
				console.error('User rejected the request:', error);
			}
		} else {
			console.error('Wallet is not installed!');
		}
	}

	async function executeTx(event: Event) {
		if (fileFactory) {
			//Prepare the tx hash list and chainId list
			let txHashes: string[] = [];
			let txChains: string[] = [];

			for (let i = 0; i < txList.length; i++) {
				txHashes.push(txList[i].txHash);
				txChains.push('' + txList[i].txChainID);
			}

			fileFactory
				.createFile('.gif', 'testus', fileChecksum, txHashes, txChains)
				.then((resp) => {
					console.log(resp.hash);
				})
				.catch((err) => {
					console.log(err);
				});
		}
	}
</script>

<Label>Now we need to connect this to the smart contract to send all the tx hashes to it</Label>
<Label>Here is how you can access file checksum: {fileChecksum}</Label>
<Label>And here is how you can access the tx list containing the txs: {txList}</Label>
<Button color="dark" class="!p-2" on:click={executeTx}>GO!</Button>
