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
		alert(flashcard.front)
	}
</script>

<div class="area_flashcard">
	<h2 class="mb-6 text-xl">{maxShow} of {store.flashcards.length} flashcards.</h2>

	{#each store.flashcards.slice(0, maxShow) as flashcard, index}
		<div class="mb-6 w-fit">
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<div
				on:click={() => handleToggleFlashcard(flashcard, index)}
				class={`front ${flashcard.status === 'showingFront' ? 'cursor-pointer' : ''} rounded-t-md bg-slate-400 px-3 py-2 front-lang-${flashcard.language}`}
			>
				{flashcard.front}
			</div>
			{#if flashcard.status !== 'showingFront'}
				<div class={`back rounded-b-md bg-slate-500 px-3 py-2`}>
					{#if flashcard.status !== 'answering'}
						<p>{flashcard.back}</p>
						{#if flashcard.extras !== ''}
							<p class="text-sm italic text-gray-400">{flashcard.extras}</p>
						{/if}
					{:else}
						<input bind:this={inputRefs[index]} bind:value={flashcard.suppliedAnswer} on:keydown={(e) => {
							if(e.key === 'Enter') handleSubmitAnswer(flashcard)
						}} />
					{/if}
				</div>
			{/if}
		</div>
	{/each}
</div>
