# 🚢 ship-htmx-app

> Scaffold and ship HTMX + Express apps fast — from your terminal.

---

## Install

```bash
go install github.com/pnpancholi/ship-htmx-app@latest
```

> Requires Go 1.21+

---

## Quick Start

```bash
ship-htmx-app new myapp
```

Follow the prompts to pick your CSS framework, then:

```bash
cd myapp
npm install
npm run dev
```

Your app is running at `http://localhost:3000` 🎉

---

## What gets scaffolded

```
myapp/
├── server.js
├── package.json
├── index.html      ← HTMX wired up, CSS included
└── src/
    └── public/
        ├── css/
        ├── js/
        └── partials/
```

---

## CSS Options

| Option | Details |
|---|---|
| None | Plain CSS, zero dependencies |
| Pico CSS | Minimal, semantic — perfect for HTMX |
| Tailwind v4 | Utility classes, no build step |

---

## Roadmap

- [x] Express + HTMX scaffolding
- [x] CSS framework selection
- [ ] Better Auth + Drizzle integration
- [ ] Hono backend support
- [ ] Go backend support

---

## Built with

- [Go](https://go.dev) + [Cobra](https://github.com/spf13/cobra)
- [Huh](https://github.com/charmbracelet/huh) — interactive prompts
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) — terminal styling


