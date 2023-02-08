{ pkgs  ? import <nixpkgs> {} }:

with pkgs;

mkshell {
  buildInputs = [
    hugo
  ];


}
