<script lang="ts">
	import { onMount } from 'svelte';
	import { formatDate } from '$lib/utils';
	import { adminState } from '$lib/admin.svelte';
	import { supabase } from '$lib/supabase';

	let password = $state('');
	let loginError = $state('');
	let isLoggingIn = $state(false);

	let posts = $state<any[]>([]);
	let isFetchingPosts = $state(true);
	let dashboardError = $state('');

	onMount(async () => {
		const authenticated = await adminState.checkAuth();
		if (authenticated) {
			await fetchPosts();
		} else {
			isFetchingPosts = false;
		}
	});

	async function login(e: SubmitEvent) {
		e.preventDefault();
		if (!password) return;

		isLoggingIn = true;
		loginError = '';

		try {
			const email = import.meta.env.VITE_ADMIN_EMAIL || 'admin@example.com';
			const { error } = await supabase.auth.signInWithPassword({
				email,
				password
			});

			if (error) {
				loginError = error.message || 'invalid password';
				isLoggingIn = false;
				return;
			}

			await fetchPosts();
		} catch (err) {
			loginError = 'connection failed. is Supabase configured correctly?';
		} finally {
			isLoggingIn = false;
		}
	}

	async function fetchPosts() {
		isFetchingPosts = true;
		dashboardError = '';
		try {
			const { data, error } = await supabase
				.from('posts')
				.select('id, slug, title, created_at, updated_at')
				.order('created_at', { ascending: false });

			if (error) {
				dashboardError = error.message || 'failed to fetch posts';
			} else {
				posts = data || [];
			}
		} catch (err) {
			dashboardError = 'failed to load posts from Supabase.';
		} finally {
			isFetchingPosts = false;
		}
	}

	async function deletePost(slug: string, title: string) {
		if (!confirm(`delete post "${title}"?`)) return;

		try {
			const { error } = await supabase
				.from('posts')
				.delete()
				.eq('slug', slug);

			if (error) {
				alert(`failed to delete post: ${error.message || 'unknown error'}`);
			} else {
				await fetchPosts();
			}
		} catch (err) {
			alert('failed to connect to Supabase');
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
