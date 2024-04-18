# Cast CLI

Personal CLI to do CLI things without boilerplate.

## Development

Build command:

```
go build -o bin/cast
```

If CLI isn't updated in Windows system, then make sure it's correct one by running:

```powershell
Get-Command cast
```

## Set of commands

### cast push "{msg}"

Perform usual line of commands for git push:

```bash
git add .
git commit -m "{msg}"
git push
```

### cast dev

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

### cast safari

Opens safari browser using [Playwright](https://playwright.dev/docs/browsers#webkit)

## Default Outputs

```
PS C:\Work\rk\cast> git add .
PS C:\Work\rk\cast> git commit -m "cast"
[main be614b9] cast
 1 file changed, 1 insertion(+), 1 deletion(-)
PS C:\Work\rk\cast> git push
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 20 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 289 bytes | 289.00 KiB/s, done.
Total 3 (delta 2), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (2/2), completed with 2 local objects.
remote: This repository moved. Please use the new location:
remote:   https://github.com/roman-koshchei/cast-cli.git
To https://github.com/roman-koshchei/do.git
   5da2c33..be614b9  main -> main
PS C:\Work\rk\cast>
```
