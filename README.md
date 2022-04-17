Teleport
--
This script will help you to add an Alias for every directory in your Shell.
Currently it just supports Linux ( Debian Based Distributions ). 

With Teleport you can set a name for each directory you want and easily access that through shell, you can set name for nested folders and access them by entering a simple word

## How to use

1. In order not to break your current profile, the script creates a new file named: `.bash_aliases ` in you home directory, you need to add below lines at the end of you `.bashrc` file in `/home/username/.bashrc` ( username should be edited to yours )
```
if [ -f ~/.bash_aliases ]; then
    . ~/.bash_aliases
fi
```
2. Place built package (`teleport`) in `/usr/local/bin` or in `/usr/bin/` you also can make a shortcut with `ln -s ` to make it easier :
   ```
    ln -s <path to teleport built file> /usr/local/bin/sv
    ``` 
    which `sv` will be a short name for teleport

3. Head to each directory you intend to and run `sv alias` which `sv` is the name of teleport built app you previously made a shortcut of and alias is the name you want to assign to that direction.
4. From now on you can jump to the directory easily by write your alias name.