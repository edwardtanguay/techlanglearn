<script lang="ts">
	import { page } from '$app/stores';
	import FlashcardArea from '../components/FlashcardArea.svelte';
	import StatsArea from '../components/StatsArea.svelte';
</script>

<FlashcardArea />
<StatsArea/>
<div class="markdown-tutorial">
	{@html $page.data.htmlContent}
</div>
