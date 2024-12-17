<script lang="ts">
	import { getStore } from '../../store.svelte';

	const store = getStore();
</script>

<header>
	<p>There are {store.flashcards.length} flashcards.</p>
</header>
