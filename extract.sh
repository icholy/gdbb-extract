
# debug application
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

# debug tests
gdbbtest () {

  # build with debug flags
  go test -c -gcflags "-N -l" "$@"

  # make sure the build didn't fail
  if [ $? != 0 ]; then return; fi

  # extract debugger comments
  gdbb-extract "*.go" > .breakpoints

  # if breakpoints were found, run on start
  if [ -s .breakpoints ]; then
    gdb -x .breakpoints -ex run *.test
  else
    gdb *.test
  fi

  # clean up
  rm .breakpoints *.test
}
