# nFetch

<img src="./screenshot.png" />

nFetch is a dependency-free, fast system information fetching tool aiming to be quite customizable.

### Customization

#### Ascii-art
The ascii-art must be set in `~/.config/nfetch/art.txt`.

#### Colors
The colors can be set at the top of the `main` function.

### Why?
I wanted a system fetching tool for the look. Sadly, the ones i saw where not what i wanted and neofetch is too slow for my liking. This is where go and boredom came into place.

### Goal
My goal was to not use any dependencies. I wanted to make sure that the performance does not suffer from shell commands or similar.

### Like it?
If you like this project and want to support me, please leave me a star on this repo. I would really appreciate it ❤️

### Build
Git clone:
```sh
# Clone the repo
git clone https://github.com/NeuroException/nfetch

# Navigate into it
cd nfetch

# Build the project
go build main.go

# Rename the outcome
mv main nfetch

# Install it system-wide
sudo mv nfetch /usr/bin
```
