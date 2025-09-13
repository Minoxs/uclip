# UCLIP

Simple clipboard tool that supports UTF8 text.
Designed to work much like clip.exe does.

Redirects output of command line tools to the Windows clipboard.
This text output can then be pasted into other programs.

This tool expects that only valid UTF8 text will be passed onto it.
To copy files and images use the standard clip.exe bundled with windows.
    
# Examples

This tools works like clip does, as shown below.

## Piping         

Places a copy of the text from stdin into windows clipboard.

```bash
echo "Cool Beans ☕" | uclip
```

## File Redirection

Places a copy of the file contents into the windows clipboard.

```bash
uclip < file.txt
```
