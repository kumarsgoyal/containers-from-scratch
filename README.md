# containers-from-scratch

Build Your Own Container Using Less than 100 Lines of Go. This is a toy
container build from scratch in Go solely for learning purpose. It uses
namespaces and cgroups.

This will start our container (needs root privilege for creating `cgroup`):

```
sudo su
go run main.go run /bin/bash
```

---

## Sources of the inspiration and information

Building Containers from Scratch with Go by Liz Rice
https://www.safaribooksonline.com/videos/building-containers-from/9781491988404

If you don't have access to safaribooksonline.com, Liz Rice gave talk on the
same topic in several conferences too.
One of them is "GOTO 2018 • Containers From Scratch":
https://www.youtube.com/watch?v=8fi7uSYlOdc

---

## Further reading

### Namespaces in Go

Part 1: Linux Namespaces
https://medium.com/@teddyking/linux-namespaces-850489d3ccf

Part 2: Namespaces in Go - Basics
https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a

Part 3: Namespaces in Go - User
https://medium.com/@teddyking/namespaces-in-go-user-a54ef9476f2a

Part 4: Namespaces in Go - reexec
https://medium.com/@teddyking/namespaces-in-go-reexec-3d1295b91af8

Part 5: Namespaces in Go - Mount
https://medium.com/@teddyking/namespaces-in-go-mount-e4c04fe9fb29

Part 6: Namespaces in Go - Network
https://medium.com/@teddyking/namespaces-in-go-network-fdcf63e76100

Part 7: Namespaces in Go - UTS
https://medium.com/@teddyking/namespaces-in-go-uts-d47aebcdf00e

---

## Bonus tip: Setting up VS Code for cross-platform development

I have used fedora to develop the container from scratch. However, if you are
using a non-Linux OS (like macOS or Windows),the development experience is not
great because a lot of pieces of this application is Linux specific. For
example, calls like `syscall.Sethostname` or the `Cloneflags` field in the
`syscall.SysProcAttr{}` struct is not available in intellisense in VSCode when
the dev environment is not Linux. VS Code will mark those lines as errors,
because they are platform specific and declared in the standard library in Go
for Linux only.

Fortunately the workaround is very simple. Search for `"go.toolsEnvVars"` in VS
Code settings, copy it to User Settings and change it to:

```
    "go.toolsEnvVars": {
        "GOOS": "linux"
    }
```

Restarting VS Code after that will recognize all Linux specific declarations and
will not see them as errors. Go-to-definition will work properly too.

---
