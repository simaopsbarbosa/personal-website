# Project Setup Phase

- Initialize a new Git repository for version control.
- Define your database schema. You will need a Users table (for your admin account) and a Posts table.
- Select your database. For a fast, simple personal blog, SQLite is highly efficient. Use PostgreSQL if you anticipate complex relational data later.
- Initialize your Go module using `go mod init <module-name>`.
- Create a `.env` file to securely store your database connection string, server port, and JWT secret. Add this file to `.gitignore`.
- Manually insert your hashed password into the database. Since you are the only user, building a registration system is an unnecessary security risk.

---

# Backend Checklist & Endpoints

## Public Endpoints

- `GET /api/posts` — Retrieve a paginated list of blog posts.
- `GET /api/posts/:slug` — Retrieve a single blog post by its URL-friendly slug.

## Protected Endpoints (Admin Only)

- `POST /api/auth/login` — Verify credentials and return an authentication token.
- `POST /api/posts` — Create a new blog post.
- `PUT /api/posts/:slug` — Update an existing blog post.
- `DELETE /api/posts/:slug` — Remove a blog post.

## Security & Performance Tasks

- Use `bcrypt` or `argon2` to hash your password. Never store plain text.
- Issue JWTs (JSON Web Tokens) or secure sessions upon login. Set them as `HttpOnly`, `Secure`, and `SameSite=Strict` cookies to prevent XSS and CSRF attacks.
- Implement CORS (Cross-Origin Resource Sharing) middleware. Restrict access strictly to your future Svelte domain.
- Add rate limiting, particularly on the `/api/auth/login` route, to mitigate brute-force attacks.
- Implement input validation for all POST and PUT requests before they interact with the database.

---

# Frontend & Integration Phase

- Initialize the frontend using SvelteKit. This provides Server-Side Rendering (SSR) out of the box, which is critical for SEO and fast initial load times.
- Set up your SvelteKit environment variables to point to the Go backend URL.
- Build the public-facing portfolio pages and the blog layout.
- Create an API client utility in Svelte. This should be a wrapper around the native `fetch` API that automatically handles errors and attaches your `HttpOnly` cookie or auth token to requests hitting protected endpoints.
- Build the hidden login page and the admin dashboard for managing posts.
- Integrate a markdown parser (like `marked` or `mdsvex`) in Svelte so you can write posts in markdown and render them as HTML.
- Implement route protection in SvelteKit (`hooks.server.js` or `hooks.server.ts`) to ensure non-authenticated users are redirected away from your admin dashboard.
- Configure SvelteKit to prerender your static portfolio pages during the build process to guarantee the fastest possible delivery to users.
