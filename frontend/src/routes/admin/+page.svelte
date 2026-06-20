<script lang="ts">
	import { onMount } from 'svelte';
	import { formatDate } from '$lib/utils';
	import { adminState } from '$lib/admin.svelte';

	const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

	let password = $state('');
	let loginError = $state('');
	let isLoggingIn = $state(false);

	let posts = $state<any[]>([]);
	let isFetchingPosts = $state(true);
	let dashboardError = $state('');

	onMount(async () => {
		if (adminState.checkAuth()) {
			await fetchPosts();
		}
	});

	async function login(e: SubmitEvent) {
		e.preventDefault();
		if (!password) return;

		isLoggingIn = true;
		loginError = '';

		try {
			const response = await fetch(`${API_URL}/auth/login`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ password })
			});

			const result = await response.json();
			if (!response.ok) {
				loginError = result.error || 'invalid password';
				isLoggingIn = false;
				return;
			}

			adminState.login(result.token);
			await fetchPosts();
		} catch (err) {
			loginError = 'connection failed. is the server running?';
		} finally {
			isLoggingIn = false;
		}
	}

	async function fetchPosts() {
		isFetchingPosts = true;
		dashboardError = '';
		try {
			const res = await fetch(`${API_URL}/posts`);
			if (res.ok) {
				posts = await res.json();
			} else {
				const err = await res.json();
				if (res.status !== 404) {
					dashboardError = err.error || 'failed to fetch posts';
				} else {
					posts = [];
				}
			}
		} catch (err) {
			dashboardError = 'failed to load posts from API.';
		} finally {
			isFetchingPosts = false;
		}
	}

	async function deletePost(slug: string, title: string) {
		if (!confirm(`delete post "${title}"?`)) return;

		const token = localStorage.getItem('admin_token');
		try {
			const res = await fetch(`${API_URL}/posts/${slug}`, {
				method: 'DELETE',
				headers: {
					'Authorization': `Bearer ${token}`
				}
			});

			if (res.ok) {
				await fetchPosts();
			} else {
				const err = await res.json();
				alert(`failed to delete post: ${err.error || 'unknown error'}`);
			}
		} catch (err) {
			alert('failed to connect to backend');
		}
	}
</script>

{#if !adminState.isAuthenticated}
	<!-- login screen -->
	<div class="centered py-24">
		<form onsubmit={login} class="flex flex-col gap-3 w-full max-w-xs">
			<input
				type="password"
				bind:value={password}
				placeholder="guess the magic word"
				required
				class="border border-dashed p-2 font-mono outline-hidden transition"
			/>

			{#if loginError}
				<p class="secondary text-center">nice try, but no</p>
			{/if}

			<button
				type="submit"
				disabled={isLoggingIn}
				class="cursor-pointer border py-2 font-mono transition hover:bg-black hover:text-white disabled:opacity-50"
			>
				{isLoggingIn ? '(thinking...)' : 'submit'}
			</button>
		</form>
	</div>
{:else}
	<!-- dashboard screen -->
	<div>
		<div class="mb-8 flex items-center justify-between">
			<div>
				<h2>Admin</h2>
			</div>
			<a href="/admin/editor" class="secondary hover:underline">(+ new post)</a>
		</div>

		{#if dashboardError}
			<p class="secondary mb-4">{dashboardError}</p>
		{/if}

		{#if isFetchingPosts}
			<p class="secondary text-center py-8">fetching posts...</p>
			{:else if posts.length === 0}
			<!-- no posts -->
			<div class="border border-dashed p-12 text-center">
				<p class="secondary mb-4">no posts found</p>
				<a href="/admin/editor" class="secondary hover:underline">(create first post)</a>
			</div>
		{:else}
			<div class="space-y-4">
				{#each posts as post, index}
					<!-- post -->
					<div class="flex items-center justify-between py-2 border-b border-dashed">
						<div class="flex flex-col">
							<a href="/blog/{post.slug}">{post.title}</a>
							<div class="flex gap-2 items-center">
								<span class="secondary">{formatDate(post.created_at)}</span>
								<span class="secondary">| {post.slug}</span>
							</div>
						</div>
						
						<div class="flex gap-2">
							<a href="/admin/editor?slug={post.slug}" class="secondary hover:underline">(edit)</a>
							<button
								onclick={() => deletePost(post.slug, post.title)}
								class="secondary cursor-pointer hover:underline p-0 border-0 bg-transparent"
							>
								(delete)
							</button>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</div>
{/if}
