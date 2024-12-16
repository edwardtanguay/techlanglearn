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
	{#each store.tutorials as tutorial}
		<p>{tutorial.title}</p>
	{/each}
	<hr />

	<section class="styled-table">
		<table>
			<thead>
				<tr>
					<th>Name</th>
					<th>Age</th>
					<th>Occupation</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td>John Doe</td>
					<td>29</td>
					<td>Developer</td>
				</tr>
				<tr>
					<td>Jane Smith</td>
					<td>34</td>
					<td>Designer</td>
				</tr>
				<tr>
					<td>Sam Wilson</td>
					<td>41</td>
					<td>Manager</td>
				</tr>
			</tbody>
		</table>
	</section>
</section>
