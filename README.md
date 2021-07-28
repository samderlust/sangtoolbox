# Sangtoolbox

This is an opensource project to create an toolbox that help you saving time with boiledplate. Whatever your projects is.
The tool box helps run commandline to create your project and then create folders and files as follow the structure you provide

At the moment, sangtoolbox only support creating flutter project.

# Install

1. Clone the project

   ```
   	git clone https://github.com/samderlust/sangtoolbox.git
   ```

2. provide the template file in `template` folder. Must be a json file
3. run `make sangtoolbox` in the root folder to install the toolbox into your system

# Avaialable commands

| command        | usage                                                  | note                                                                                 |
| -------------- | ------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| flutter_create | `sangtool flutter_create <name> --template=<template>` | if `--template` is not provide, the `example.json` will be used as default template. |
|                |                                                        |
