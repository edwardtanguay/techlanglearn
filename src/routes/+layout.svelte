<script lang="ts">
	import '../app.scss';
	import { page } from '$app/stores';
</script>

<nav>
	<ul class="flex gap-3 bg-slate-300 px-2 py-1 pl-6">
		<li><a href="/" class={$page.url.pathname === '/' ? 'active' : ''}>Home</a></li>
		<li><a href="/about" class={$page.url.pathname === '/about' ? 'active' : ''}>About</a></li>
	</ul>
</nav>

<main class="p-6">
	<slot />
</main>
