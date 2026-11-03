package web

// Regenerate the shadcn-templ script bundle, format and compile the views'
// templates, and rebuild the Tailwind stylesheet. The Tailwind standalone CLI
// must be on PATH; no node or npm is involved. Only views is formatted: the
// components belong to the shadcn-templ CLI.
//go:generate go tool shadcn-templ bundle --cwd ../.. --silent
//go:generate go tool templ fmt views
//go:generate go tool templ generate -path .
//go:generate tailwindcss -i assets/css/globals.css -o assets/css/output.css --minify
