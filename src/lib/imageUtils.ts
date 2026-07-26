// converts image file to webp with compression
export async function processImageForUpload(file: File, quality = 0.8): Promise<File> {
	if (!file.type.startsWith('image/')) {
		return file;
	}

	const isGif = file.type === 'image/gif' || file.name.toLowerCase().endsWith('.gif');
	if (isGif) {
		return file;
	}

	const isSvg = file.type === 'image/svg+xml' || file.name.toLowerCase().endsWith('.svg');
	if (isSvg) {
		return file;
	}

	if (
		typeof window === 'undefined' ||
		typeof document === 'undefined' ||
		typeof Image === 'undefined'
	) {
		return file;
	}

	return new Promise((resolve) => {
		const img = new Image();
		const objectUrl = URL.createObjectURL(file);

		img.onload = () => {
			URL.revokeObjectURL(objectUrl);

			try {
				const canvas = document.createElement('canvas');
				const width = img.naturalWidth || img.width;
				const height = img.naturalHeight || img.height;

				if (!width || !height) {
					resolve(file);
					return;
				}

				canvas.width = width;
				canvas.height = height;

				const ctx = canvas.getContext('2d');
				if (!ctx) {
					resolve(file);
					return;
				}

				ctx.drawImage(img, 0, 0);

				canvas.toBlob(
					(blob) => {
						if (!blob) {
							resolve(file);
							return;
						}

						const lastDot = file.name.lastIndexOf('.');
						const nameWithoutExt = lastDot !== -1 ? file.name.substring(0, lastDot) : file.name;
						const newFileName = `${nameWithoutExt}.webp`;

						const convertedFile = new File([blob], newFileName, {
							type: 'image/webp',
							lastModified: Date.now()
						});

						resolve(convertedFile);
					},
					'image/webp',
					quality
				);
			} catch {
				resolve(file);
			}
		};

		img.onerror = () => {
			URL.revokeObjectURL(objectUrl);
			resolve(file);
		};

		img.src = objectUrl;
	});
}
