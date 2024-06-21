
{
  description = "A go dev Env for my development";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        defaultPackage = pkgs.mkShell{

            buildInputs = with pkgs; 
            [
              go 
              golangci-lint
              gofumpt 
            ];
            shellHook = ''
            zsh
            go version 
            echo "Entering Go dev env"
            '';
          };
      }
    );
}
