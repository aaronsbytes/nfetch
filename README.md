# ⚡ nFetch

<img src="./screenshot.png" />

⚡ nFetch is a dependency-free, fast system information fetching tool aiming for maximum customizability on peek performance

### 🎨 Customization
Config location is `~/.config/nfetch/nfetch.conf`

See `sample.nfetch.conf` for an example!

Use `<tags>` to insert infos, colors and styles!

### Tags
Infos<br/>
`<os>`
`<distro>`
`<kernel>`
`<shell>`
`<wm>`
`<user>`
`<host>`
`<cpu>`
`<memory>`
`<disk>`
`<packages>`
`<flatpaks>`
`<time>`
`<uptime>`

Foreground<br/>
`<fg-red>`
`<fg-green>`
`<fg-blue>`
`<fg-yellow>`
`<fg-magenta>`
`<fg-cyan>`
`<fg-white>`
`<fg-black>`

Background<br/>
`<bg-red>`
`<bg-green>`
`<bg-blue>`
`<bg-yellow>`
`<bg-magenta>`
`<bg-cyan>`
`<bg-white>`
`<bg-black>`

Styles<br/>
`<bold>`
`<underline>`
`<blink>`
`<inverse>`
`<reset>`



### 💭 Why?
I wanted a system fetching tool for the look. Sadly, the ones i saw where not what i wanted and neofetch is too slow for my liking. This is where Go and boredom came into place.

### 💪 Goal
My goal was to not use any dependencies or shell commands. I wanted to achieve maximum performance while fetching accurate informations. My next goal is to publish it to some package repositories to make installation easier across different linux distros.

### 📝 Recently implemented
- ✅ Config File
- ✅ WM
- ✅ SHELL
- ✅ DISK
- ✅ PACKAGES (dnf, apt, pacman, flatpak)

### ✏️ Todo
- GPU
- RESOLUTION
- THEME
- ICONS
- TERMINAL FONT
- TERMINAL

### 💙 Like it?
If you like this project and want to support me, please leave me a ⭐ on this repo. I would really appreciate it ❤️

### 👥 Contribution
If you have a feature suggestion, improvement or fix, feel free to open an issue or pr for it.

### ⚙️ Build
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
