<script lang="ts">
	import { getStore } from '../../store.svelte';
	import './styles.scss';

	const store = getStore();

	const handleToggleSortColumn = (fieldIdCode: string) => {
		store.toggleSortColumn(fieldIdCode);
	};
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
					<th class="cursor-pointer" on:click={() => handleToggleSortColumn('rank')}>Rank</th>
					<th class="cursor-pointer" on:click={() => handleToggleSortColumn('year')}>Year</th>
					<th class="cursor-pointer" on:click={() => handleToggleSortColumn('language')}>Language</th>
					<th>Topics</th>
					<th>Platform</th>
					<th>Title</th>
					<th>Status</th>
				</tr>
			</thead>
			<tbody>
				{#each store.tutorials as tutorial}
					<tr>
						<td>{tutorial.rank}</td>
						<td>{tutorial.year}</td>
						<td>{tutorial.language}</td>
						<td>{tutorial.topics}</td>
						<td>{tutorial.platform}</td>
						<td>
							<p><a class="underline" target="_blank" href={tutorial.url}>{tutorial.title}</a></p>
							<p class="text-xs text-orange-800">{tutorial.fileIdCode}.md</p>
						</td>
						<td>{tutorial.status}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</section>
</section>
