{pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
  packages = with pkgs; [
    docker-compose
    tilt
    gopls
    go
  ];
}
