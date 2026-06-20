<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { adminState } from '$lib/admin.svelte';

	let { children } = $props();
	let isLoading = $state(true);

	onMount(() => {
		const authed = adminState.checkAuth();
		if (!authed) {
			if (window.location.pathname !== '/admin' && window.location.pathname !== '/admin/') {
				goto('/admin');
			}
		}
		isLoading = false;
	});
</script>

{#if isLoading}
	<p class="secondary py-12 text-center">loading...</p>
{:else}
	<div class="mx-10 mt-6">
		{#if adminState.isAuthenticated}
			<div class="mb-2 flex justify-end">
				<button
					onclick={() => adminState.logout()}
					class="secondary cursor-pointer border-0 bg-transparent p-0 hover:underline"
				>
					(logout)
				</button>
			</div>
		{/if}
		{@render children()}
	</div>
{/if}
