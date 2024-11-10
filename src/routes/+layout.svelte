<script lang="ts">
	import '../app.scss';
	import { page } from '$app/stores';
</script>

<nav class="flex justify-between bg-slate-300 ">
	<ul class="flex gap-3 px-2 py-1 pl-6">
		<li><a href="/" class={$page.url.pathname === '/' ? 'active' : ''}>Home</a></li>
		<li><a href="/about" class={$page.url.pathname === '/about' ? 'active' : ''}>About</a></li>
	</ul>
	<p class="mr-3 text-xs flex items-center">v0.001</p>
</nav>

<main class="p-6">
	<slot />
</main>
