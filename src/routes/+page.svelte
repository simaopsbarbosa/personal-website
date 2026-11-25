<script lang="ts">
	import Portrait from '$lib/components/Portrait.svelte';

	let { data } = $props();
</script>

<div class="centered flex flex-col gap-6 md:mt-10">
	<h1 class="centered">Simão Barbosa</h1>

	<!-- picture of me -->
	<Portrait />

	<p class="max-w-sm px-4 text-center">
		{data.pageMetaTags.description}
	</p>

	<p>
		co-coordinator of <a href="https://apps.apple.com/pt/app/uni/id1585951339">uni</a> @
		<a href="https://nifeup.com/">NIFEUP</a>
	</p>

	<div class="flex gap-4" data-nosnippet>
		<a class="dotted-underline" href="https://github.com/simaopsbarbosa">github</a>
		<a class="dotted-underline" href="https://www.linkedin.com/in/simaopsbarbosa/">linkedin</a>
		<a class="dotted-underline" href="mailto:simaopsbarbosa@gmail.com">email</a>
	</div>
</div>
