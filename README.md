# DO CLI

Personal CLI to do CLI things without boilerplate.

## Set of commands

### do push "{msg}"

Perform usual line of commands for git push:

```bash
git add .
git commit -m "{msg}"
git push
```

### do dev

Identifies which programming language and which command is used to run project.
And then run project. To avoid process of checking which one to use in which project.

So if project has file `pnpm-lock.json` then:

```bash
pnpm run dev
```

If project has file `npm-lock.json` then:

```bash
npm run dev
```

### do safari

Opens safari browser using [Playwright](https://playwright.dev/docs/browsers#webkit)
