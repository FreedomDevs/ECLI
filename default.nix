{
  stdenv,
  pkgs,
}:
pkgs.buildGoModule {
  pname = "ecli";
  version = "1.0";

  src = ./.;

  GOAMD64 = stdenv.hostPlatform.goAMD64 or "v1";

  vendorHash = "sha256-T8mcyv19XWAE3egX+5J//2gmKm7u39F07o8Qymoo6Rg=";

  ldflags = [
    "-X ecli/cmd.ShellPath=${pkgs.dash}/bin/dash"
    "-X ecli/cmd.ScriptPath=${placeholder "out"}/share/ecli/create_svc_network.sh"
  ];

  nativeBuildInputs = [pkgs.installShellFiles];

  postInstall = ''
    mkdir -p $out/share/ecli
    cp create_svc_network.sh $out/share/ecli/create_svc_network.sh
    chmod +x $out/share/ecli/create_svc_network.sh

    $out/bin/ecli completion bash > ecli.bash
    installShellCompletion --bash ecli.bash

    $out/bin/ecli completion zsh > _ecli
    installShellCompletion --zsh _ecli

    $out/bin/ecli completion fish > ecli.fish
    installShellCompletion --fish ecli.fish
  '';

  meta = with pkgs.lib; {
    description = "Хуйня для генерации шаблонных репо, и другой хуйни для ElysiumSMP";
    homepage = "https://github.com/FreedomDevs/ECLI";
    license = licenses.mit;
    maintainers = [foksik mikinol];
  };
}
