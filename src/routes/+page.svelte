<script>
	import { page } from '$app/stores';
</script>

<main class="p-6">
	<header class="mb-6">
		<h1 class="main">Tech Language Learning</h1>
		<p>On this site I record my learning of tech topics in various languages</p>
	</header>

	<div class="markdown">
		{@html $page.data.content}
	</div>
</main>
