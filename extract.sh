gow () {
  while true; do
    clear
    echo "go $@"
    go "$@"
    if [ $? -eq 0 ]; then break; fi
    inotifywait -e modify *.go 2> /dev/null
  done
}

gdbb () {

  # build with debug flags
  go build -gcflags "-N -l" -o out

  # make sure the build didn't fail
  if [ $? != 0 ]; then return; fi

  # extract debugger comments
  gdbb-extract "*.go" > .breakpoints

  # break on main if no breakpoints were found
  if [ ! -s .breakpoints ]; then echo "break main.main" > .breakpoints; fi

  # launch gdb
  gdb -x .breakpoints -ex run --args out "$@"

  # clean up
  rm .breakpoints out
}
