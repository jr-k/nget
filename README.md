# nget

*Pronounced "nugget"* 🍗

Expose files from a Docker container via temporary HTTP links for easy download.

> **💡 Tip:** `nget` is integrated with [github.com/jr-k/d4s](https://github.com/jr-k/d4s) when you `shell` into a volume, perfect for quickly grabbing files!

## Quick Start

### Run

**With host network (simpler, recommended):**
```bash
docker run -it --rm --network host -v /path/to/data:/data ghcr.io/jr-k/nget
```

**With port mapping:**
```bash
docker run -it --rm -p 33000-33100:33000-33100 -v /path/to/data:/data ghcr.io/jr-k/nget
```

### Usage

```bash
nget myfile.txt          # Expose a file (expires in 10 min)
nget -e 60 myfile.txt    # Custom expiration (60 seconds)
nget myfolder/           # Expose a directory (auto tar'd)
nget ls                  # List active sessions
nget kill 1              # Kill session by number or UUID
nget prune               # Kill all sessions
```

## Options

| Option | Description |
|--------|-------------|
| `-e <seconds>` | Set expiration time (default: 600) |

## Commands

| Command | Description |
|---------|-------------|
| `ls` | List active sessions |
| `kill <id>` | Kill a session |
| `prune` | Kill all sessions |

## License

MIT
