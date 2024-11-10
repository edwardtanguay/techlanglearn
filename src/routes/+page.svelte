<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import FlashcardArea from '../components/FlashcardArea.svelte';
	import StatsArea from '../components/StatsArea/StatsArea.svelte';

	type PageStatus = 'loading' | 'ready' | 'error';

	let pageStatus: PageStatus = $state("loading")
	let message = $state('');

	onMount(async () => {
		const response = await fetch('/api/pd');
		if (response.ok) {
			pageStatus = "ready"
		} else {
			const error = await response.json();
			pageStatus = "error"
			message = error.message;
		}
	});
</script>

{#if pageStatus === "ready"}
	<FlashcardArea />
	<StatsArea />
	<div class="markdown-tutorial">
		{@html $page.data.htmlContent}
	</div>
{:else if pageStatus === "loading"}
	<p>loading...</p>
{:else if pageStatus === "error"}
	<p>error: {message}</p>
{/if}
