<script lang="ts">
	import { getStore } from '../../store.svelte';
	import './styles.scss';

	const store = getStore();
</script>

<div class="area_flashcard">
	<h2 class="mb-6 text-xl">There are {store.flashcards.length} flashcards.</h2>

	{#each store.flashcards as flashcard}
		<div class="mb-6 w-fit">
			<p class={`bg-slate-400 px-3 py-2 rounded-t-md lang-${flashcard.language}`}>{flashcard.front}</p>
			<p class="bg-slate-500 text-yellow-300 px-3 py-2 rounded-b-md">{flashcard.back}</p>
		</div>
	{/each}
</div>
