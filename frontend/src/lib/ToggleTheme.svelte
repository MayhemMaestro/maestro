<script lang="ts">
	export let dark: boolean;

	const setDark = (e: KeyboardEvent | MouseEvent, chooseDark: boolean) => {
		if (e.type === 'click' || (e as KeyboardEvent).key === 'Enter') {
			dark = chooseDark;
			localStorage.setItem('dark', chooseDark.toString());
			const htmlElement = document.getElementsByTagName('html')[0];

			if (!chooseDark) {
				htmlElement.classList.replace('bg-zinc-900', 'bg-zinc-50');
				htmlElement.classList.replace('text-slate-300', 'text-stone-700');
			} else {
				htmlElement.classList.replace('bg-zinc-50', 'bg-zinc-900');
				htmlElement.classList.replace('text-stone-700', 'text-slate-300');
			}
		}
	};
</script>

<div class="cursor-pointer rounded-lg w-full h-min flex flex-row">
	<svg
		role="button"
		tabindex={1}
		on:click={(e) => setDark(e, false)}
		on:keydown={(e) => setDark(e, false)}
		xmlns="http://www.w3.org/2000/svg"
		width="30px"
		height="30px"
		viewBox="0 0 24 24"
		class="text-yellow-600 dark:text-yellow-300 ml-8 hover:text-yellow-800 dark:hover:text-yellow-600"
		{...$$props}
	>
		<g fill="none" stroke="currentColor">
			<circle cx="12" cy="12" r="3.5" />
			<path
				stroke-linecap="round"
				d="M12 5V3m0 18v-2m4.95-11.95l1.414-1.414M5.636 18.364L7.05 16.95M19 12h2M3 12h2m11.95 4.95l1.414 1.414M5.636 5.636L7.05 7.05"
			/>
		</g>
	</svg>

	<svg
		role="button"
		tabindex={2}
		on:click={(e) => setDark(e, true)}
		on:keydown={(e) => setDark(e, true)}
		xmlns="http://www.w3.org/2000/svg"
		width="30px"
		height="30px"
		viewBox="0 0 256 256"
		class="text-blue-600 dark:text-blue-300 ml-8 hover:text-blue-800 dark:hover:text-blue-600"
		{...$$props}
	>
		<path
			fill="currentColor"
			d="M232.13 143.64a6 6 0 0 0-6-1.49a90.07 90.07 0 0 1-112.27-112.3a6 6 0 0 0-7.49-7.48a102.88 102.88 0 0 0-51.89 36.31a102 102 0 0 0 142.84 142.84a102.88 102.88 0 0 0 36.31-51.89a6 6 0 0 0-1.5-5.99m-42 48.29a90 90 0 0 1-126-126a90.9 90.9 0 0 1 35.52-28.27a102.06 102.06 0 0 0 118.69 118.69a90.9 90.9 0 0 1-28.24 35.58Z"
		/>
	</svg>
</div>
