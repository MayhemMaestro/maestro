<script lang="ts">
	import Icon from '@iconify/svelte';
	import IconifyIcon from '@iconify/svelte';
	import Toggle from '$lib/ToggleTheme.svelte';
	export let dark: boolean;
</script>

<nav
	class="bg-stone-200 dark:text-stone-300 text-stone-800 dark:bg-stone-950 border-r-2 border-stone-900
    dark:border-stone-200 h-full w-44 fixed shadow-lg hover:shadow-stone-900 hover:dark:shadow-stone-200 duration-200"
>
	<div class="flex flex-col pt-4 h-full">
		<a class="w-full nav-item z-10 text-cyan-600 hover:text-orange-600 duration-400 ml-4" href="/">
			<Icon icon="fluent:music-note-2-24-regular" height="25px" class="rounded-lg" />
			<p class="pl-2 dark:text-stone-300 text-stone-800">Home</p>
		</a>
		<a class="w-full nav-item z-10 text-lime-600 hover:text-orange-600 duration-400 ml-4" href="/">
			<Icon icon="ic:baseline-history" height="25px" class="rounded-lg" />
			<p class="pl-2 dark:text-stone-300 text-stone-800">History</p>
		</a>

		<a class="w-full nav-item z-10 text-rose-600 hover:text-orange-600 duration-400 ml-4" href="/">
			<IconifyIcon icon="fluent-emoji-high-contrast:magic-wand" height="25px" class="rounded-lg" />
			<p class="pl-2 dark:text-stone-300 text-stone-800">Configure</p>
		</a>
		<div class="w-full nav-item !mt-auto z-10">
			<div>
				<Toggle bind:dark />
			</div>
		</div>
	</div>
</nav>

<style lang="postcss">
	.nav-item {
		@apply flex;
		@apply mb-2;
		@apply mt-2;
		@apply pl-2;
	}
	.nav-item *:hover {
		@apply cursor-pointer;
		color: #ff8400;
		transition-duration: 200ms;
	}
</style>
