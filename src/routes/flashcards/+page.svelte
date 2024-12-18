<script lang="ts">
	import { getStore } from '../../store.svelte';
	import './styles.scss';

	const store = getStore();
</script>

<div class="area_flashcard">
	<h2 class="mb-6 text-xl">There are {store.flashcards.length} flashcards.</h2>

	{#each store.flashcards as flashcard}
		<div class="mb-6 w-fit">
			<p class={`front rounded-t-md bg-slate-400 px-3 py-2 front-lang-${flashcard.language}`}>
				{flashcard.front}
			</p>
			<p class={`back rounded-b-md bg-slate-500 px-3 py-2 back-lang-${flashcard.language}`}>
				{flashcard.back}
			</p>
		</div>
	{/each}
</div>
