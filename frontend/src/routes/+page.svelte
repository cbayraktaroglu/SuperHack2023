<script lang="ts">
	import { Progress } from '$lib/store/store';
	import UploadWidget from './UploadWidget.svelte';
	import BlckTable from './BLCKTable.svelte';
	import logo from '$lib/images/Logo.webp';
	import logo_fallback from '$lib/images/Logo.png';

	// Progress indicator
	let progress: number;
	Progress.subscribe((value) => {
		progress = value;
	});
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

<section>
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
		flex: 0.6;
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