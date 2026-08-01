---
title: Publishing Obsidian Notes to a Website
tags:
  - software
  - obsidian
publish: "true"
---
I spent more time than I'd like to admit managing my blog posts in two places. I'd write them in markdown in Obsidian and then copy and paste them into my websites HTML. It was not pretty... So in an effort to make my life easier I setup this blog! Every single post is automatically imported from my personal Obsidian Vault. This doc will go over how everything works and possible improvements.  
## Tech Stack
- Hugo (any static site generator will do)
- Github (for both storing my notes and hosting this blog via Github Pages)
- Obsidian (to write the notes) and the community Git plugin (to push them to a Github repo)

Every part of the stack is fully replaceable, so if anything happens to any of the above tools, it should be easy enough to find and use an alternative.

![](/attachments/publish_ObsidianFlow.svg)
The basic idea is fairly simple:
1. Push your notes to a private repository
2. Configure your hugo site to periodically scan and pull your private notes 
## Privacy
We do not necessarily want to publish all our notes to our website. My solution for this is by using frontmatter. Frontmatter is a feature that allows you to store metadata in markdown files, and many static site generators work natively with it. For this setup, I include a text property called "publish". If it is set to true, the markdown file will be pulled to the Hugo repo, if it is set to anything else or not set at all, the file will be ignored. If you don't want to risk accidentally publishing a private note, you could always use a separate vault.
## Pushing 
For this step, I use the Obsidian [Git Plugin](https://github.com/Vinzent03/obsidian-git). I have the default hotkey for Save (⌘S) and use a custom hotkey for Commit-And-Sync (⌘⇧S). I also pay for Obsidian Sync so I can easily access my notes on my other devices. 
## Pulling
The Hugo repo is where the bulk of the logic lies! It has a couple of responsibilities: 
1. Periodically scanning the private notes repo  
2. Pulling "published" files
	- Keep in mind, this logic can be quite complex. I'd recommend deleting all files and re-pulling all down, otherwise renames can leave you with duplicates 
3. Generating and building the static site
## Summary
To see a working example, check out this website's [repo](https://github.com/dylanCz/blog). I still have some plans for improvements but this has served me well and I don't foresee any major changes at the time of writing.  
