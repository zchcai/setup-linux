#!/bin/bash

if [ ! -f ~/.vimrc ]; then
	cp ./vimrc ~/.vimrc
fi
if ! cmp ~/.vimrc ./vimrc >/dev/null 2>&1
then
	cp ./vimrc ~/.vimrc
fi

if [ ! -f ~/.tmux.conf ]; then
	cp ./tmux.conf ~/.tmux.conf
fi

# install NERDTree
mkdir -p ~/.vim/autoload ~/.vim/bundle
curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim
cd ~/.vim/bundle
git clone https://github.com/scrooloose/nerdtree.git  
