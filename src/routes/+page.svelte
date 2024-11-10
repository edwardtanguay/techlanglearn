<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import FlashcardArea from '../components/FlashcardArea.svelte';
	import StatsArea from '../components/StatsArea/StatsArea.svelte';

	let showSite = $state(false);

	onMount(async () => {
		const response = await fetch('/api/pd');
		const data = await response.json();
		showSite = true;
	});
</script>

{#if showSite}
	<FlashcardArea />
	<StatsArea />
	<div class="markdown-tutorial">
		{@html $page.data.htmlContent}
	</div>
{:else}
	<p>loading...</p>
{/if}
