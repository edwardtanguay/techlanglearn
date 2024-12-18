<script lang="ts">
	import * as dataModel from '../../db/dataModel';
	import './styles.scss';

	const stats = dataModel.getStats();
	const flashcards = dataModel.getFlashcards();
</script>

<section class="statsArea w-fit rounded border-2 border-slate-300 bg-slate-200 p-6 font-mono">
	<p>Total days: <span class="highlight">{stats.totalDays}</span></p>
	<p>
		Average time per day watching tutorials: <span class="highlight bg-green-300 px-1"
			>{stats.averageDurationPerDay}</span
		>
	</p>
	<p>
		Total time watching tutorials since Oct 28: <span class="highlight">{stats.totalDuration}</span>
	</p>
	<p>Number of flashcards recorded: <span class="highlight">{flashcards.length}</span></p>
	<ul class="mt-3">
		{#each stats.timeUnits as timeUnit}
			<li class={timeUnit.duration === '00:00:00' ? 'text-red-700' : 'text-green-700'}>
				{timeUnit.calendarDate} - {timeUnit.duration}
			</li>
		{/each}
	</ul>
</section>
