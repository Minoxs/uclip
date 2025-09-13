using System;
using System.Runtime.InteropServices;

namespace uclip;

public class DisposableString : IDisposable
{
    private readonly IntPtr _ptr;

    private DisposableString(IntPtr ptr)
    {
        _ptr = ptr;
    }

    public static DisposableString Unicode(string value)
    {
        IntPtr data = Marshal.StringToHGlobalUni(value);
        return new DisposableString(data);
    }

    public static implicit operator IntPtr(DisposableString str)
    {
        return str._ptr;
    }

    public void Dispose()
    {
        GC.SuppressFinalize(this);
        Marshal.FreeHGlobal(_ptr);
    }
}
