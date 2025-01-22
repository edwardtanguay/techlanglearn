<script lang="ts">
	import IconLinkedIn from '../../components/IconLinkedIn.svelte';
	import IconYoutube from '../../components/IconYoutube.svelte';
	import { getStore } from '../../store.svelte';
	import type { Tutorial } from '../../types';
	import './styles.scss';

	const store = getStore();

	const handleToggleSortColumn = (fieldIdCode: string) => {
		store.toggleSortColumn(fieldIdCode);
	};

	handleToggleSortColumn('rank');

	const getStatusClass = (tutorial: Tutorial) => {
		if (['finished', 'discontinued'].includes(tutorial.status)) {
			return tutorial.status;
		} else {
			return '';
		}
	};

	let showingControls = false;
	const topics = store.getSelectableTopics();
	let selectedTopic = topics[0];
	const selectTopic = (topic: string) => {
		store.filterTutorials(topic);
	};
	setTimeout(() => {
		showingControls = true;
	}, 1000);
</script>

<section class="pageTutorials">
	<header class="flex gap-3">
		<h1 class="main">{store.filteredTutorials.length} Tutorials</h1>
		{#if showingControls}
			<select
				bind:value={selectedTopic}
				class="inline-block rounded-md border border-gray-300 bg-gray-50 px-2 py-1 text-sm text-gray-700 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500"
			>
				{#each topics as topic}
					<option
						value={topic}
						selected={selectedTopic === topic}
						on:click={() => selectTopic(topic)}>{topic}</option
					>
				{/each}
			</select>
		{/if}
	</header>
	<hr />

	<section class="styled-table">
		<table>
			<thead>
				<tr>
					<th class="cursor-pointer" on:click={() => handleToggleSortColumn('rank')}>Rank</th>
					<th class="cursor-pointer" on:click={() => handleToggleSortColumn('year')}>Year</th>
					<th class="cursor-pointer" on:click={() => handleToggleSortColumn('language')}
						>Language</th
					>
					<th>Topics</th>
					<th>Category</th>
					<th>Duration</th>
					<th>Title</th>
				</tr>
			</thead>
			<tbody>
				{#each store.filteredTutorials as tutorial}
					<tr class={getStatusClass(tutorial)}>
						<td>{tutorial.rank.toFixed(2)}</td>
						<td>{tutorial.year === 0 ? '' : tutorial.year}</td>
						<td>{tutorial.language}</td>
						<td>{tutorial.topics}</td>
						<td title={tutorial.category?.description}
							>{tutorial.category !== null ? tutorial.category.name : ''}</td
						>
						<td>{tutorial.duration}</td>
						<td>
							<p class="flex gap-2">
								<a class="underline" target="_blank" href={tutorial.url}>{tutorial.title}</a>
								{#if tutorial.platform === 'linkedInLearning'}
									<IconLinkedIn size={22} />
								{/if}
								{#if tutorial.platform === 'youtube'}
									<IconYoutube width={31} />
								{/if}
								{#if tutorial.status === 'finished'}
									<span class="-ml-[.1rem] font-bold text-green-800">finished</span>
								{/if}
								{#if tutorial.status === 'discontinued'}
									<span class="-ml-[.1rem] font-bold italic text-red-800">discontinued</span>
								{/if}
								{#if !['finished', 'discontinued'].includes(tutorial.status)}
									<span class="-ml-[.1rem] text-orange-700">{tutorial.status}</span>
								{/if}
							</p>
							<p class="text-xs text-orange-800">{tutorial.fileIdCode}.md</p>
							<p class="text-xs italic">{tutorial.description}</p>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</section>
</section>
