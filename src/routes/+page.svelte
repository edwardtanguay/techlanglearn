<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import FlashcardArea from '../components/FlashcardArea.svelte';
	import StatsArea from '../components/StatsArea/StatsArea.svelte';

	onMount(() => {
		console.log('111134', 'loaded');
	})
</script>

<FlashcardArea />
<StatsArea/>
<div class="markdown-tutorial">
	{@html $page.data.htmlContent}
</div>
