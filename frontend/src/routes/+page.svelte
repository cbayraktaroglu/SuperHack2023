<script lang="ts">
	import { Progress } from '$lib/store/store';
	import UploadWidget from './UploadWidget.svelte';
	import BlckTable from './BLCKTable.svelte';
	import logo from '$lib/images/Logo.webp';
	import logo_fallback from '$lib/images/Logo.png';
	import HueButton from './HueButton.svelte';
	import detectEthereumProvider from '@metamask/detect-provider';
	import { onMount } from 'svelte';
	import { Label } from 'flowbite-svelte';
	// Progress indicator
	let progress: number;
	Progress.subscribe((value) => {
		progress = value;
	});

	onMount(async () => {
		await getConnectedMetaMaskWallet();
	});

	// MetaMask wallet connection
	let walletAddress: string;
	async function connectMetaMaskWallet(): Promise<void> {
		const metaMaskEth = await detectEthereumProvider();
		if (!metaMaskEth) {
			console.log('MetaMask extension not found');
			return;
		}

		try {
			const accounts = (await metaMaskEth.request({
				method: 'eth_requestAccounts',
				params: []
			})) as string[];
			walletAddress = accounts[0];
			console.log(accounts);
			// Get chain Id and console log it
			const chainId = (await metaMaskEth.request({
				method: 'eth_chainId',
				params: []
			})) as string;
			console.log('Chain Id', chainId);
		} catch (error) {
			console.log(error);
		}
	}

	// Function for getting connected MetaMask wallet and will be called on every page load
	async function getConnectedMetaMaskWallet(): Promise<void> {
		const metaMaskEth = await detectEthereumProvider();
		if (!metaMaskEth) {
			console.log('MetaMask extension not found');
			return;
		}
		try {
			const accounts = (await metaMaskEth.request({
				method: 'eth_accounts',
				params: []
			})) as string[];
			if (accounts.length == 0) {
				console.log('No accounts found. Please Connect Again!');
				return;
			} else {
				walletAddress = accounts[0];
				console.log(accounts);
			}
		} catch (error) {
			console.log(error);
		}
	}

	// Function for making transaction on MetaMask wallet
	async function makeTransaction(): Promise<void> {
		const metaMaskEth = await detectEthereumProvider();
		if (!metaMaskEth) {
			console.log('MetaMask extension not found');
			return;
		}

		const transactionParams: { from: string; to: string; value: string }[] = [
			{
				from: walletAddress,
				to: '0x00',
				value: '0'
			}
		];

		try {
			await metaMaskEth.request({
				method: 'eth_sendTransaction',
				params: [
					{
						from: walletAddress,
						to: '0x00',
						value: '0'
					}
				] as { from: string; to: string; value: string }[]
			});
			// .then((txHash: string) => console.log(txHash))
			// .catch((error: any) => console.error(error));
		} catch (error) {
			console.log(error);
		}
	}

	// Function for getting connected MetaMask wallet and will be called on every page load
	async function getConnectedMetaMaskWallet(): Promise<void> {
		const metaMaskEth = await detectEthereumProvider();
		if (!metaMaskEth) {
			console.log('MetaMask extension not found');
			return;
		}
		try {
			const accounts = (await metaMaskEth.request({
				method: 'eth_accounts',
				params: []
			})) as string[];
			if (accounts.length == 0) {
				console.log('No accounts found. Please Connect Again!');
				return;
			} else {
				walletAddress = accounts[0];
				console.log(accounts);
			}
		} catch (error) {
			console.log(error);
		}
	}

	// Function for making transaction on MetaMask wallet
	async function makeTransaction(): Promise<void> {
		const metaMaskEth = await detectEthereumProvider();
		if (!metaMaskEth) {
			console.log('MetaMask extension not found');
			return;
		}

		const transactionParams: { from: string; to: string; value: string }[] = [
			{
				from: walletAddress,
				to: '0x00',
				value: '0'
			}
		];

		try {
			await metaMaskEth.request({
				method: 'eth_sendTransaction',
				params: [
					{
						from: walletAddress,
						to: '0x00',
						value: '0'
					}
				] as { from: string; to: string; value: string }[]
			});
			// .then((txHash: string) => console.log(txHash))
			// .catch((error: any) => console.error(error));
		} catch (error) {
			console.log(error);
		}
	}
</script>

<svelte:head>
	<title>KeepIt</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<div class="welcome">
	<picture>
		<source srcset={logo} type="image/webp" />
		<img src={logo_fallback} alt="Welcome" />
	</picture>
</div>

	<Label class="p-3">
		<HueButton
			buttonText={walletAddress && walletAddress.length > 0
				? `Connected: ${walletAddress.substring(0, 6)}...${walletAddress.substring(38)}`
				: 'Connect Your Wallet!'}
			triggerFunction={connectMetaMaskWallet}
		/>
	</Label>

	<Label class="p-3">
		<HueButton
			buttonText={walletAddress && walletAddress.length > 0
				? `Connected: ${walletAddress.substring(0, 6)}...${walletAddress.substring(38)}`
				: 'Connect Your Wallet!'}
			triggerFunction={connectMetaMaskWallet}
		/>
	</Label>

	{#if progress == 0}
		<UploadWidget />
	{/if}

	{#if progress == 1}
		<BlckTable />
	{/if}
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.1;
	}

	.welcome {
		margin-top: 1vh; /* Adjust the value as needed to control vertical positioning */
		position: relative;
	}

	.welcome picture {
		display: block;
		text-align: center;
	}

	.welcome img {
		width: 50%;
		display: inline-block;
	}
</style>
