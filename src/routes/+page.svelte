<script lang="ts">
	import { page } from '$app/stores';
	import FlashcardArea from '../components/FlashcardArea.svelte';
	import StatsArea from '../components/StatsArea/StatsArea.svelte';
	import { getStore } from '../store.svelte';

	const store = getStore();

</script>

{#if store.pageStatus === 'ready'}
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
