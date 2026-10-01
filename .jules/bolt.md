## 2024-05-24 - [Over-fetching in SvelteKit Supabase Queries]

**Learning:** When fetching data for list views (like a blog index), it's easy to accidentally fetch the entire row (including large text columns like `content`) when only a few fields (`title`, `slug`, `created_at`) are needed for rendering. This causes unnecessary database load and larger JSON payloads sent to the client.
**Action:** Always explicitly specify required columns in Supabase `.select()` calls (`.select('id, slug, title')`) instead of defaulting to fetching all or mostly all fields, especially for data that isn't rendered in the current view.

## 2024-05-24 - [Lazy Loading Component Images]

**Learning:** List items that contain images (like a list of projects with logos or banners) can quickly degrade initial page load performance if all images try to load simultaneously, even those far below the fold.
**Action:** Add `loading="lazy"` as a default to `<img>` tags inside reusable list item components to ensure only visible images block the initial render.
