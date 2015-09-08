#! /bin/bash

# Copy subl script onto path
curl https://raw.githubusercontent.com/joshwyatt/subl/master/subl > /usr/bin/subl
chmod +x /usr/bin/subl

# Remove subl function in ~/.zshrc
sed '/subl/,/}/d' ~/.zshrc > ~/tmp.zshrc
mv ~/tmp.zshrc ~/.zshrc

# Change .gitconfig editor to subl
sed 's/nano/subl/' ~/.gitconfig > ~/tmp.gitconfig && mv ~/tmp.gitconfig ~/.gitconfig

