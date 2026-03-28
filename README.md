# Personal Website & Blog

A full-stack personal website and blog platform built with a Go backend and a SvelteKit frontend.

## Project Structure

- **`/backend`**: Go REST API using the Gin framework and SQLite.
- **`/frontend`**: SvelteKit 5 application styled with TailwindCSS.

## Tech Stack

### Frontend
- **Framework:** [SvelteKit 5](https://svelte.dev/)
- **Styling:** [TailwindCSS 4](https://tailwindcss.com/)
- **Type Safety:** TypeScript
- **Testing:** Vitest & Playwright

### Backend
- **Language:** [Go](https://go.dev/)
- **API Framework:** [Gin](https://gin-gonic.com/)
- **Database:** SQLite
- **Documentation:** Swagger / Open API 2.0
- **Authentication:** Bcrypt password hashing

## Getting Started

### Prerequisites
- Go 1.21+
- Node.js 20+
- pnpm (recommended) or npm

### Running the Backend
1. Navigate to the backend directory:
   ```bash
   cd backend
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Run the server:
   ```bash
   go run .
   ```
   The API will be available at `http://localhost:8080`.
   Swagger documentation is available at `http://localhost:8080/swagger/index.html`.

### Running the Frontend
1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Install dependencies:
   ```bash
   pnpm install
   ```
3. Start the development server:
   ```bash
   pnpm dev
   ```
   The website will be available at `http://localhost:5173`.

## License
[MIT](LICENSE)
