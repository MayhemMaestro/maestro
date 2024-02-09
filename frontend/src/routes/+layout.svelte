<script lang="ts">
	import '../app.css';
	import Navbar from '$lib/Navbar.svelte';
	import { onMount } from 'svelte';
	let dark = false;
	onMount(() => {
		let pref = window.matchMedia('(prefers-color-scheme: dark)');
		let localDark = localStorage.getItem('dark');
		if (localDark === 'true') {
			dark = true;
		} else if (localDark === 'false') {
			dark = false;
		} else if (pref.matches) {
			dark = true;
		}
		const htmlElement = document.getElementsByTagName('html')[0];

		if (!dark) {
			htmlElement.classList.replace('bg-zinc-900', 'bg-zinc-50');
			htmlElement.classList.replace('text-slate-300', 'text-stone-700');
		} else {
			htmlElement.classList.replace('bg-zinc-50', 'bg-zinc-900');
			htmlElement.classList.replace('text-stone-700', 'text-slate-300');
		}
	});
</script>

<main class="{dark ? 'dark' : ''} ">
	<Navbar bind:dark />

	<div
		class=" ml-44 bg-zinc-50 dark:bg-zinc-900 dark:text-stone-300 text-stone-700 h-screen
        transition-all duration-500
      "
	>
		<slot />
	</div>
</main>
