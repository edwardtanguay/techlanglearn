<script>
	import { page } from '$app/stores';
</script>

<main class="p-6">
	<header>
		<h1 class="main">Tech Language Learning</h1>
		<p>On this site I record my learning of tech topics while watching video tutorials in numerous languages.</p>
		<p>
			This site was made with <b>SvelteKit 5</b>, see
			<a target="_blank" class="underline" href="https://github.com/edwardtanguay/techlanglearn"
				>repo</a
			>
			and live
			<a target="_blank" class="underline" href="https://techlanglearn.vercel.app">site</a>.
		</p>
		<p>
			Also see my <a target="_blank" class="underline" href="https://tanguay-eu.vercel.app/howtos">other sites</a>.
		</p>
	</header>

	<div class="markdown-tutorial">
		{@html $page.data.content}
	</div>
</main>
