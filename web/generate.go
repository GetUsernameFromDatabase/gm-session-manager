package web

// Regenerate the shadcn-templ script bundle, templ components and the
// Tailwind stylesheet. The Tailwind standalone CLI must be on PATH; no node
// or npm is involved.
//go:generate go tool shadcn-templ bundle --cwd .. --silent
//go:generate go tool templ generate -path .
//go:generate tailwindcss -i assets/css/globals.css -o assets/css/output.css --minify
