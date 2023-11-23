{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/21.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = nixpkgs.legacyPackages.${system};
    in {
      packages.website = pkgs.stdenv.mkDerivation {
        name = "integratn-io";
        src = self;
        buildPhase = "${pkgs.hugo}/bin/hugo";
        installPhase = "cp -r * $out";
      };
      devShell = pkgs.mkShell {
        packages = with pkgs; [
          hugo
        ];
      };
    });
}
