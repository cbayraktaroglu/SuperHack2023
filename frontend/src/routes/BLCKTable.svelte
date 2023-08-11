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
