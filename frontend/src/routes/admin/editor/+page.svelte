<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';

	const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

	let isEditing = $state(false);
	let originalSlug = $state('');

	// post states
	let title = $state('');
	let slug = $state('');
	let content = $state('<h2>My Article</h2>\n<p>feeling inspired, si?</p>');

	// UI states
	let isSaving = $state(false);
	let errorMessage = $state('');
	let successMessage = $state('');

	// image upload states
	let isUploadingImage = $state(false);
	let uploadedImageUrl = $state('');
	let imageUploadError = $state('');

	onMount(async () => {
		const token = localStorage.getItem('admin_token');
		if (!token) {
			goto('/admin');
			return;
		}

		const slugParam = page.url.searchParams.get('slug');
		if (slugParam) {
			isEditing = true;
			originalSlug = slugParam;
			await fetchPostDetails(slugParam);
		}
	});

	async function fetchPostDetails(slugParam: string) {
		try {
			const res = await fetch(`${API_URL}/posts/${slugParam}`);
			if (res.ok) {
				const post = await res.json();
				title = post.title;
				slug = post.slug;
				content = post.content;
			} else {
				errorMessage = 'failed to fetch post details for editing.';
			}
		} catch (err) {
			errorMessage = 'error connecting to the backend API.';
		}
	}

	async function savePost() {
		if (!title || !content) {
			errorMessage = 'title and content required.';
			return;
		}

		isSaving = true;
		errorMessage = '';
		successMessage = '';

		const token = localStorage.getItem('admin_token');
		const url = isEditing ? `${API_URL}/posts/${originalSlug}` : `${API_URL}/posts`;
		const method = isEditing ? 'PUT' : 'POST';

		try {
			const res = await fetch(url, {
				method,
				headers: {
					'Content-Type': 'application/json',
					'Authorization': `Bearer ${token}`
				},
				body: JSON.stringify({ title, content })
			});

			const result = await res.json();
			if (!res.ok) {
				errorMessage = result.error || 'failed to save post.';
				isSaving = false;
				return;
			}

			successMessage = isEditing ? 'post updated!' : 'post created!';
			setTimeout(() => {
				goto('/admin');
			}, 800);
		} catch (err) {
			errorMessage = 'failed to connect to backend.';
		} finally {
			isSaving = false;
		}
	}

	async function handleImageUpload(e: Event) {
		const target = e.target as HTMLInputElement;
		if (!target.files || target.files.length === 0) return;

		const file = target.files[0];
		isUploadingImage = true;
		imageUploadError = '';
		uploadedImageUrl = '';

		const token = localStorage.getItem('admin_token');
		const formData = new FormData();
		formData.append('image', file);

		try {
			const res = await fetch(`${API_URL}/upload`, {
				method: 'POST',
				headers: {
					'Authorization': `Bearer ${token}`
				},
				body: formData
			});

			const result = await res.json();
			if (!res.ok) {
				imageUploadError = result.error || 'failed to upload image.';
				return;
			}

			uploadedImageUrl = result.url;
			// append image tag in raw HTML
			const imgTag = `<img src="${result.url}" alt="" class="max-w-full my-4" />`;
			content = content + '\n' + imgTag;
		} catch (err) {
			imageUploadError = 'network error uploading image.';
		} finally {
			isUploadingImage = false;
			target.value = '';
		}
	}
</script>

<div class="space-y-6">
	<!-- page header -->
	<div class="flex items-center justify-between border-b border-dashed pb-2">
		<h2>{isEditing ? 'Edit Post' : 'New Post'}</h2>
		<div class="flex gap-4">
			<a href="/admin" class="secondary hover:underline">(cancel)</a>
			<button
				onclick={savePost}
				disabled={isSaving}
				class="secondary cursor-pointer hover:underline p-0 border-0 bg-transparent"
			>
				{isSaving ? '(saving...)' : '(save post)'}
			</button>
		</div>
	</div>

	<!-- feedback -->
	{#if errorMessage}
		<p class="secondary">{errorMessage}</p>
	{/if}
	{#if successMessage}
		<p class="secondary">{successMessage}</p>
	{/if}

	<!-- split screen -->
	<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
		<!-- left side -->
		<div class="space-y-4">
			<!-- title -->
			<div>
				<label for="title" class="secondary block mb-1">title</label>
				<input
					type="text"
					id="title"
					bind:value={title}
					placeholder="post title"
					class="border border-dashed p-2 w-full outline-hidden focus:border-black"
				/>
			</div>

			<!-- image importer -->
			<div class="border border-dashed p-3">
				<label for="image-upload" class="secondary block mb-1">upload files</label>
				<div class="flex items-center gap-4 mt-2">
					<input
						type="file"
						id="image-upload"
						accept="image/*"
						onchange={handleImageUpload}
						disabled={isUploadingImage}
						class=" file:mr-2 file:border file:border-dashed file:bg-transparent file:px-2 file:py-1 file: file:cursor-pointer file:font-semibold hover:file:bg-black hover:file:text-white"
					/>
					{#if isUploadingImage}
						<span class="secondary animate-pulse">uploading...</span>
					{/if}
				</div>
				{#if uploadedImageUrl}
					<p class="secondary mt-2">image tag appended to editor</p>
				{/if}
				{#if imageUploadError}
					<p class="secondary mt-2">{imageUploadError}</p>
				{/if}
			</div>

			<!-- HTML textarea -->
			<div class="flex flex-col h-[50vh]">
				<label for="content" class="secondary block mb-1">HTML code</label>
				<textarea
					id="content"
					bind:value={content}
					placeholder="write HTML here..."
					class="border border-dashed p-3 w-full grow font-mono leading-relaxed outline-hidden focus:border-black resize-none"
				></textarea>
			</div>
		</div>

		<!-- right side -->
		<div class="border border-dashed p-6 max-h-[68vh] overflow-y-auto ">
			<h3 class="secondary border-b border-dashed pb-2 mb-4">live preview</h3>
			<div>
				<h1>{title || 'Untitled Post'}</h1>
				<hr class="my-4 border-dashed" />
				<article class="prose max-w-none">
					{@html content}
				</article>
			</div>
		</div>
	</div>
</div>

<style>
	:global(.prose h1) {
		font-size: 1.8rem;
		font-weight: bold;
		margin-top: 1.5rem;
		margin-bottom: 0.5rem;
	}
	:global(.prose h2) {
		font-size: 1.5rem;
		font-weight: bold;
		margin-top: 1.2rem;
		margin-bottom: 0.4rem;
	}
	:global(.prose p) {
		margin-bottom: 1rem;
		line-height: 1.6;
	}
	:global(.prose code) {
		background-color: #fff;
		padding: 0.2rem 0.4rem;
		font-family: monospace;
	}
	:global(.prose pre) {
		background-color: #000;
		color: #f9fafb;
		padding: 1rem;
		overflow-x: auto;
		margin-bottom: 1rem;
	}
</style>
