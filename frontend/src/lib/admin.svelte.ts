import { goto } from '$app/navigation';

export class AdminState {
	#isAuthenticated = $state(false);

	get isAuthenticated() {
		return this.#isAuthenticated;
	}

	set isAuthenticated(value: boolean) {
		this.#isAuthenticated = value;
	}

	login(token: string) {
		localStorage.setItem('admin_token', token);
		this.#isAuthenticated = true;
	}

	logout() {
		localStorage.removeItem('admin_token');
		this.#isAuthenticated = false;
		goto('/');
	}

	checkAuth() {
		const token = localStorage.getItem('admin_token');
		this.#isAuthenticated = !!token;
		return this.#isAuthenticated;
	}
}

export const adminState = new AdminState();
