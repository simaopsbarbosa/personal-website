// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
import type { MetaTagsProps } from 'svelte-meta-tags';

declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		interface PageData {
			pageMetaTags?: Readonly<MetaTagsProps>;
		}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
