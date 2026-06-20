import { goto } from '$app/navigation';
import { supabase } from './supabase';

export class AdminState {
	#isAuthenticated = $state(false);
	#isInitialized = $state(false);

	constructor() {
		if (typeof window !== 'undefined') {
			supabase.auth.onAuthStateChange((event, session) => {
				this.#isAuthenticated = !!session;
				this.#isInitialized = true;
			});
		}
	}

	get isAuthenticated() {
		return this.#isAuthenticated;
	}

	get isInitialized() {
		return this.#isInitialized;
	}

	async checkAuth() {
		const {
			data: { session }
		} = await supabase.auth.getSession();
		this.#isAuthenticated = !!session;
		this.#isInitialized = true;
		return this.#isAuthenticated;
	}

	async logout() {
		await supabase.auth.signOut();
		this.#isAuthenticated = false;
		goto('/');
	}
}

export const adminState = new AdminState();
