<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { supabase } from '$lib/supabase';
	import { adminState } from '$lib/admin.svelte';

	let isEditing = $state(false);
	let originalSlug = $state('');
	let postID = $state<number | null>(null);

	// post states
	let title = $state('');
	let slug = $state('');
	let content = $state('<h2>My Article</h2>\n<p>feeling inspired, si?</p>');
	let draft = $state(true);
	let originalDraft = $state(true);

	// UI states
	let isSaving = $state(false);
	let errorMessage = $state('');
	let successMessage = $state('');

	// image upload states
	let isUploadingImage = $state(false);
	let uploadedImageUrl = $state('');
	let imageUploadError = $state('');

	onMount(async () => {
		const authenticated = await adminState.checkAuth();
		if (!authenticated) {
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
			const { data: post, error } = await supabase
				.from('posts')
				.select('*')
				.eq('slug', slugParam)
				.single();

			if (error || !post) {
				errorMessage = 'failed to fetch post details for editing.';
			} else {
				postID = post.id;
				title = post.title;
				slug = post.slug;
				content = post.content;
				draft = post.draft ?? false;
				originalDraft = post.draft ?? false;
			}
		} catch (err) {
			errorMessage = 'error connecting to Supabase.';
		}
	}

	function slugify(input: string): string {
		let base = input.toLowerCase().trim();
		base = base.replace(/[^a-z0-9]+/g, '-');
		base = base.replace(/^-+|-+$/g, '');
		return base || 'post';
	}

	async function generateUniqueSlug(titleStr: string, excludeId?: number): Promise<string> {
		const baseSlug = slugify(titleStr);
		let candidate = baseSlug;
		let suffix = 2;

		while (true) {
			let query = supabase.from('posts').select('id').eq('slug', candidate);
			if (excludeId) {
				query = query.neq('id', excludeId);
			}
			const { data, error } = await query;
			if (!data || data.length === 0) {
				return candidate;
			}
			candidate = `${baseSlug}-${suffix}`;
			suffix++;
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

		try {
			const finalSlug = await generateUniqueSlug(title, postID || undefined);
			const now = new Date().toISOString();

			let error;
			if (isEditing && postID) {
				const updatePayload: {
					title: string;
					content: string;
					slug: string;
					draft: boolean;
					updated_at: string;
					created_at?: string;
				} = {
					title,
					content,
					slug: finalSlug,
					draft,
					updated_at: now
				};

				// If changing from draft to published, set created_at to now
				if (originalDraft && !draft) {
					updatePayload.created_at = now;
				}

				const { error: err } = await supabase.from('posts').update(updatePayload).eq('id', postID);
				error = err;
			} else {
				const insertPayload: {
					title: string;
					content: string;
					slug: string;
					draft: boolean;
					created_at?: string;
					updated_at?: string;
				} = {
					title,
					content,
					slug: finalSlug,
					draft
				};

				// If created directly as published, set created_at and updated_at
				if (!draft) {
					insertPayload.created_at = now;
					insertPayload.updated_at = now;
				}

				const { error: err } = await supabase.from('posts').insert([insertPayload]);
				error = err;
			}

			if (error) {
				errorMessage = error.message || 'failed to save post.';
				isSaving = false;
				return;
			}

			successMessage = isEditing ? 'post updated!' : 'post created!';
			setTimeout(() => {
				goto('/admin');
			}, 800);
		} catch {
			errorMessage = 'failed to connect to Supabase.';
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

		try {
			const ext = file.name.split('.').pop();
			const fileName = `${Date.now()}.${ext}`;
			const filePath = `${fileName}`;

			const { data, error } = await supabase.storage.from('blog-images').upload(filePath, file);

			if (error) {
				imageUploadError = error.message || 'failed to upload image.';
				return;
			}

			const {
				data: { publicUrl }
			} = supabase.storage.from('blog-images').getPublicUrl(filePath);

			uploadedImageUrl = publicUrl;
			// append image tag in raw HTML
			const imgTag = `<img src="${publicUrl}" alt="" class="max-w-full my-4" />`;
			content = content + '\n' + imgTag;
		} catch (err) {
			imageUploadError = 'network error uploading image.';
		} finally {
			isUploadingImage = false;
			target.value = '';
		}
	}

	let textareaElement = $state<HTMLTextAreaElement>();

	function insertHTML(prefix: string, suffix: string) {
		if (!textareaElement) {
			content = content + prefix + suffix;
			return;
		}
		const start = textareaElement.selectionStart;
		const end = textareaElement.selectionEnd;
		const text = content;
		const selected = text.substring(start, end);
		const replacement = prefix + selected + suffix;
		const before = text.substring(0, start);
		const after = text.substring(end, text.length);
		content = before + replacement + after;

		// focus back and set selection
		setTimeout(() => {
			if (!textareaElement) return;
			textareaElement.focus();
			if (selected) {
				textareaElement.selectionStart = start;
				textareaElement.selectionEnd = start + replacement.length;
			} else {
				textareaElement.selectionStart = textareaElement.selectionEnd = start + prefix.length;
			}
		});
	}
</script>

<div class="flex flex-col space-y-4 pb-4 md:h-[calc(100vh-5rem)]">
	<!-- page header -->
	<div class="flex flex-none items-center justify-between border-b border-dashed pb-2">
		<h2>{isEditing ? 'Edit Post' : 'New Post'}</h2>
		<div class="flex gap-4">
			<a href="/admin" class="secondary hover:underline">(cancel)</a>
			<button
				onclick={savePost}
				disabled={isSaving}
				class="secondary cursor-pointer border-0 bg-transparent p-0 hover:underline"
			>
				{isSaving ? '(saving...)' : '(save post)'}
			</button>
		</div>
	</div>

	<!-- feedback -->
	{#if errorMessage}
		<p class="secondary flex-none">{errorMessage}</p>
	{/if}
	{#if successMessage}
		<p class="secondary flex-none">{successMessage}</p>
	{/if}

	<!-- split screen -->
	<div class="grid min-h-0 flex-1 grid-cols-1 gap-6 md:grid-cols-2">
		<!-- left side -->
		<div class="flex flex-col space-y-4 md:h-full md:min-h-0">
			<!-- title -->
			<div class="flex-none">
				<label for="title" class="secondary mb-1 block">title</label>
				<input
					type="text"
					id="title"
					bind:value={title}
					placeholder="post title"
					class="w-full border border-dashed p-2 outline-hidden focus:border-black"
				/>
			</div>

			<!-- draft checkbox -->
			<div class="flex flex-none items-center gap-2">
				<input
					type="checkbox"
					id="draft"
					bind:checked={draft}
					class="h-4 w-4 cursor-pointer border border-dashed accent-black"
				/>
				<label for="draft" class="secondary cursor-pointer select-none">
					save as draft (hidden from public blog)
				</label>
			</div>

			<!-- image importer -->
			<div class="flex-none border border-dashed p-3">
				<label for="image-upload" class="secondary mb-1 block">upload files</label>
				<div class="mt-2 flex items-center gap-4">
					<input
						type="file"
						id="image-upload"
						accept="image/*"
						onchange={handleImageUpload}
						disabled={isUploadingImage}
						class=" file: file:mr-2 file:cursor-pointer file:border file:border-dashed file:bg-transparent file:px-2 file:py-1 file:font-semibold hover:file:bg-black hover:file:text-white"
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
			<div class="flex h-[50vh] flex-col md:min-h-0 md:flex-1">
				<div class="mb-1 flex flex-none flex-wrap items-center justify-between gap-2">
					<label for="content" class="secondary">HTML code</label>
					<div class="flex flex-wrap gap-2 text-xs">
						<button
							onclick={() => insertHTML('<h2>', '</h2>')}
							class="secondary cursor-pointer border-0 bg-transparent p-0 hover:underline"
							>(h2)</button
						>
						<button
							onclick={() => insertHTML('<p>', '</p>')}
							class="secondary cursor-pointer border-0 bg-transparent p-0 hover:underline"
							>(p)</button
						>
						<button
							onclick={() => insertHTML('<strong>', '</strong>')}
							class="secondary cursor-pointer border-0 bg-transparent p-0 hover:underline"
							>(bold)</button
						>
						<button
							onclick={() => insertHTML('<em>', '</em>')}
							class="secondary cursor-pointer border-0 bg-transparent p-0 hover:underline"
							>(italic)</button
						>
						<button
							onclick={() =>
								insertHTML('<a href="https://example.com" class="dotted-underline">', '</a>')}
							class="secondary cursor-pointer border-0 bg-transparent p-0 hover:underline"
							>(link)</button
						>
						<button
							onclick={() => insertHTML('<blockquote>', '</blockquote>')}
							class="secondary cursor-pointer border-0 bg-transparent p-0 hover:underline"
							>(quote)</button
						>
						<button
							onclick={() => insertHTML('<ul>\n\t<li>', '</li>\n\t<li></li>\n</ul>')}
							class="secondary cursor-pointer border-0 bg-transparent p-0 hover:underline"
							>(list)</button
						>
						<button
							onclick={() => insertHTML('<code>', '</code>')}
							class="secondary cursor-pointer border-0 bg-transparent p-0 hover:underline"
							>(code)</button
						>
					</div>
				</div>
				<textarea
					id="content"
					bind:this={textareaElement}
					bind:value={content}
					placeholder="write HTML here..."
					class="min-h-0 w-full flex-1 resize-none border border-dashed p-3 font-mono leading-relaxed outline-hidden focus:border-black"
				></textarea>
			</div>
		</div>

		<!-- right side -->
		<div
			class="max-h-[68vh] min-h-0 overflow-y-auto border border-dashed p-6 md:h-full md:max-h-none"
		>
			<h3 class="secondary mb-4 border-b border-dashed pb-2">live preview</h3>
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
		color: black;
		font-size: 1.8rem;
		font-weight: bold;
		margin-top: 1.5rem;
		margin-bottom: 0.5rem;
	}
	:global(.prose h2) {
		color: black;
		font-size: 1.5rem;
		font-weight: bold;
		margin-top: 1.2rem;
		margin-bottom: 0.4rem;
	}
	:global(.prose p) {
		color: black;
		margin-bottom: 1rem;
		line-height: 1.6;
	}
	:global(.prose code) {
		background-color: black;
		color: white;
		padding: 0.2rem 0.4rem;
		font-family: monospace;
		font-weight: 400;
	}
	:global(.prose pre) {
		background-color: black;
		border-radius: 0px !important;
		color: #f9fafb;
		padding: 1rem;
		overflow-x: auto;
		margin-bottom: 1rem;
	}
</style>
