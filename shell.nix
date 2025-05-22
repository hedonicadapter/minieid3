{pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
  packages = with pkgs;
    [
      docker
      docker-compose
      tilt
      gopls
      go
      azure-cli
    ]
    ++ (pkgs.lib.optionals pkgs.stdenv.isDarwin [
      colima
    ]);

  shellHook =
    if pkgs.stdenv.isDarwin
    then ''
      if ! colima status >/dev/null 2>&1; then
        echo "Starting colima"
        colima start
      else
        echo "Colima is already running"
      fi
    ''
    else "";
}
