{
  pkgs,
  ...
}:
rec {
  default = pkgs.mkShell {
    nativeBuildInputs = [
      # keep-sorted start
      (pkgs.mkGoEnv { pwd = ./.; })
      pkgs.cocogitto
      pkgs.gotestdox
      pkgs.remake
      pkgs.unstable.go
      pkgs.unstable.golangci-lint
      # keep-sorted end
    ];
  };

  ci = default.overrideAttrs (
    final: prev: {
      nativeBuildInputs = [
        # keep-sorted start
        pkgs.bashNonInteractive
        # keep-sorted end
      ]
      ++ prev.nativeBuildInputs;

      shellHook = ''
        export GOCACHE=$(go env GOCACHE)
        export GOMODCACHE=$(go env GOMODCACHE)
      '';
    }
  );

  terminal = default.overrideAttrs (
    final: prev: {
      nativeBuildInputs = [
        # keep-sorted start
        pkgs.bash
        pkgs.gomod2nix
        pkgs.nix-update
        pkgs.unstable.gotests
        pkgs.unstable.gotools
        # keep-sorted end
      ]
      ++ prev.nativeBuildInputs;

      shellHook = ''
        cog install-hook --all --overwrite
      '';
    }
  );
}
