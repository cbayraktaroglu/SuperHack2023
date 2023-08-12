<script lang="ts">
	import type { TxInfoContainer } from '$lib/types/SmartContractBridgeData';
	import { Label } from 'flowbite-svelte';
	import { Progress, FileSHA256Checksum, TxList } from '$lib/store/store';

	// Subscribe to changes
	let fileChecksum: string;
	FileSHA256Checksum.subscribe((value) => {
		fileChecksum = value;
	});

	let txList: TxInfoContainer[];
	TxList.subscribe((value) => {
		txList = value;
	});
</script>

<Label>Now we need to connect this to the smart contract to send all the tx hashes to it</Label>
<Label>Here is how you can access file checksum: {fileChecksum}</Label>
<Label>And here is how you can access the tx list containing the txs: {txList}</Label>
