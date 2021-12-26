#! /bin/sh

echo "Initialising f1-cli installer"

#Check if directory exists, if it does remove it
if [ -d $HOME/.f1 ]
then
  rm -r $HOME/.f1
fi

#Creating a temporary directory
TMP_DIR=$(mktemp -d)
CURR_DIR=$(pwd)

#Creating a hidden directory at $HOME
mkdir $HOME/.f1

if [ $? -eq 0 ]
then
  echo "Downloading..."

  #Getting tarball link using GitHub api
  #Downloading tar file using wget
  link=$(curl -s https://api.github.com/repos/MrVSiK/f1-cli/releases/latest | grep "tarball_url.*" | cut -d : -f 2,3 | tr -d \" | tr -d , )
  wget $link -O $TMP_DIR/tempfile

  echo "Installing..."

  #Extracting tar file into f1_source
  mkdir $TMP_DIR/f1_source && tar -xf $TMP_DIR/tempfile -C $TMP_DIR/f1_source --strip-components 1

  #Generating binary file
  cd $TMP_DIR/f1_source && go build -o ./build/f1

  #Moving binary file into .f1
  mkdir $HOME/.f1/bin && mv $TMP_DIR/f1_source/build/f1 $HOME/.f1/bin/

  export PATH="$PATH:$HOME/.f1/bin"

  echo "Installation Successful"
  echo "Add the following line to ~/.bashrc or ~/.bash_profile"
  echo "PATH=\"\$PATH:\$HOME/.f1/bin\""

  #Cleaning up
  rm -r $TMP_DIR
  cd $CURR_DIR
else
  echo "Error"
fi
