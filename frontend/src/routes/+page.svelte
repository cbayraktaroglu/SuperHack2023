<script lang="ts">
	import type { EncodeResponse, EncodeResponseData, BlockList } from '$lib/types/EncodeResponse';
	import welcome from '$lib/images/svelte-welcome.webp';
	import welcome_fallback from '$lib/images/svelte-welcome.png';
	import { NumberInput, Fileupload, Label, Helper, Button } from 'flowbite-svelte';
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import HueButton from './HueButton.svelte';

	// Progress indicator
	let progress: number = 0;

	// Parameters for local file upload
	let files: FileList;
	let selectedFile: File | null;

	// Block size that is being inputted by the user
	let value: number;

	// Encode endpoint response
	let blockList: BlockList[];

	$: if (files) {
		console.log(files);

		for (const file of files) {
			selectedFile = file;
			console.log(`${file.name}: ${file.size} bytes`);
		}
	}

	$: if (blockList) {
		console.log('Ho HEY!');
		console.log(blockList);
	}

	function handleFileUpload(event: Event): void {
		if (selectedFile) {
			if (selectedFile.size < value || value == 0) {
				console.log('Secified block size can not be bigger than the original file size!');
				return;
			}
			const formData = new FormData();
			formData.append('file', selectedFile);
			formData.append('blockSize', '' + value);

			fetch('http://localhost:3000/encode', {
				method: 'POST',
				body: formData
			})
				.then((response) => {
					if (response.ok) {
						return response.json();
					} else {
						throw new Error('Error uploading file: ' + response.statusText);
					}
				})
				.then((serverResponse: EncodeResponse) => {
					blockList = serverResponse.Data.FileList;
					console.log(serverResponse.Data);
					progress = 1;
				})
				.catch((error) => {
					console.error('Error uploading file :', error);
				});
		}
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
		<p class="p-3 mb-2 text-xl text-black-500 dark:text-gray-400">Let's get started</p>

		<Label class="space-y-2 mb-2">
			<Fileupload bind:files />
		</Label>

		{#if files}
			{#each Array.from(files) as file}
				<p class="p-3">{file.size} bytes</p>
				<Label class="p-3">
					<span>Single Block Size</span>
					<NumberInput bind:value />
					<Helper class="text-sm py-2">
						Determines the number of transactions needed for the file upload. As this value
						increases, required number of transactions will decrease. <strong
							>This value cannot exceed the original file's size!</strong
						>
					</Helper>
				</Label>
				<Label class="p-3">
					<HueButton buttonText="Upload!" triggerFunction={handleFileUpload} />
				</Label>
			{/each}
		{/if}
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
							<HueButton buttonText="->" triggerFunction={handleFileUpload} />
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
