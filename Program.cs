using System;
using System.Text;
using uclip;

Console.InputEncoding = Encoding.UTF8;
string input = Console.In.ReadToEnd();
bool success = Clipboard.SetData(input);
if (!success)
    throw new ArgumentException("Failed to set data to clipboard");
