<script lang="ts">
	import { page } from '$app/stores';
	import { getRandomItemsFromArray } from '../tools';
	import type { Flashcard } from '../types';
	import * as dataModel from '../dataModel';

	const _flashcards = dataModel.getFlashcards();
	$: flashcards = getRandomItemsFromArray(_flashcards, 3);

	const handleFlashcardToggle = (flashcard: Flashcard) => {
		flashcards = flashcards.map((f) => (f === flashcard ? { ...f, isOpen: !f.isOpen } : f));
	};
</script>

<main class="p-6">
	<section class="w-fit rounded bg-slate-400 p-6 font-mono">
		<h2 class="mb-3 text-xl">{flashcards.length} of {_flashcards.length} Flashcards</h2>
		<div class="flex flex-wrap gap-3">
			{#each flashcards as flashcard}
				<!-- svelte-ignore a11y_positive_tabindex -->
				<!-- svelte-ignore a11y_click_events_have_key_events -->
				<div
					role="button"
					tabindex="1"
					aria-pressed={flashcard.isOpen}
					class="w-fit cursor-pointer select-none rounded border border-slate-500 bg-slate-700 p-2 text-sm text-slate-300 hover:opacity-90"
					on:click={() => handleFlashcardToggle(flashcard)}
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
	<header>
		<h1 class="main">
			Tech and Language Learning <span class="text-xs"
				>by <a href="https://tanguay.info" target="_blank" class="underline">Edward Tanguay</a
				></span
			>
		</h1>
		<p>
			On this site I record my learning of tech topics while watching video tutorials in numerous
			languages.
		</p>
		<p>
			This site was made with <b>SvelteKit 5</b>, see
			<a target="_blank" class="underline" href="https://github.com/edwardtanguay/techlanglearn"
				>repo</a
			>
			and
			<a target="_blank" class="underline" href="https://techlanglearn.vercel.app">live site</a>.
		</p>
		<p>
			Also see my <a target="_blank" class="underline" href="https://tanguay-eu.vercel.app/howtos"
				>other sites</a
			>.
		</p>
	</header>

	<div class="markdown-tutorial">
		{@html $page.data.htmlContent}
	</div>
</main>
