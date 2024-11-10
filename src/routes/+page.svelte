<script lang="ts">
	import { page } from '$app/stores';
	import FlashcardArea from '../components/FlashcardArea.svelte';
	import StatsArea from '../components/StatsArea/StatsArea.svelte';
	import { getStore } from '../store.svelte';
	import './styles.scss';

	const store = getStore();
</script>

{#if store.pageStatus === 'loading'}
	<div class="overlay">
		<div class="overlay-text">Parsing data...</div>
	</div>
{/if}
<main class="{store.pageStatus === 'loading' ? 'blurArea' : ''}">
	{#if ['ready', 'loading'].includes(store.pageStatus)}
		<FlashcardArea />
		<StatsArea />
		<div class="markdown-tutorial">
			{@html $page.data.htmlContent}
		</div>
	{:else if store.pageStatus === 'error'}
		<p>error: {store.errorMessage}</p>
	{/if}
</main>
