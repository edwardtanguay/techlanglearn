<script lang="ts">
	import { getStore } from '../../store.svelte';
	import './styles.scss';

	const store = getStore();
</script>

<header>
	<h1 class="main">Tutorials</h1>
	<p class="test123">There are {store.tutorials.length} tutorials.</p>
	{#each store.tutorials as tutorial}
		<p>{tutorial.title}</p>
	{/each}
</header>
