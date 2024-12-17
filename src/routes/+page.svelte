<script lang="ts">
	import { page } from '$app/stores';
	import StatsArea from '../components/StatsArea/StatsArea.svelte';
	import { getStore } from '../store.svelte';
	import './styles.scss';

	const store = getStore();
</script>

<main id="vocab">
	{#if ['ready', 'loading'].includes(store.pageStatus)}
		<StatsArea />
		<div class="markdown-tutorial">
			{@html $page.data.htmlContent}
		</div>
	{:else if store.pageStatus === 'error'}
		<p>error: {store.errorMessage}</p>
	{/if}
</main>
