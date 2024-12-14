# quickvast

CLI tools for managing Vast.ai instances.

What I want from this tool:

- [ ] Quickly choose machine types from a refined, customized list of options.
  - This would not just be a filter, it would be a customized, weighted scoring system allowing me to quickly choose the best machine for my needs.
  - It would provide warnings for machines about why they may not be the best choice, such as low RAM, low driver versions, etc.
  - It would have a nice, colorful UI allowing me to quickly see the best options.
  - At first this would be a single weighted config, one that could be hard-coded into the tool, but later I could add multiple configs via files perhaps
- [ ] Quickly create a new instance.
  - Using the same feature above, quickly create an instance.
  - Being able to upload scripts and profiles into the machine easily.
  - Being able to choose and configure the 'template' of the machine.
- [ ] Quickly manage instances.
  - Being able to quickly retrieve necessary connection information, open the browser for tooling, and copy authentication details.
  - Being able to quickly SSH into the machine.
- [ ] Quickly destroy instances.
  - Being able to destroy instances quickly.
- [ ] Long term monitoring for pricing
  - My concern is that I might leave an instance running for too long and rack up a huge bill.
  - I would like each instance to have a time limit for my usage, both a soft and hard limit.
  - The soft limit will send me a message on Discord, and the hard limit will destroy the instance after another message.
  -

## About

This is the official Wails React-TS template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
