#!/usr/bin/env bash
set -euo pipefail

GREEN="\033[0;32m"; YELLOW="\033[1;33m"; RED="\033[0;31m"; NC="\033[0m"

if [ $# -lt 1 ]; then
  echo -e "${RED}Usage:${NC} venv-run.sh <command> [args]"
  exit 1
fi

PYTHON_BIN="${PYTHON_BIN:-python3}"
TMP_VENV="$(mktemp -d -t algos-venv-XXXXXX)"

cleanup() { rm -rf "${TMP_VENV}" || true; }
trap cleanup EXIT

echo -e "${YELLOW}Creating virtual environment...${NC}"
"${PYTHON_BIN}" -m venv "${TMP_VENV}"
# shellcheck disable=SC1091
source "${TMP_VENV}/bin/activate"
export PIP_DISABLE_PIP_VERSION_CHECK=1
python -m pip install --upgrade pip >/dev/null
echo -e "${GREEN}✓ Virtual environment ready!${NC}"

CUR_REQ="${PWD}/requirements.txt"
ROOT_REQ="$(cd "$(git rev-parse --show-toplevel 2>/dev/null || echo "${PWD}/../..")" 2>/dev/null && pwd)/requirements.txt"

if [ -f "${ROOT_REQ}" ]; then
  echo -e "${YELLOW}Installing repo requirements...${NC}"
  pip install -r "${ROOT_REQ}"
fi
if [ -f "${CUR_REQ}" ]; then
  echo -e "${YELLOW}Installing task requirements...${NC}"
  pip install -r "${CUR_REQ}"
else
  echo -e "${YELLOW}No requirements.txt found, will auto‑install missing imports if any.${NC}"
fi

map_pkg() {
  case "$1" in
    sklearn) echo "scikit-learn" ;;
    cv2) echo "opencv-python" ;;
    PIL) echo "Pillow" ;;
    bs4) echo "beautifulsoup4" ;;
    yaml) echo "PyYAML" ;;
    *) echo "$1" ;;
  esac
}

run_with_autoinstall() {
  local attempts=0
  local log="$(mktemp)"
  while true; do
    if "$@" > >(tee "${log}") 2> >(tee -a "${log}" >&2); then
      rm -f "${log}"
      return 0
    fi
    if [[ "$1" != "python" && "$1" != "python3" ]]; then
      rm -f "${log}"
      return 1
    fi
    local missing
    missing="$(grep -E "ModuleNotFoundError: No module named '([^']+)'" "${log}" | tail -n1 | sed -E "s/.*ModuleNotFoundError: No module named '([^']+)'.*/\1/")" || true
    if [ -n "${missing:-}" ]; then
      attempts=$((attempts+1))
      if [ $attempts -gt 5 ]; then
        echo -e "${RED}Too many missing modules. Aborting.${NC}"
        rm -f "${log}"
        return 1
      fi
      local pkg
      pkg="$(map_pkg "${missing}")"
      echo -e "${YELLOW}→ Installing missing module '${missing}' as package '${pkg}'...${NC}"
      if ! pip install "${pkg}"; then
        echo -e "${RED}Failed to install ${pkg}${NC}"
        rm -f "${log}"
        return 1
      fi
      continue
    fi
    rm -f "${log}"
    return 1
  done
}

echo -e "${GREEN}→ Running:${NC} $*"
if ! run_with_autoinstall "$@"; then
  echo -e "${RED}Command failed.${NC}"
  exit 1
fi
echo -e "${GREEN}✓ Done. Cleaning up temporary venv.${NC}"