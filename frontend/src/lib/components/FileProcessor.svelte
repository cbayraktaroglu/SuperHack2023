<script lang="ts">
	import type { BLCKFile, BLCK } from '$lib/types/BLCKFile';
	import {
		Progress,
		ProcessedInput,
		FileSHA256Checksum,
		FileName,
		FileExtension
	} from '$lib/store/store';
	import { NumberInput, Fileupload, Label, Helper } from 'flowbite-svelte';
	import HueButton from '$lib/components/HueButton.svelte';
	import CryptoJS from 'crypto-js';

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

	// Divides the file into .blcks and calculates the checksum of the original file
	const processTheFile = async (file: File) => {
		const reader = new FileReader();

		return new Promise<BLCKFile>((resolve, reject) => {
			reader.onload = () => {
				if (selectedFile) {
					// Init the return variable
					let retInfo: BLCKFile = {
						checksum: '',
						files: []
					};

					// Calculate the checksum
					const wordArray = CryptoJS.lib.WordArray.create(reader.result);
					const checksum = CryptoJS.SHA256(wordArray).toString(CryptoJS.enc.Hex);

					// Assign the checksum
					retInfo.checksum = checksum;

					// Start file processing
					var lastFileSize = selectedFile.size % value;
					var numberOfFiles = (selectedFile.size - lastFileSize) / value;

					console.log('Last file size: ' + lastFileSize);
					console.log('Number of files: ' + numberOfFiles);

					var blockList: BLCK[] = [];

					for (let i = 0; i < numberOfFiles + 1; i++) {
						let upperLimit = value;

						if (i == numberOfFiles) {
							upperLimit = lastFileSize;
						}

						let buffer = new Uint8Array(reader.result as ArrayBuffer, i * value, upperLimit);

						let item: BLCK = {
							file_name: i + '.blck',
							file_size: buffer.length,
							data: buffer
						};

						if (i == numberOfFiles) {
							item.file_size = lastFileSize;
						}

						if (!(i == numberOfFiles && lastFileSize == 0)) {
							blockList[i] = item;
						}
					}

					// Assign blocklist
					retInfo.files = blockList;

					resolve(retInfo);
				}

				reject();
			};

			reader.onerror = (error) => {
				reject(error);
			};

			reader.readAsArrayBuffer(file);
		});
	};

	// File handler
	function handleFile(event: Event): void {
		if (selectedFile) {
			// Sanity check
			if (selectedFile.size < value || value == 0) {
				console.log('Secified block size can not be bigger than the original file size!');
				return;
			}
			processTheFile(selectedFile)
				.then((result) => {
					console.log(`Checksum of ${selectedFile?.name}: ${result.checksum}`);
					console.log(result.files);

					// Update the stored variables
					ProcessedInput.set(result);
					FileSHA256Checksum.set(result.checksum);
					FileName.set(selectedFile.name.split('.', 2)[0]);
					FileExtension.set(selectedFile.name.split('.', 2)[1]);
					Progress.set(1);
				})
				.catch((error) => {
					console.error('Error calculating checksum:', error);
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
			<HueButton buttonText="Upload!" triggerFunction={handleFile} />
		</Label>
	{/each}
{/if}

<style>
</style>
