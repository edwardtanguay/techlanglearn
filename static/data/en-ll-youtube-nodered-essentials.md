# Node-RED Essentials

https://www.youtube.com/watch?v=ksGeUD26Mw0&list=PLyNBB9VCLmo1hyO-4fIZ08gqFcXBkHy-6

- duration: 01:00:00
- language: en
- topics: nodered
- rank: 4.3
- description: 19 videos to show the basics of Node-RED
- year: 2020
- STATUS: started

## Introduction, 0:56, 2025-01-18

- can run on Rasberry Pi
- low code
- removes boiling plate

## Editor Components, 2:31, 2025-01-18

- you can hide and show flows, or disable them
- there is context data
- shift-click to select all connected nodes

## Creating a flow, 1:42, 2025-01-19

https://www.youtube.com/watch?v=46Ak61c_ymc&list=PLyNBB9VCLmo1hyO-4fIZ08gqFcXBkHy-6&index=3

- start with inject/debug
- you can drop a node onto the wire

## Wiring nodes, 2:08, 2025-01-19

https://www.youtube.com/watch?v=_ST7fiBLlfw&list=PLyNBB9VCLmo1hyO-4fIZ08gqFcXBkHy-6&index=4

- you can release the mouse over any part of the target node
- hold CTRL down to chain nodes
- CTRL-click on workspace is add dialog
- CTRL-click on wire to add node between two nodes
- detach wire: SHIFT-click
- delete wire: SHIFT-click then release on empty space
- can SHIFT-click to copy all wires

## Editing nodes, 2:47, 2025-01-20

https://www.youtube.com/watch?v=JLrMswaViC4&list=PLyNBB9VCLmo1hyO-4fIZ08gqFcXBkHy-6&index=5

- you can change the icons
- blue dot: changes to be deployed
- you can disable so messages won't run through it, will be blocked

## Deploying flows, 1:52, nnn

https://www.youtube.com/watch?v=xjUnVzbMUCI&list=PLyNBB9VCLmo1hyO-4fIZ08gqFcXBkHy-6&index=6

- deploy: restart just modified flows

## Information sidebar, 1:34, nnn

https://www.youtube.com/watch?v=agOKgoOOI1M&list=PLyNBB9VCLmo1hyO-4fIZ08gqFcXBkHy-6&index=7

- you can revel the node in the workspace
- get help on each kind of node on the side

## Debug sidebar, 3:04, 2025-01-20

https://www.youtube.com/watch?v=TDxL02SInps&list=PLyNBB9VCLmo1hyO-4fIZ08gqFcXBkHy-6&index=8

- click to jump
- click date/milliseconds to display as date
- copy the path so that you can paste it in code
- selected nodes didn't seem to work
- you can open debug in its own window

## Configuration Nodes Sidebar, 1:18, 2025-01-20

- mqtt broker node, not shown
- can select config nodes upper right
- can select unused

## Context Sidebar, 1:36, 2025-01-20

- node, flow, global
- to see contents, you have to click refresh
- how to set the values

```
context.set("name", "Hans");
context.flow.set("count", 3);
global.set("version", 2.1);
```

## Subflows, 2:57, 2025-01-20

- you convert two nodes into a subflow
- you then use variables
- difficult, look into this

## Working with Context, 4:44, 2025-01-20

- decide on the scope
- you can get outside the subflow global to global
- there is a component to save context to a file
- easily done with change node
- remember you are passing objects by reference

## Introduction to Messages, 2:20, 2025-01-20

- a message is an object
- usually has payload
- can also send JSON

## Working with Messages, 2:49, 2025-01-20

- to change messages, use the change node, but for more complex changes, the function node
- change can also access flow and global spaces
- can also use jsonata

## Message Sequences, 4:15, nnn




## VOCAB - ITALIAN

```
it can be executed
può essere eseguito
2025-01-18 23:50:05
```
