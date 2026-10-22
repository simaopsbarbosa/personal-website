<script lang="ts">
	import { onDestroy } from 'svelte';
	import simao from '$lib/assets/simao.png';
	import { animateAsciiImage } from '$lib/asciiAnimation';

	let portrait: HTMLImageElement;
	let asciiCanvas: HTMLCanvasElement;
	let asciiVisible = $state(false);
	let stopAnimation: (() => void) | undefined;
	let silhouetteAlpha: Uint8Array | undefined;

	function showAscii() {
		if (!portrait.complete || !portrait.naturalWidth || asciiVisible) return;

		try {
			stopAnimation = animateAsciiImage(asciiCanvas, portrait, {
				columns: 80,
				framesPerSecond: 16,
				shimmer: 1,
				color: '#000000'
			});
			asciiVisible = true;
		} catch (error) {
			console.error('Could not render the ASCII portrait:', error);
		}
	}

	function hideAscii() {
		stopAnimation?.();
		stopAnimation = undefined;
		asciiVisible = false;
	}

	function pointerIsOnPortrait(event: PointerEvent) {
		if (!portrait.naturalWidth || !portrait.naturalHeight) return false;

		if (!silhouetteAlpha) {
			const maskCanvas = document.createElement('canvas');
			maskCanvas.width = portrait.naturalWidth;
			maskCanvas.height = portrait.naturalHeight;
			const context = maskCanvas.getContext('2d', { willReadFrequently: true });
			if (!context) return false;

			context.drawImage(portrait, 0, 0);
			const rgba = context.getImageData(0, 0, maskCanvas.width, maskCanvas.height).data;
			silhouetteAlpha = new Uint8Array(maskCanvas.width * maskCanvas.height);
			for (let pixel = 0; pixel < silhouetteAlpha.length; pixel++) {
				silhouetteAlpha[pixel] = rgba[pixel * 4 + 3];
			}
		}

		const bounds = portrait.getBoundingClientRect();
		const x = event.clientX - bounds.left;
		const y = event.clientY - bounds.top;
		if (x < 0 || y < 0 || x >= bounds.width || y >= bounds.height) return false;

		const sourceX = Math.floor((x / bounds.width) * portrait.naturalWidth);
		const sourceY = Math.floor((y / bounds.height) * portrait.naturalHeight);
		return silhouetteAlpha[sourceY * portrait.naturalWidth + sourceX] > 24;
	}

	function handlePointerMove(event: PointerEvent) {
		if (event.pointerType === 'mouse' && pointerIsOnPortrait(event)) showAscii();
		else hideAscii();
	}

	onDestroy(() => {
		stopAnimation?.();
	});
</script>

<div
	class="relative w-fit max-w-[70vw] md:max-w-2xs"
	role="img"
	aria-label={asciiVisible
		? 'Animated ASCII portrait of Simão Barbosa'
		: 'Portrait of Simão Barbosa'}
	onpointermove={handlePointerMove}
	onpointerleave={hideAscii}
>
	<img
		bind:this={portrait}
		class="portrait-layer portrait-source rendering-pixelated block max-w-full"
		class:ascii-visible={asciiVisible}
		src={simao}
		alt=""
		title="hi ;)"
	/>
	<div class="portrait-layer portrait-ascii absolute inset-0" class:ascii-visible={asciiVisible}>
		<canvas bind:this={asciiCanvas} class="h-full w-full" aria-hidden="true"></canvas>
	</div>
</div>

<style>
	.portrait-layer {
		transition: opacity 225ms ease;
	}

	.portrait-source.ascii-visible,
	.portrait-ascii {
		opacity: 0;
	}

	.portrait-ascii.ascii-visible {
		opacity: 1;
	}

	@media (prefers-reduced-motion: reduce) {
		.portrait-layer {
			transition: none;
		}
	}
</style>
