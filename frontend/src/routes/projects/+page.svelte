<script>
    import uniLogo from '$lib/assets/logos/uni.svg';
    import incluiLogo from '$lib/assets/logos/inclui.svg';
    import fallboundLogo from '$lib/assets/logos/fallbound.svg';
    import sixerLogo from '$lib/assets/logos/sixer.png';
    import pingpongLogo from '$lib/assets/logos/pingpong.png';
    import whatimonLogo from '$lib/assets/logos/whatimon.svg';

    const projects = [
        {
            title: 'uni',
            description: 'Uni is a mobile app designed to help students of the University of Porto to manage their academic life.',
            link: 'https://github.com/NIAEFEUP/uni',
            logo: uniLogo
        },
        {
            title: 'Inclui',
            description: 'Inclui is a mobile app for finding and reporting accessibility features in public spaces, powered by community-verified data.',
            link: 'https://github.com/simaopsbarbosa/inclui',
            logo: incluiLogo
        },
        {
            title: 'what i\'m on',
            description: 'what i\'m on is a personal space to showcase your current obsessions and crowdsource your next favorite thing.',
            link: 'https://github.com/peucastro/what-im-on/',
            logo: whatimonLogo
        },
        {
            title: 'PingPong Q&A',
            description: 'PingPong Q&A is a collaborative Q&A platform for structured knowledge sharing, featuring reputation scoring, smart search, and content moderation.',
            link: 'https://github.com/simaopsbarbosa/pingpong',
            logo: pingpongLogo
        },
        {
            title: 'Sixer',
            description: 'Sixer is a collaborative Q&A platform for structured knowledge sharing, featuring reputation scoring, smart search, and content moderation.',
            link: 'https://github.com/simaopsbarbosa/sixer',
            logo: sixerLogo
        },
        {
            title: 'Fallbound',
            description: 'Fallbound is a 2D platformer game where the player must fall down a series of platforms and gain power while defeating enemies. How far can you fall?',
            link: 'https://github.com/simaopsbarbosa/fallbound',
            logo: fallboundLogo
        },
    ];
</script>


<h2 class="centered">My Work</h2>
<p class="centered">The projects I'm proudest of</p>

<div class="mt-12 mx-10">
    <div class="space-y-3">
        {#each projects as project}
            <div class="flex">
                <div class="w-38 flex justify-center items-center">
                    {#if project.logo}
                        <a href="{project.link}"><img src={project.logo} alt="{project.title} logo" class="max-h-10 w-auto max-w-20" /></a>
                    {/if}
                </div>
                <div class="w-9/10 pl-4 border-l border-dotted">
                    <p>{project.description}</p>
                    <a href={project.link} target="_blank" rel="noopener noreferrer" class="hover:underline secondary">({project.link.slice(8)})</a>
                </div>
            </div>
        {/each}
    </div>
</div>