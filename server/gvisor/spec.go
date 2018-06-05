/*
* Honeytrap
* Copyright (C) 2016-2017 DutchSec (https://dutchsec.com/)
*
* This program is free software; you can redistribute it and/or modify it under
* the terms of the GNU Affero General Public License version 3 as published by the
* Free Software Foundation.
*
* This program is distributed in the hope that it will be useful, but WITHOUT
* ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
* FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License for more
* details.
*
* You should have received a copy of the GNU Affero General Public License
* version 3 along with this program in the file "LICENSE".  If not, see
* <http://www.gnu.org/licenses/agpl-3.0.txt>.
*
* See https://honeytrap.io/ for more details. All requests should be sent to
* licensing@honeytrap.io
*
* The interactive user interfaces in modified source and object code versions
* of this program must display Appropriate Legal Notices, as required under
* Section 5 of the GNU Affero General Public License version 3.
*
* In accordance with Section 7(b) of the GNU Affero General Public License version 3,
* these Appropriate Legal Notices must retain the display of the "Powered by
* Honeytrap" logo and retain the original copyright notice. If the display of the
* logo is not reasonably feasible for technical reasons, the Appropriate Legal Notices
* must display the words "Powered by Honeytrap" and retain the original copyright notice.
 */
package gvisor

import (
	"encoding/json"
	"fmt"
	_ "log"

	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func DefaultSpec() *specs.Spec {
	specBytes := []byte(`{
  "ociVersion": "1.0.0",
  "process": {
    "user": {
      "uid": 0,
      "gid": 0
    },
    "args": [
            "dropbear","-F","-p", "8022", "-R", "-E"
    ],
    "env": [
      "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
      "HOSTNAME=7ff99a5a341a"
    ],
    "terminal": true,
"consoleSize": {
"height": 80,
"width": 80
},
    "cwd": "/",
    "capabilities": {
      "bounding": [
        "CAP_CHOWN",
        "CAP_DAC_OVERRIDE",
        "CAP_FSETID",
        "CAP_FOWNER",
        "CAP_MKNOD",
        "CAP_NET_RAW",
        "CAP_SETGID",
        "CAP_SETUID",
        "CAP_SETFCAP",
        "CAP_SETPCAP",
        "CAP_NET_BIND_SERVICE",
        "CAP_SYS_CHROOT",
        "CAP_KILL",
        "CAP_AUDIT_WRITE"
      ],
      "effective": [
        "CAP_CHOWN",
        "CAP_DAC_OVERRIDE",
        "CAP_FSETID",
        "CAP_FOWNER",
        "CAP_MKNOD",
        "CAP_NET_RAW",
        "CAP_SETGID",
        "CAP_SETUID",
        "CAP_SETFCAP",
        "CAP_SETPCAP",
        "CAP_NET_BIND_SERVICE",
        "CAP_SYS_CHROOT",
        "CAP_KILL",
        "CAP_AUDIT_WRITE"
      ],
      "inheritable": [
        "CAP_CHOWN",
        "CAP_DAC_OVERRIDE",
        "CAP_FSETID",
        "CAP_FOWNER",
        "CAP_MKNOD",
        "CAP_NET_RAW",
        "CAP_SETGID",
        "CAP_SETUID",
        "CAP_SETFCAP",
        "CAP_SETPCAP",
        "CAP_NET_BIND_SERVICE",
        "CAP_SYS_CHROOT",
        "CAP_KILL",
        "CAP_AUDIT_WRITE"
      ],
      "permitted": [
        "CAP_CHOWN",
        "CAP_DAC_OVERRIDE",
        "CAP_FSETID",
        "CAP_FOWNER",
        "CAP_MKNOD",
        "CAP_NET_RAW",
        "CAP_SETGID",
        "CAP_SETUID",
        "CAP_SETFCAP",
        "CAP_SETPCAP",
        "CAP_NET_BIND_SERVICE",
        "CAP_SYS_CHROOT",
        "CAP_KILL",
        "CAP_AUDIT_WRITE"
      ]
    },
    "apparmorProfile": "docker-default",
    "oomScoreAdj": 0
  },
  "root": {
    "path": "/chroot"
  },
  "hostname": "7ff99a5a341a",
  "mounts": [
    {
      "destination": "/proc",
      "type": "proc",
      "source": "proc",
      "options": [
        "nosuid",
        "noexec",
        "nodev"
      ]
    },
    {
      "destination": "/dev",
      "type": "tmpfs",
      "source": "tmpfs",
      "options": [
        "nosuid",
        "strictatime",
        "mode=755",
        "size=65536k"
      ]
    },
    {
      "destination": "/dev/pts",
      "type": "devpts",
      "source": "devpts",
      "options": [
        "nosuid",
        "noexec",
        "newinstance",
        "ptmxmode=0666",
        "mode=0620",
        "gid=5"
      ]
    },
    {
      "destination": "/sys",
      "type": "sysfs",
      "source": "sysfs",
      "options": [
        "nosuid",
        "noexec",
        "nodev",
        "ro"
      ]
    },
    {
      "destination": "/sys/fs/cgroup",
      "type": "cgroup",
      "source": "cgroup",
      "options": [
        "ro",
        "nosuid",
        "noexec",
        "nodev"
      ]
    },
    {
      "destination": "/dev/mqueue",
      "type": "mqueue",
      "source": "mqueue",
      "options": [
        "nosuid",
        "noexec",
        "nodev"
      ]
    },
    {
      "destination": "/etc/resolv.conf",
      "type": "bind",
      "source": "/var/lib/docker/containers/7ff99a5a341a21b6a4d76b6f049c09a3bd5d05d22bdb7e6cd47f9e8f5e61d87f/resolv.conf",
      "options": [
        "rbind",
        "rprivate"
      ]
    },
    {
      "destination": "/etc/hostname",
      "type": "bind",
      "source": "/var/lib/docker/containers/7ff99a5a341a21b6a4d76b6f049c09a3bd5d05d22bdb7e6cd47f9e8f5e61d87f/hostname",
      "options": [
        "rbind",
        "rprivate"
      ]
    },
    {
      "destination": "/etc/hosts",
      "type": "bind",
      "source": "/var/lib/docker/containers/7ff99a5a341a21b6a4d76b6f049c09a3bd5d05d22bdb7e6cd47f9e8f5e61d87f/hosts",
      "options": [
        "rbind",
        "rprivate"
      ]
    },
    {
      "destination": "/dev/shm",
      "type": "bind",
      "source": "/var/lib/docker/containers/7ff99a5a341a21b6a4d76b6f049c09a3bd5d05d22bdb7e6cd47f9e8f5e61d87f/shm",
      "options": [
        "rbind",
        "rprivate"
      ]
    }
  ],
  "hooks": {
    "prestart": [
    ]
  },
  "linux": {
    "resources": {
      "devices": [
        {
          "allow": false,
          "access": "rwm"
        },
        {
          "allow": true,
          "type": "c",
          "major": 1,
          "minor": 5,
          "access": "rwm"
        },
        {
          "allow": true,
          "type": "c",
          "major": 1,
          "minor": 3,
          "access": "rwm"
        },
        {
          "allow": true,
          "type": "c",
          "major": 1,
          "minor": 9,
          "access": "rwm"
        },
        {
          "allow": true,
          "type": "c",
          "major": 1,
          "minor": 8,
          "access": "rwm"
        },
        {
          "allow": true,
          "type": "c",
          "major": 5,
          "minor": 0,
          "access": "rwm"
        },
        {
          "allow": true,
          "type": "c",
          "major": 5,
          "minor": 1,
          "access": "rwm"
        },
        {
          "allow": false,
          "type": "c",
          "major": 10,
          "minor": 229,
          "access": "rwm"
        }
      ],
      "memory": {},
      "cpu": {
        "shares": 0
      },
      "pids": {
        "limit": 0
      },
      "blockIO": {
        "weight": 0
      }
    },
    "cgroupsPath": "/docker/7ff99a5a341a21b6a4d76b6f049c09a3bd5d05d22bdb7e6cd47f9e8f5e61d87f",
    "namespaces": [
      {
        "type": "mount"
      },
      {
              "type": "network",
              "path":"/var/run/netns/NS1"
      },
      {
        "type": "uts"
      },
      {
        "type": "pid"
      },
      {
        "type": "ipc"
      }
    ],
    "seccomp": {
      "defaultAction": "SCMP_ACT_ERRNO",
      "architectures": [
        "SCMP_ARCH_X86_64",
        "SCMP_ARCH_X86",
        "SCMP_ARCH_X32"
      ],
      "syscalls": [
        {
          "names": [
            "accept"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "accept4"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "access"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "adjtimex"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "alarm"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "bind"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "brk"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "capget"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "capset"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "chdir"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "chmod"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "chown"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "chown32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "clock_getres"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "clock_gettime"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "clock_nanosleep"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "close"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "connect"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "copy_file_range"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "creat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "dup"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "dup2"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "dup3"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "epoll_create"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "epoll_create1"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "epoll_ctl"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "epoll_ctl_old"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "epoll_pwait"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "epoll_wait"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "epoll_wait_old"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "eventfd"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "eventfd2"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "execve"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "execveat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "exit"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "exit_group"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "faccessat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fadvise64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fadvise64_64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fallocate"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fanotify_mark"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fchdir"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fchmod"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fchmodat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fchown"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fchown32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fchownat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fcntl"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fcntl64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fdatasync"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fgetxattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "flistxattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "flock"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fork"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fremovexattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fsetxattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fstat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fstat64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fstatat64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fstatfs"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fstatfs64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "fsync"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "ftruncate"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "ftruncate64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "futex"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "futimesat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getcpu"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getcwd"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getdents"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getdents64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getegid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getegid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "geteuid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "geteuid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getgid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getgid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getgroups"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getgroups32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getitimer"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getpeername"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getpgid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getpgrp"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getpid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getppid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getpriority"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getrandom"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getresgid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getresgid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getresuid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getresuid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getrlimit"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "get_robust_list"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getrusage"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getsid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getsockname"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getsockopt"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "get_thread_area"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "gettid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "gettimeofday"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getuid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getuid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "getxattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "inotify_add_watch"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "inotify_init"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "inotify_init1"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "inotify_rm_watch"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "io_cancel"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "ioctl"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "io_destroy"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "io_getevents"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "ioprio_get"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "ioprio_set"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "io_setup"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "io_submit"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "ipc"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "kill"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "lchown"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "lchown32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "lgetxattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "link"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "linkat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "listen"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "listxattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "llistxattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "_llseek"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "lremovexattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "lseek"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "lsetxattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "lstat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "lstat64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "madvise"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "memfd_create"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mincore"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mkdir"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mkdirat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mknod"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mknodat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mlock"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mlock2"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mlockall"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mmap"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mmap2"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mprotect"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mq_getsetattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mq_notify"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mq_open"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mq_timedreceive"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mq_timedsend"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mq_unlink"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "mremap"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "msgctl"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "msgget"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "msgrcv"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "msgsnd"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "msync"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "munlock"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "munlockall"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "munmap"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "nanosleep"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "newfstatat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "_newselect"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "open"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "openat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "pause"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "pipe"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "pipe2"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "poll"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "ppoll"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "prctl"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "pread64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "preadv"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "preadv2"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "prlimit64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "pselect6"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "pwrite64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "pwritev"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "pwritev2"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "read"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "readahead"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "readlink"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "readlinkat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "readv"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "recv"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "recvfrom"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "recvmmsg"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "recvmsg"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "remap_file_pages"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "removexattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "rename"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "renameat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "renameat2"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "restart_syscall"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "rmdir"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "rt_sigaction"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "rt_sigpending"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "rt_sigprocmask"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "rt_sigqueueinfo"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "rt_sigreturn"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "rt_sigsuspend"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "rt_sigtimedwait"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "rt_tgsigqueueinfo"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sched_getaffinity"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sched_getattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sched_getparam"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sched_get_priority_max"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sched_get_priority_min"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sched_getscheduler"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sched_rr_get_interval"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sched_setaffinity"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sched_setattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sched_setparam"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sched_setscheduler"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sched_yield"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "seccomp"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "select"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "semctl"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "semget"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "semop"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "semtimedop"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "send"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sendfile"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sendfile64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sendmmsg"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sendmsg"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sendto"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setfsgid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setfsgid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setfsuid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setfsuid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setgid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setgid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setgroups"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setgroups32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setitimer"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setpgid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setpriority"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setregid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setregid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setresgid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setresgid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setresuid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setresuid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setreuid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setreuid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setrlimit"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "set_robust_list"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setsid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setsockopt"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "set_thread_area"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "set_tid_address"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setuid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setuid32"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "setxattr"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "shmat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "shmctl"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "shmdt"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "shmget"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "shutdown"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sigaltstack"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "signalfd"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "signalfd4"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sigreturn"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "socket"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "socketcall"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "socketpair"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "splice"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "stat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "stat64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "statfs"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "statfs64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "symlink"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "symlinkat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sync"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sync_file_range"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "syncfs"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "sysinfo"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "syslog"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "tee"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "tgkill"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "time"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "timer_create"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "timer_delete"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "timerfd_create"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "timerfd_gettime"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "timerfd_settime"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "timer_getoverrun"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "timer_gettime"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "timer_settime"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "times"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "tkill"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "truncate"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "truncate64"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "ugetrlimit"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "umask"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "uname"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "unlink"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "unlinkat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "utime"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "utimensat"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "utimes"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "vfork"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "vmsplice"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "wait4"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "waitid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "waitpid"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "write"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "writev"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "personality"
          ],
          "action": "SCMP_ACT_ALLOW",
          "args": [
            {
              "index": 0,
              "value": 0,
              "op": "SCMP_CMP_EQ"
            }
          ]
        },
        {
          "names": [
            "personality"
          ],
          "action": "SCMP_ACT_ALLOW",
          "args": [
            {
              "index": 0,
              "value": 8,
              "op": "SCMP_CMP_EQ"
            }
          ]
        },
        {
          "names": [
            "personality"
          ],
          "action": "SCMP_ACT_ALLOW",
          "args": [
            {
              "index": 0,
              "value": 131072,
              "op": "SCMP_CMP_EQ"
            }
          ]
        },
        {
          "names": [
            "personality"
          ],
          "action": "SCMP_ACT_ALLOW",
          "args": [
            {
              "index": 0,
              "value": 131080,
              "op": "SCMP_CMP_EQ"
            }
          ]
        },
        {
          "names": [
            "personality"
          ],
          "action": "SCMP_ACT_ALLOW",
          "args": [
            {
              "index": 0,
              "value": 4294967295,
              "op": "SCMP_CMP_EQ"
            }
          ]
        },
        {
          "names": [
            "arch_prctl"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "modify_ldt"
          ],
          "action": "SCMP_ACT_ALLOW"
        },
        {
          "names": [
            "clone"
          ],
          "action": "SCMP_ACT_ALLOW",
          "args": [
            {
              "index": 0,
              "value": 2080505856,
              "op": "SCMP_CMP_MASKED_EQ"
            }
          ]
        },
        {
          "names": [
            "chroot"
          ],
          "action": "SCMP_ACT_ALLOW"
        }
      ]
    },
    "maskedPaths": [
      "/proc/kcore",
      "/proc/latency_stats",
      "/proc/timer_list",
      "/proc/timer_stats",
      "/proc/sched_debug",
      "/proc/scsi",
      "/sys/firmware"
    ],
    "readonlyPaths": [
      "/proc/asound",
      "/proc/bus",
      "/proc/fs",
      "/proc/irq",
      "/proc/sys",
      "/proc/sysrq-trigger"
    ]
  }
}`)

	var spec specs.Spec
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		panic(fmt.Errorf("error unmarshaling spec: %v\n %s", err, string(specBytes)))
	}

	return &spec
}
