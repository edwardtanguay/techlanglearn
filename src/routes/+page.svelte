<script>
	import { page } from '$app/stores';
</script>

<h1>Tech Language Learning</h1>
<p>On this site I record my learning of tech topics in various languages</p>

<div class="markdown">
	{@html $page.data.content}
</div>
