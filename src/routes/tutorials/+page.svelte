<script lang="ts">
	import { getStore } from '../../store.svelte';
	import './styles.scss';

	const store = getStore();
</script>

<section class="pageTutorials">
	<header>
		<h1 class="main">Tutorials</h1>
		<p>There are {store.tutorials.length} tutorials.</p>
	</header>
	<hr />

	<section class="styled-table">
		<table>
			<thead>
				<tr>
					<th>Rank</th>
					<th>Year</th>
					<th>Language</th>
					<th>Topics</th>
					<th>Title</th>
					<th>Platform</th>
				</tr>
			</thead>
			<tbody>
				{#each store.tutorials as tutorial}
					<tr>
						<td>{tutorial.rank}</td>
						<td>{tutorial.year}</td>
						<td>{tutorial.language}</td>
						<td>{tutorial.topics}</td>
						<td>
							<p><a class="underline" target="_blank" href={tutorial.url}>{tutorial.title}</a></p>
							<p class="text-xs text-orange-800">{tutorial.fileIdCode}.md</p>
						</td>
						<td>{tutorial.platform}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</section>
</section>
