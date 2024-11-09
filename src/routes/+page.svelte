<script lang="ts">
	import { page } from '$app/stores';
	import FlashcardArea from '../components/FlashcardArea.svelte';
</script>

<FlashcardArea />
<div class="markdown-tutorial">
	{@html $page.data.htmlContent}
</div>
