<script lang="ts">
	import { getStore } from '../store.svelte';

	const store = getStore();
	const randomFlashcards = store.randomFlashcards;
</script>

<section class="mb-3 w-fit rounded bg-slate-400 p-6 font-mono">
	<div class="flex flex-wrap gap-3">
		{#each randomFlashcards as flashcard}
			<!-- svelte-ignore a11y_positive_tabindex -->
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<div
				role="button"
				tabindex="1"
				aria-pressed={flashcard.isOpen}
				class="w-fit cursor-pointer select-none rounded border border-slate-500 bg-slate-700 p-2 text-sm text-slate-300 hover:opacity-90"
				onclick={() => store.handleRandomFlashcardToggle(flashcard)}
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
