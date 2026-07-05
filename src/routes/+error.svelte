<script lang="ts">
	import { page } from '$app/state';
</script>

<div class="centered mt-20 flex flex-col gap-6">
	<h1 class="centered text-4xl font-bold">{page.status || 404}</h1>

	<p class="max-w-md text-center">
		{#if page.status === 404}
			The page you are looking for doesn't exist.
		{:else}
			Something went wrong. Please try again later.
		{/if}
	</p>

	<div>
		<a class="dotted-underline" href="/">Go back home</a>
	</div>
</div>
