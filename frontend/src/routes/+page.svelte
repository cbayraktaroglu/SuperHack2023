<script lang="ts">
	import type { BlockList } from '$lib/types/EncodeResponse';
	import { Progress, TheBlockList } from '$lib/store/store';
	import UploadWidget from './UploadWidget.svelte';
	import welcome from '$lib/images/svelte-welcome.webp';
	import welcome_fallback from '$lib/images/svelte-welcome.png';
	import { Button } from 'flowbite-svelte';
	import { Icon } from 'flowbite-svelte-icons';
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';

	// Progress indicator
	let progress: number;
	Progress.subscribe((value) => {
		progress = value;
	});

	// Encode endpoint response
	let blockList: BlockList[];
	TheBlockList.subscribe((value) => {
		blockList = value;
	});

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
				})
				.catch((error) => {
					console.error('Error uploading file :', error);
				});
		};
	}
</script>

<svelte:head>
	<title>KeepIt</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<h1>
		<span class="welcome">
			<picture>
				<source srcset={welcome} type="image/webp" />
				<img src={welcome_fallback} alt="Welcome" />
			</picture>
		</span>
	</h1>

	{#if progress == 0}
		<UploadWidget />
	{/if}

	{#if progress == 1}
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
	{/if}
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}

	h1 {
		width: 100%;
	}

	.welcome {
		display: block;
		position: relative;
		width: 100%;
		height: 0;
		padding: 0 0 calc(100% * 495 / 2048) 0;
	}

	.welcome img {
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
		display: block;
	}
</style>
