<script lang="ts">
import * as dataModel from '../../dataModel';
import './styles.scss';

const stats = dataModel.getStats()

</script>

<section class="statsArea w-fit rounded bg-slate-200 border-2 border-slate-300 p-6 font-mono">
	<p>Total days: <span class="highlight">{stats.totalDays}</span></p>
	<p>Total time watching tutorials: <span class="highlight">{stats.totalDuration}</span></p>
	<p>Average time per day watching tutorials: <span class="highlight bg-yellow-200 px-1">{stats.averageDurationPerDay}</span></p>
	<ul class="mt-3">
		{#each stats.timeUnits as timeUnit}
			<li>{timeUnit.calendarDate} - {timeUnit.duration}</li>
		{/each}
	</ul>
</section>