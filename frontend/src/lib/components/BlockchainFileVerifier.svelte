<script lang="ts">
	import { Label } from 'flowbite-svelte';
	import { Progress, TxHashContainingTheFileCreation } from '$lib/store/store';

	// Subscribe to the values
	let targetTX: string;
	TxHashContainingTheFileCreation.subscribe((value) => {
		targetTX = value;
	});
</script>

<Label>
	Now you need to get the emmited event from {targetTX}, extract the newly created File contract
	address. Then call the verify with users input
</Label>
