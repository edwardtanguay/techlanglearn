<script lang="ts">
	import { page } from '$app/stores';
	import FlashcardArea from '../components/FlashcardArea.svelte';
	import StatsArea from '../components/StatsArea/StatsArea.svelte';
	import { getStore } from '../store.svelte';

	const store = getStore();

	const reloadData = async () => {
		store.setPageStatus('loading');
		const response = await fetch('/api/pd');
		if (response.ok) {
			store.setPageStatus('ready');
		} else {
			const error = await response.json();
			store.setPageStatus('error');
			store.setErrorMessage(error.message);
		}
	};
</script>

{#if store.pageStatus === 'ready'}
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
{:else if store.pageStatus === 'loading'}
	<p>loading...</p>
{:else if store.pageStatus === 'error'}
	<p>error: {store.errorMessage}</p>
{/if}
