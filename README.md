# JCommandChain (JCC)

JCC is A Simple CLI Tool That Allows you to alias Commands/scripts or run multiple of them concurrently

> **Note**
> this is Only Tested and Used on Windows On My Side

### Building the Executable
```bash
go build -o jcc.exe
```
Now Add The Executable to Path

### Usage
Creating The Config File
if Exe is Added To Path
```
  jcc --init
```

if Exe is not Added To Path
```
  ./jcc.exe --init
```
##### Empty Config Created
```json
  {
    aliases: {},
    multiple: {}
  }
```

### The Config {jcc.config.json}
Example
```json
  {
    aliases: {
      "client": "LiveReloadWebServer 'path/client' --port 1200 -useSsl --useLiveReload"
      "tailwind": "npx tailwindcss -i ./client/input.css -o ./client/output.css --watch"
    },
    multiple: {
      "run": ["client","tailwind"] 
    }
  }
```
##### Donts in Config
- Dont Have Same Names In both Aliaes and Multiple
```json
  {
    aliases: {
      "client": "LiveReloadWebServer 'path/client' --port 1200 -useSsl --useLiveReload"
      "tailwind": "npx tailwindcss -i ./client/input.css -o ./client/output.css --watch"
    },
    multiple: {
      // dont
      "client": ["client","tailwind"] 
    }
  }
```
- Dont use Commands Directly in Multiple like Shown in "run" command
```json
  {
    aliases: {
      "client": "LiveReloadWebServer 'path/client' --port 1200 -useSsl --useLiveReload"
      "tailwind": "npx tailwindcss -i ./client/input.css -o ./client/output.css --watch"
    },
    multiple: {
      // dont
      "run": ["client","npx tailwindcss -i ./client/input.css -o ./client/output.css --watch"] 
    }
 }
```

##### Using Your Commands
Using the first Correct Config given above in Example 
- To only use One command
```bash
  jcc client
```
- To run Multiple Commands directly
```bash
  jcc client tailwind
```
- To Run Multiple Commands Defind in Config
```
  jcc run 
```

