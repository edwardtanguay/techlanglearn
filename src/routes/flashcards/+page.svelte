<script lang="ts">
	import { getStore } from '../../store.svelte';
	import type { Flashcard } from '../../types';
	import './styles.scss';

	const store = getStore();

	let inputRefs: Record<number, HTMLInputElement | null> = {};

	const maxShow = 10;

	const handleToggleFlashcard = (flashcard: Flashcard, index: number) => {
		flashcard.status = 'answering';
		setTimeout(() => {
			inputRefs[index]?.focus();
		}, 10);
	};

	const handleSubmitAnswer = (flashcard: Flashcard) => {
		if (flashcard.suppliedAnswer === flashcard.back) {
			flashcard.status = 'correct';
		} else {
			flashcard.status = 'incorrect';
			setTimeout(() => {
				flashcard.suppliedAnswer = '';
				flashcard.status = 'showingFrontOnly';
			}, 5000);
		}
	};
</script>

{#snippet FlashcardBack(flashcard: Flashcard)}
	<p class="text-green-700">{flashcard.back}</p>
	{#if flashcard.extras !== ''}
		<p class="text-sm italic text-gray-400">{flashcard.extras}</p>
	{/if}
{/snippet}

<div class="area_flashcard">
	<h2 class="mb-6 text-xl">{maxShow} of {store.flashcards.length} flashcards.</h2>

	{#each store.flashcards.slice(0, maxShow) as flashcard, index}
		<div class="mb-6 w-fit">
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			{#if flashcard.status !== 'correct'}
				<div
					on:click={() => handleToggleFlashcard(flashcard, index)}
					class={`front ${flashcard.status === 'showingFrontOnly' ? 'cursor-pointer' : ''} rounded-t-md bg-slate-400 px-3 py-2 front-lang-${flashcard.language}`}
				>
					{flashcard.front}
				</div>
			{/if}

			{#if flashcard.status === 'answering'}
				<div class={`back rounded-b-md bg-slate-500 px-3 py-2`}>
					<input
						class="w-[30rem]"
						bind:this={inputRefs[index]}
						bind:value={flashcard.suppliedAnswer}
						on:keydown={(e) => {
							if (e.key === 'Enter') handleSubmitAnswer(flashcard);
						}}
					/>
				</div>
			{/if}
			{#if flashcard.status === 'incorrect'}
				<div class={`back rounded-b-md bg-slate-500 px-3 py-2`}>
					<p class="text-red-800">{flashcard.suppliedAnswer}</p>
					{@render FlashcardBack(flashcard)}
				</div>
			{/if}
			{#if flashcard.status === 'correct'}
				<div class={`backCorrect rounded-md border border-green-500 bg-green-100 px-3 py-2`}>
					{@render FlashcardBack(flashcard)}
				</div>
			{/if}
		</div>
	{/each}
</div>
