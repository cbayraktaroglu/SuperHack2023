<script lang="ts">
	import type { EncodeResponse, EncodeResponseData, BlockList } from '$lib/types/EncodeResponse';
	import { Progress, TheBlockList } from '$lib/store/store';
	import { NumberInput, Fileupload, Label, Helper } from 'flowbite-svelte';
	import HueButton from './HueButton.svelte';

	// Parameters for local file upload
	let files: FileList;
	let selectedFile: File | null;

	// Block size that is being inputted by the user
	let value: number;

	// Listener for updating the data
	$: if (files) {
		console.log(files);

		for (const file of files) {
			selectedFile = file;
			console.log(`${file.name}: ${file.size} bytes`);
		}
	}

	// Function for handling first file upload
	function handleFileUpload(event: Event): void {
		console.log('Ho hey!');
		if (selectedFile) {
			console.log('Ho hey!!');
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
					console.log(serverResponse.Data);

					// Update the stored variables
					TheBlockList.set(serverResponse.Data.FileList);
					Progress.set(1);
				})
				.catch((error) => {
					console.error('Error uploading file :', error);
				});
		}
	}
</script>

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
				Determines the number of transactions needed for the file upload. As this value increases,
				required number of transactions will decrease. <strong
					>This value cannot exceed the original file's size!</strong
				>
			</Helper>
		</Label>
		<Label class="p-3">
			<HueButton buttonText="Upload!" triggerFunction={handleFileUpload} />
		</Label>
	{/each}
{/if}

<style>
</style>
