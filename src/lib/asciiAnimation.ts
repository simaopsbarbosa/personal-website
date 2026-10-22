export type AsciiAnimationOptions = {
	columns?: number;
	fontSize?: number;
	cellWidth?: number;
	lineHeight?: number;
	framesPerSecond?: number;
	shimmer?: number;
	characters?: string;
	color?: string;
	opacity?: number;
};

const defaults = {
	columns: 80,
	fontSize: 10,
	cellWidth: 6,
	lineHeight: 10,
	framesPerSecond: 16,
	shimmer: 1,
	characters: '.:-=+*#%@',
	color: '#000000',
	opacity: 1
} satisfies Required<AsciiAnimationOptions>;

/** Draws a loaded image as subtly animated ASCII. Returns a cleanup function. */
export function animateAsciiImage(
	canvas: HTMLCanvasElement,
	image: HTMLImageElement,
	options: AsciiAnimationOptions = {}
): () => void {
	if (!image.complete || !image.naturalWidth) {
		throw new Error('The source image must be loaded before starting the ASCII animation.');
	}

	const settings = { ...defaults, ...options };
	if (
		settings.columns <= 0 ||
		settings.cellWidth <= 0 ||
		settings.lineHeight <= 0 ||
		settings.framesPerSecond <= 0 ||
		!settings.characters.length
	) {
		throw new Error('ASCII animation dimensions, frame rate, and characters must be nonempty.');
	}
	const columns = Math.max(1, Math.round(settings.columns));
	const rows = Math.max(
		1,
		Math.round(
			(columns * image.naturalHeight * settings.cellWidth) /
				(image.naturalWidth * settings.lineHeight)
		)
	);
	const sample = document.createElement('canvas');
	sample.width = columns;
	sample.height = rows;
	const sampleContext = sample.getContext('2d', { willReadFrequently: true });
	const context = canvas.getContext('2d');
	if (!sampleContext || !context) throw new Error('Canvas 2D is unavailable.');

	sampleContext.drawImage(image, 0, 0, columns, rows);
	const pixels = sampleContext.getImageData(0, 0, columns, rows).data;
	const width = columns * settings.cellWidth;
	const height = rows * settings.lineHeight;
	const pixelRatio = Math.min(window.devicePixelRatio || 1, 2);
	canvas.width = Math.round(width * pixelRatio);
	canvas.height = Math.round(height * pixelRatio);
	context.setTransform(pixelRatio, 0, 0, pixelRatio, 0, 0);
	context.font = `${settings.fontSize}px monospace`;
	context.textBaseline = 'top';

	const motionPreference = window.matchMedia('(prefers-reduced-motion: reduce)');
	const characters = settings.characters;
	let frameId = 0;
	let lastDraw = -Infinity;

	function draw(time: number) {
		if (!context) return;
		context.clearRect(0, 0, width, height);

		for (let y = 0; y < rows; y++) {
			for (let x = 0; x < columns; x++) {
				const pixel = (y * columns + x) * 4;
				const alpha = pixels[pixel + 3] / 255;
				if (alpha < 0.12) continue;

				const brightness =
					(0.2126 * pixels[pixel] + 0.7152 * pixels[pixel + 1] + 0.0722 * pixels[pixel + 2]) / 255;
				const base = Math.round((1 - brightness) * (characters.length - 1));
				const wave = motionPreference.matches
					? 0
					: settings.shimmer *
						(Math.sin(x * 0.19 + y * 0.13 + time * 0.0015) +
							0.5 * Math.sin(x * 0.43 - y * 0.17 + time * 0.001));
				const characterIndex = Math.max(
					0,
					Math.min(characters.length - 1, base + Math.round(wave))
				);

				context.globalAlpha = alpha * settings.opacity;
				context.fillStyle = settings.color;
				context.fillText(
					characters[characterIndex],
					x * settings.cellWidth,
					y * settings.lineHeight
				);
			}
		}
		context.globalAlpha = 1;
	}

	function tick(time: number) {
		if (time - lastDraw >= 1000 / settings.framesPerSecond) {
			draw(time);
			lastDraw = time;
		}
		frameId = requestAnimationFrame(tick);
	}

	function updateMotion() {
		cancelAnimationFrame(frameId);
		lastDraw = -Infinity;
		draw(0);
		if (!motionPreference.matches) frameId = requestAnimationFrame(tick);
	}

	motionPreference.addEventListener('change', updateMotion);
	updateMotion();

	return () => {
		cancelAnimationFrame(frameId);
		motionPreference.removeEventListener('change', updateMotion);
	};
}
