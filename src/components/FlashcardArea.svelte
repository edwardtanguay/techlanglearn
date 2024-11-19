<script lang="ts">
	import type { Flashcard } from '../types';
	import * as dataModel from '../dataModel';
	import { getRandomItemsFromArray } from '../tools';
	import { getStore} from '../store.svelte';

	const store = getStore();

	const _flashcards = dataModel.getFlashcards();
	let flashcards = $state(getRandomItemsFromArray(_flashcards, 3));

	const allOpen = $derived(flashcards[0].isOpen && flashcards[1].isOpen && flashcards[2].isOpen);

	const handleFlashcardToggle = (flashcard: Flashcard) => {
		flashcards = flashcards.map((f) => (f === flashcard ? { ...f, isOpen: !f.isOpen } : f));
	};
</script>


<section class="w-fit rounded bg-slate-400 p-6 font-mono mb-3">
	<div class="flex flex-wrap gap-3">
		{#each flashcards as flashcard}
			<!-- svelte-ignore a11y_positive_tabindex -->
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<div
				role="button"
				tabindex="1"
				aria-pressed={flashcard.isOpen}
				class="w-fit cursor-pointer select-none rounded border border-slate-500 bg-slate-700 p-2 text-sm text-slate-300 hover:opacity-90"
				onclick={() => handleFlashcardToggle(flashcard)}
			>
				<p>{flashcard.front}</p>
				{#if flashcard.isOpen}
					<p class="text-yellow-300">{flashcard.back}</p>
				{:else}
					<p>{'_'.repeat(flashcard.back.length)}</p>
				{/if}
			</div>
		{/each}
	</div>
</section>