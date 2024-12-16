<script lang="ts">
	import { getStore } from '../../store.svelte';

	const store = getStore();
</script>

<header>
	<h1 class="main">Tutorials</h1>
	<p>There are {store.tutorials.length} tutorials.</p>
</header>
