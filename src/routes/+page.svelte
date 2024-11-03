<script>
	import { page } from '$app/stores';
	import flashcards from '../data/flashcards.json';
</script>

<main class="p-6">
	<section class="rounded bg-slate-400 p-6 font-mono">
		<h2 class="mb-3 text-xl">Flashcards</h2>
		<div class="flex gap-3 flex-wrap">
			{#each flashcards as flashcard}
				<div class="w-fit rounded border border-slate-500 p-2 bg-slate-700 text-slate-300 text-sm">
					<p>{flashcard.front}</p>
					<p>{flashcard.back}</p>
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
