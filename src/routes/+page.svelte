<script lang="ts">
	import { page } from '$app/stores';
	import FlashcardArea from '../components/FlashcardArea.svelte';
	import StatsArea from '../components/StatsArea/StatsArea.svelte';
	import { getStore } from '../store';

	const store = getStore();

	type PageStatus = 'loading' | 'ready' | 'error';

	let pageStatus: PageStatus = $state('ready');
	let message = $state('');

	const reloadData = async () => {
		pageStatus = 'loading';
		const response = await fetch('/api/pd');
		if (response.ok) {
			console.log(111144, 'ok');
			pageStatus = 'ready';
		} else {
			console.log(111144, 'NOT OK');
			const error = await response.json();
			pageStatus = 'error';
			message = error.message;
		}
	};
</script>

<p class="text-red-600">{store.siteLocation}</p>
{#if pageStatus === 'ready'}
	{#if store.siteLocation === 'dev'}
		<button class="mb-3 rounded border border-slate-600 bg-slate-400 px-1" onclick={reloadData}
			>Parse data</button
		>
	{/if}
	<FlashcardArea />
	<StatsArea />
	<div class="markdown-tutorial">
		{@html $page.data.htmlContent}
	</div>
{:else if pageStatus === 'loading'}
	<p>loading...</p>
{:else if pageStatus === 'error'}
	<p>error: {message}</p>
{/if}
