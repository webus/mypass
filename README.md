# MyPass
Personal secure passwords manager with CLI based on [Golang](https://golang.org) and [BoltDB](https://github.com/boltdb/bolt).

### What is it ?

This is a very simple utility with which you can conveniently and safely store your passwords.

### How can i use it ?

- Just download latest stable release or build latest version from source code.
- After the first launch it will create the settings file. You can find it in `~/mypass.cfg`.
- Add your first password in the database. To do this, just type the command `mypass add amazon`. Instead of "Amazon" you can write any word that will be key for the password. After you run the command opens the text editor (the default editor `vi` you can specify another editor in the settings file). Write in the editor the password that you want to remember, and close the editor with save all. That's all !
- To get the password just type the command `mypass get amazon`. In this example just replace `amazon` to your key. 
- You can also add logins to MyPass. Just type `mypass add -l amazon my_login`. In this example just replace `amazon` to your key and `my_login` to your login.

### Features

- Simple command line tool
- Once the password is copied to the clipboard.
- All passwords and logins are stored in encrypted form
- You can choose the storage location of the password database. For example, you can specify ~/Dropbox to store the database in Dropbox :)

### Build from source

- You must have installed [Golang](https://golang.org).
- You must have installed [Gb](https://getgb.io/).
- Install `xsel` and `xclip` by this command: `sudo apt-get install xsel xclip`.
- Install build tools by this command: `sudo apt-get install build-essential`.
- Clone this repository `git clone https://github.com/webus/mypass`.
- Go to repository and build it by this command: `cd mypass && make`.
- Copy `mypass` to your bin folder: `cp bin/mypass ~/bin`
- Have fun! :)
