<script lang="ts">
	import { getStore } from '../../store.svelte';
	import type { Flashcard } from '../../types';
	import './styles.scss';
	import IconCheck from '../../components/IconCheck.svelte';

	const store = getStore();

	let inputRefs: Record<number, HTMLInputElement | null> = {};

	const maxShow = 10;

	const handleToggleFlashcard = (flashcard: Flashcard, index: number) => {
		flashcard.status = 'answering';
		setTimeout(() => {
			inputRefs[index]?.focus();
		}, 10);
	};

	const handleSubmitAnswer = (flashcard: Flashcard, index: number) => {
		if (flashcard.suppliedAnswer === flashcard.back) {
			flashcard.status = 'correct';
		} else {
			flashcard.status = 'incorrect';
		}
	};

	const handleCloseFlashcard = (flashcard: Flashcard, index: number) => {
		flashcard.suppliedAnswer = '';
		flashcard.status = 'answering';
		setTimeout(() => {
			inputRefs[index]?.focus();
		}, 10);
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
						class={`w-[30rem] ${flashcard.suppliedAnswer === flashcard.back ? 'bg-green-100' : ''}`}
						bind:this={inputRefs[index]}
						bind:value={flashcard.suppliedAnswer}
						on:keydown={(e) => {
							if (e.key === 'Enter') handleSubmitAnswer(flashcard, index);
						}}
					/>
				</div>
			{/if}
			{#if flashcard.status === 'incorrect'}
				<div class={`back rounded-b-md bg-slate-500 px-3 py-2`}>
					<p class="text-red-800">{flashcard.suppliedAnswer}</p>
					{@render FlashcardBack(flashcard)}
					<div class="mt-2 flex justify-end">
						<button
							class="mb-1 rounded border border-slate-600 bg-slate-200 px-2 py-1 text-xs"
							on:click={(e) => handleCloseFlashcard(flashcard, index)}>try again</button
						>
					</div>
				</div>
			{/if}
			{#if flashcard.status === 'correct'}
				<div
					class={`backCorrect flex gap-1 rounded-md border border-green-500 bg-green-100 px-3 py-2`}
				>
					<IconCheck size={22} class="text-green-800" />
					<div>
						{@render FlashcardBack(flashcard)}
					</div>
				</div>
			{/if}
		</div>
	{/each}
</div>
