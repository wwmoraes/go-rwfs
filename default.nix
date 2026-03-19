{
  pkgs,
  ...
}:

pkgs.buildGoApplication rec {
  pname = "go-rwfs-integration";
  version = "0.0.0";

  src =
    with pkgs.lib.fileset;
    toSource {
      root = ./.;
      fileset = unions [
        (fileFilter (file: file.hasExt "go") ./.)
        (maybeMissing ./go.sum)
        ./go.mod
      ];
    };

  modules = ./gomod2nix.toml;
  subPackages = [ "cmd/internal/integration" ];

  CGO_ENABLED = 0;
  GOFLAGS = [
    "-race"
    "-mod=readonly"
    "-trimpath"
  ];

  ldflags = [
    "-s"
    "-w"
    "-buildid="
    "-X main.version=${version}"
  ];

  meta = {
    description = "golang read-write filesystem interfaces";
    homepage = "https://github.com/wwmoraes/go-rwfs";
    license = pkgs.lib.licenses.mit;
    maintainers = [ pkgs.lib.maintainers.wwmoraes or "wwmoraes" ];
    mainProgram = "integration";
  };
}
