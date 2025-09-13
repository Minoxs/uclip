using System.Runtime.InteropServices;

namespace uclip;

public static partial class Clipboard {
    [LibraryImport("user32.dll")]
    [return: MarshalAs(UnmanagedType.Bool)]
    private static partial bool OpenClipboard(IntPtr hWnd);

    [LibraryImport("user32.dll")]
    private static partial IntPtr SetClipboardData(uint uFormat, IntPtr data);

    static Clipboard() {
        bool result = OpenClipboard(IntPtr.Zero);
        if (!result)
            throw new ArgumentException("Failed to initialize clipboard");
    }

    public static bool SetData(string value) {
        const int cfUnicode = 13;
        using var data = DisposableString.Unicode(value);
        return SetClipboardData(cfUnicode, data) != 0;
    }
}
