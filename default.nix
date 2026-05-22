{ pkgs ? import <nixpkgs> {} }:

pkgs.buildGoModule {
  pname = "ecli";
  version = "1.0";

  src = ./.;

  vendorHash = "sha256-T8mcyv19XWAE3egX+5J//2gmKm7u39F07o8Qymoo6Rg=";

  meta = with pkgs.lib; {
    description = "Хуйня для генерации шаблонных репо, и другой хуйни для ElysiumSMP";
    homepage = "https://github.com/FreedomDevs/ECLI";
    license = licenses.mit;
    maintainers = [ foksik mikinol ];
  };
}
