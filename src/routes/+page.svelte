<script>
	import { page } from '$app/stores';
</script>

<main class="p-6">
	<section class="font-mono bg-slate-400 p-6 rounded">
		<h2 class="text-xl mb-3">Flashcards</h2>
		<div>
			<p>the front</p>
			<p>the back</p>
		</div>
	</section>
	<header>
		<h1 class="main">Tech and Language Learning <span class="text-xs">by <a href="https://tanguay.info" target="_blank" class="underline">Edward Tanguay</a></span></h1>
		<p>On this site I record my learning of tech topics while watching video tutorials in numerous languages.</p>
		<p>
			This site was made with <b>SvelteKit 5</b>, see
			<a target="_blank" class="underline" href="https://github.com/edwardtanguay/techlanglearn"
				>repo</a
			>
			and 
			<a target="_blank" class="underline" href="https://techlanglearn.vercel.app">live site</a>.
		</p>
		<p>
			Also see my <a target="_blank" class="underline" href="https://tanguay-eu.vercel.app/howtos">other sites</a>.
		</p>
	</header>

	<div class="markdown-tutorial">
		{@html $page.data.htmlContent}
	</div>
</main>
