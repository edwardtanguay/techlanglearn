<script lang="ts">
	interface Item {
		id: number;
		label: string;
	}

	let items: Item[] = [
		{ id: 1, label: 'Input 1' },
		{ id: 2, label: 'Input 2' },
		{ id: 3, label: 'Input 3' }
	];

	let inputRefs: Record<number, HTMLInputElement | null> = {};

	function focusInput(id: number) {
		inputRefs[id]?.focus();
	}
</script>

<div>
	{#each items as item}
		<div>
			<label for={String(item.id)}>{item.label}</label>
			<input id={String(item.id)} bind:this={inputRefs[item.id]} placeholder={item.label} />
			<button on:click={() => focusInput(item.id)}>Focus {item.label}</button>
		</div>
	{/each}
</div>

<style>
	div {
		margin-bottom: 10px;
	}

	label {
		margin-right: 5px;
	}

	input {
		margin-right: 10px;
		padding: 5px;
	}

	button {
		padding: 5px 10px;
	}
</style>
