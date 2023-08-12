<script lang="ts">
	import type { TxInfoContainer, TxInfoFileJson } from '$lib/types/SmartContractBridgeData';
	import type { Factory } from '$lib/types/contracts';
	import { Label } from 'flowbite-svelte';
	import {
		Progress,
		FileSHA256Checksum,
		TxList,
		FileName,
		FileExtension,
		TxHashContainingTheFileCreation
	} from '$lib/store/store';
	import { onMount } from 'svelte';
	import { Factory__factory } from '$lib/types/contracts';
	import { Button } from 'flowbite-svelte';
	import { ethers } from 'ethers';

	import HueButton from './HueButton.svelte';
	import { downloadTxInfoAsJsonFile } from '$lib/components/DownloadTx';
	// Subscribe to changes
	let fileChecksum: string;
	FileSHA256Checksum.subscribe((value) => {
		fileChecksum = value;
	});

	let txList: TxInfoContainer[];
	TxList.subscribe((value) => {
		txList = value;
	});

	let nameOfTheFile: string;
	FileName.subscribe((value) => {
		nameOfTheFile = value;
	});

	let extensionOfTheFile: string;
	FileExtension.subscribe((value) => {
		extensionOfTheFile = '.' + value;
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
				.createFile(extensionOfTheFile, nameOfTheFile, fileChecksum, txHashes, txChains)
				.then((resp) => {
					TxHashContainingTheFileCreation.set(resp.hash);
					Progress.set(3);
				})
				.catch((err) => {
					console.log(err);
				});
		}
	}
	function createJsonAndDownload(): void {
		let json = {
			file_name: nameOfTheFile,
			file_type: extensionOfTheFile,
			check_sum: fileChecksum,
			txInfo: txList
		} as TxInfoFileJson;

		const fileName: string = nameOfTheFile + '-txhash.json';
		downloadTxInfoAsJsonFile(json, fileName);
	}
</script>

<div>
	<div class="grid grid-cols-2 gap-16 content-center py-16">
		<h4>To assign an ethereum address to your file press the GO!</h4>

		<HueButton buttonText="GO!" triggerFunction={executeTx} />

		<h4>You can dowload transaction information as Json!</h4>

		<HueButton buttonText="Download!" triggerFunction={createJsonAndDownload} />
	</div>
</div>
