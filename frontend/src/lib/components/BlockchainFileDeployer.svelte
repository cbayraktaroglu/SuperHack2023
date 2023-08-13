<script lang="ts">
	import type { TxInfoContainer, TxInfoFileJson } from '$lib/types/SmartContractBridgeData';
	import type { KeepItFileFactory } from '$lib/types/contracts';
	import { KeepItFileFactory__factory } from '$lib/types/contracts';
	import { downloadTxInfoAsJsonFile } from '$lib/components/DownloadTx';

	import {
		Progress,
		FileSHA256Checksum,
		TxList,
		FileName,
		FileExtension,
		TxHashContainingTheFileCreation
	} from '$lib/store/store';

	import { onMount } from 'svelte';
	import { ethers } from 'ethers';

	import HueButton from './HueButton.svelte';
	import { Alert } from 'flowbite-svelte';
	import { Icon } from 'flowbite-svelte-icons';

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
	let fileFactory: KeepItFileFactory;
	let signer: ethers.JsonRpcSigner;
	let factoryAddresses: SmartContractAddressMap = {
		// Optimism Goerli
		'420': '0xcf91A26C978c33fCe244412cBcb602C63F749A8b',
		// Base Goerli
		'84531': '0xfbBc9950cB64912EDd88bCf47f6D85957C2aBcd0',
		// Zora Goerli
		'999': '0xfbBc9950cB64912EDd88bCf47f6D85957C2aBcd0',
		// Polygon Mumbai
		'80001': '0xce2dA00922faf10dd5bE5229666691eB28FcB75D',
		// Mode Sepolia
		'919': '0xA4575B1d61AA4fE963373e8FD535427205B91135'
	};

	// Variable for displaying the alert
	let isChainSupported: boolean = true;

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

				// Check if the current chain is supported
				let chainId = await provider.getNetwork();
				let currentChainAddress = factoryAddresses[chainId.chainId.toString()];

				if (currentChainAddress === undefined) {
					isChainSupported = false;
					throw new Error('unsupported chain');
				}

				// Get the contract
				fileFactory = KeepItFileFactory__factory.connect(currentChainAddress, signer);
			} catch (error) {
				console.error('User rejected the request:', error);
			}
		} else {
			console.error('Wallet is not installed!');
		}
	}

	async function executeTx(event: Event) {
		try {
			let chainId = await provider.getNetwork();
			let currentChainAddress = factoryAddresses[chainId.chainId.toString()];
			if (currentChainAddress === undefined) {
				throw new Error('unsupported chain');
			}

			// Re init the contract
			signer = await provider.getSigner();
			fileFactory = KeepItFileFactory__factory.connect(currentChainAddress, signer);
			isChainSupported = true;
			if (fileFactory) {
				//Prepare the tx hash list and chainId list
				let txHashes: string[] = [];
				let txChains: string[] = [];

				for (let i = 0; i < txList.length; i++) {
					txHashes.push(txList[i].tx_hash);
					txChains.push('' + txList[i].chain_ID);
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
		} catch (err) {
			isChainSupported = false;
			const match = err.message.match(/network changed: \d+ => (\d+)/);
			const extractedValue = match ? match[1] : null;

			console.log(extractedValue);
			let currentChainAddress = factoryAddresses[extractedValue];
			if (currentChainAddress === undefined) {
				throw new Error('unsupported chain');
			}

			// Re init the contract
			provider = new ethers.BrowserProvider(window.ethereum);
			signer = await provider.getSigner();
			fileFactory = KeepItFileFactory__factory.connect(currentChainAddress, signer);
			isChainSupported = true;
		}
	}
	function createJsonAndDownload(): void {
		let json = {
			file_name: nameOfTheFile,
			file_type: extensionOfTheFile,
			check_sum: fileChecksum,
			tx_info: txList
		} as TxInfoFileJson;

		const fileName: string = nameOfTheFile + '-txhash.json';
		downloadTxInfoAsJsonFile(json, fileName);
	}
</script>

<div>
	{#if !isChainSupported}
		<Alert color="yellow">
			<Icon name="info-circle-solid" slot="icon" class="w-4 h-4" />
			<span class="font-medium">Warning!</span>
			Address assigning on blockchain is currently only supported in: Optimism Goerli, Base Goerli, Zora
			Goerli, Polygon Mumbai, Mode Sepolia
		</Alert>
	{/if}
	<div class="grid grid-cols-2 gap-16 content-center py-16">
		<h4>To assign an ethereum address to your file press the keepIt public button!</h4>

		<HueButton buttonText="keepIt public" triggerFunction={executeTx} />

		<h4>You can dowload transaction information as Json!</h4>

		<HueButton buttonText="keepIt private" triggerFunction={createJsonAndDownload} />
	</div>
</div>
