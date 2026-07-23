<script>
	import uniLogo from '$lib/assets/logos/uni.svg';
	import uniBanner from '$lib/assets/projects/uni.webp';
	import incluiLogo from '$lib/assets/logos/inclui.svg';
	import incluiBanner from '$lib/assets/projects/inclui.webp';
	import fallboundLogo from '$lib/assets/logos/fallbound.svg';
	import sixerLogo from '$lib/assets/logos/sixer.png';
	import pingpongLogo from '$lib/assets/logos/pingpong.png';
	import whatimonLogo from '$lib/assets/logos/whatimon.svg';
	import whatimonBanner from '$lib/assets/projects/whatimon.webp';

	import FeaturedProject from '$lib/components/projects/FeaturedProject.svelte';
	import Project from '$lib/components/projects/Project.svelte';

	const projects = [
		{
			title: 'uni',
			description:
				'Uni is a mobile app designed to help students of the University of Porto to manage their academic life.',
			link: 'https://github.com/NIAEFEUP/uni',
			logo: uniLogo,
			banner: uniBanner
		},
		{
			title: 'Inclui',
			description:
				'Inclui is a mobile app for finding and reporting accessibility features in public spaces, powered by community data.',
			link: 'https://github.com/simaopsbarbosa/inclui',
			logo: incluiLogo,
			banner: incluiBanner
		},
		{
			title: "what i'm on",
			description:
				"what i'm on is a personal space to showcase your current obsessions and crowdsource your next favorite thing. Built during the Shift APPens 26 hackathon, winner of the Best Design award.",
			link: 'https://github.com/peucastro/what-im-on/',
			logo: whatimonLogo,
			banner: whatimonBanner
		},
		{
			title: 'PingPong Q&A',
			description:
				'PingPong Q&A is a collaborative Q&A platform for structured knowledge sharing, featuring reputation scoring, smart search, and content moderation.',
			link: 'https://github.com/simaopsbarbosa/pingpong',
			logo: pingpongLogo
		},
		{
			title: 'Sixer',
			description:
				'Sixer is a freelancing platform where freelancers can offer their services and customers can browse and purchase them.',
			link: 'https://github.com/simaopsbarbosa/sixer',
			logo: sixerLogo
		},
		{
			title: 'Fallbound',
			description:
				'Fallbound is a 2D platformer game where the player must fall down a series of platforms and gain power while defeating enemies. How far can you fall?',
			link: 'https://github.com/simaopsbarbosa/fallbound',
			logo: fallboundLogo
		}
	];
</script>

<h2 class="centered">My Work</h2>
<p class="centered">The projects I'm proudest of</p>

<div class="mx-10 my-12">
	<div class="space-y-6">
		{#each projects as project}
			{#if project.banner}
				<FeaturedProject {...project} />
			{:else}
				<Project {...project} />
			{/if}
		{/each}
	</div>
</div>
