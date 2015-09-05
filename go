#! /bin/bash

# Get directory of this script
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}"  )" && pwd  )

# Remove subl function in ~/.zshrc
sed '/subl/,/}/d' ~/.zshrc > ~/tmp.zshrc
mv ~/tmp.zshrc ~/.zshrc

# Set execute permissions and copy subl script onto path
chmod u+x $DIR/subl
cp $DIR/subl /usr/local/bin

# Change .gitconfig editor to subl
sed 's/nano/subl' ~/.gitconfig > ~/tmp.gitconfig && mv ~/tmp.gitconfig ~/.gitconfig

