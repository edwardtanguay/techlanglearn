<script lang="ts">
	import '../app.scss';
	import { page } from '$app/stores';
	import { getStore } from '../store';

	const store = getStore();
</script>

<nav class="flex justify-between bg-slate-300">
	<ul class="flex gap-3 px-2 py-1 pl-6">
		<li><a href="/" class={$page.url.pathname === '/' ? 'active' : ''}>Home</a></li>
		<li><a href="/about" class={$page.url.pathname === '/about' ? 'active' : ''}>About</a></li>
	</ul>
	<div class="flex gap-3">
		<p class="text-xs items-center flex">{store.siteLocation}</p>
		<p class="mr-3 flex items-center text-xs">v0.001</p>
	</div>
</nav>

<main class="p-6">
	<slot />
</main>
