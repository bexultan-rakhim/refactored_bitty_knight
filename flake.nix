{
  description = "dev env to build bitty knight";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.11";

  outputs = { self, nixpkgs, ... }:
    let
      goVersion = 22;
      supportedSystems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forEachSupportedSystem = f: nixpkgs.lib.genAttrs supportedSystems (system: f {
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ self.overlays.default ];
        };
      });
    in
    {
      overlays.default = final: prev: {
        go = final."go_1_${toString goVersion}";
      };

      devShells = forEachSupportedSystem ({ pkgs }: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gotools
            golangci-lint
            gopls
            # Raylib Dependencies
            pkg-config
            xorg.libX11
            xorg.libXcursor
            xorg.libXi
            xorg.libXinerama
            xorg.libXrandr
            libGL
            libxkbcommon
            wayland
            alsa-lib
            pipewire
          ];

          # Essential for CGO to find the libraries above
          LD_LIBRARY_PATH = pkgs.lib.makeLibraryPath (with pkgs; [
            libGL
            xorg.libX11
            alsa-lib
            pipewire
          ]);

          shellHook = ''
            export CGO_ENABLED=1
            export SHELL="${pkgs.bashInteractive}/bin/bash"
          '';
        };
      });
    };
}
