using System;
using System.Text;
using uclip;

if (!Console.IsInputRedirected)
{
    Console.WriteLine(
        """
        UCLIP
        
        Description:
            Redirects output of command line tools to the Windows clipboard.
            This text output can then be pasted into other programs.
            
            This tool expects that only valid UTF8 text will be passed onto it.
            To copy files and images use the standard clip.exe bundled with windows.
            
        Examples:
            (Piping)                
            cat file.txt | uclip  Places a copy of the text from stdin into windows clipboard.
            (File Redirection)
            uclip < file.txt      Places a copy of the file content into the windows clipboard.
        """
    );
    return;
}

Console.InputEncoding = Encoding.UTF8;
string input = Console.In.ReadToEnd();
bool success = Clipboard.SetData(input);
if (!success)
    throw new ArgumentException("Failed to set data to clipboard");
