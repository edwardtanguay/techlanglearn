<script lang="ts">
	import '../app.scss';
	import { page } from '$app/stores';
	import { getStore } from '../store.svelte';

	const store = getStore();

	const reloadData = async () => {
		store.setPageStatus('loading');
		const response = await fetch('/api/pd');
		if (response.ok) {
			store.setPageStatus('ready');
		} else {
			const error = await response.json();
			store.setPageStatus('error');
			store.setErrorMessage(error.message);
		}
	};
</script>

<nav class="flex justify-between bg-slate-300">
	<ul class="flex gap-3 px-2 py-1 pl-6">
		<li><a href="/" class={$page.url.pathname === '/' ? 'active' : ''}>Stats</a></li>
		<li><a href="/tutorials" class={$page.url.pathname === '/tutorials' ? 'active' : ''}>Tutorials</a></li>
		<li><a href="/flashcards" class={$page.url.pathname === '/flashcards' ? 'active' : ''}>Flashcards</a></li>
		<li><a href="/about" class={$page.url.pathname === '/about' ? 'active' : ''}>About</a></li>
	</ul>
	<div class="flex items-center gap-3">
		{#if store.siteLocation === 'dev'}
			<button class="rounded border border-slate-600 bg-slate-400 px-1 text-xs" onclick={reloadData}
				>Parse data</button
			>
		{/if}
		<p class="flex items-center text-xs">{store.siteLocation}</p>
		<p class="mr-3 flex items-center text-xs">v0.003</p>
	</div>
</nav>

{#if store.pageStatus === 'loading'}
	<div class="overlay">
		<div class="overlay-text">Parsing data...</div>
	</div>
{/if}
<main class={`p-6 ${store.pageStatus === 'loading' ? 'blurArea' : ''}`}>
	<slot />
</main>
