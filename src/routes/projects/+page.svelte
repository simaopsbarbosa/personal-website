<script lang="ts">
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
	import type { ProjectLink } from '$lib/components/projects/types';

	// Each label is displayed as written; the first link is also used by the project image.
	const projects: {
		title: string;
		description: string;
		links: [ProjectLink, ...ProjectLink[]];
		logo: string;
		banner?: string;
	}[] = [
		{
			title: 'uni',
			description:
				'I co-coordinate uni at NIFEUP. The app gives University of Porto students one place to check class schedules, exam dates, grades, and campus resources.',
			links: [
				{ label: 'github', href: 'https://github.com/NIAEFEUP/uni' },
				{
					label: 'google play',
					href: 'https://play.google.com/store/apps/details?id=pt.up.fe.ni.uni&hl=pt_PT&gl=US&pli=1&pcampaignid=pcampaignidMKT-Other-global-all-co-prtnr-py-PartBadge-Mar2515-1'
				},
				{ label: 'app store', href: 'https://apps.apple.com/pt/app/uni/id1585951339' }
			],
			logo: uniLogo,
			banner: uniBanner
		},
		{
			title: 'Inclui',
			description:
				'I built Inclui with four classmates at FEUP. The app lets people find and report accessibility features like wheelchair ramps and Braille menus so they can choose places that meet their needs.',
			links: [{ label: 'github', href: 'https://github.com/simaopsbarbosa/inclui' }],
			logo: incluiLogo,
			banner: incluiBanner
		},
		{
			title: "what i'm on",
			description:
				"what i'm on is a personal space to showcase your current obsessions and crowdsource your next favorite thing. Built during the Shift APPens 26 hackathon, winner of the Best Design award.",
			links: [
				{ label: 'github', href: 'https://github.com/peucastro/what-im-on/' },
				{ label: 'website', href: 'https://whatimon.vercel.app/' }
			],
			logo: whatimonLogo,
			banner: whatimonBanner
		},
		{
			title: 'PingPong Q&A',
			description:
				'PingPong Q&A is a collaborative Q&A platform for structured knowledge sharing, featuring reputation scoring, smart search, and content moderation.',
			links: [{ label: 'github', href: 'https://github.com/simaopsbarbosa/pingpong' }],
			logo: pingpongLogo
		},
		{
			title: 'Sixer',
			description:
				'Sixer is a freelancing platform where freelancers can offer their services and customers can browse and purchase them.',
			links: [{ label: 'github', href: 'https://github.com/simaopsbarbosa/sixer' }],
			logo: sixerLogo
		},
		{
			title: 'Fallbound',
			description:
				'Fallbound is a 2D platformer game where the player must fall down a series of platforms and gain power while defeating enemies. How far can you fall?',
			links: [
				{ label: 'github', href: 'https://github.com/simaopsbarbosa/fallbound' },
				{ label: 'play web port', href: 'https://fallbound.simaob.com' }
			],
			logo: fallboundLogo
		}
	];
</script>

<h1 class="centered">My Work</h1>
<p class="centered">The projects I'm proudest of</p>

<div class="mx-10 my-12">
	<div class="space-y-6">
		{#each projects as project (project.title)}
			{#if project.banner}
				<FeaturedProject {...project} banner={project.banner} />
			{:else}
				<Project {...project} />
			{/if}
		{/each}
	</div>
</div>
