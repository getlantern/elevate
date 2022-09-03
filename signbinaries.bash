#!/usr/bin/env bash

###############################################################################
#
# This script regenerates the source files that embed the platform-specific
# executables.
#
###############################################################################

set -euo pipefail

function die() {
  echo "$*"
  exit 1
}

if [ -z "$BNS_CERT" ] || [ -z "$BNS_CERT_PASS" ]
then
	die "$0: Please set BNS_CERT and BNS_CERT_PASS to the bns_cert.p12 signing key and the password for that key"
fi

osslsigncode sign -pkcs12 "$BNS_CERT" -pass "$BNS_CERT_PASS" -in binaries/windows_unsigned/elevate_unsigned.exe -out binaries/windows/elevate.exe || die "Could not sign windows"
#codesign -s "Developer ID Application: Brave New Software Project, Inc" -f binaries/osx/cocoasudo || die "Could not sign macintosh"
codesign --options runtime --strict --timestamp --force --deep -r="designated => anchor trusted and identifier com.getlantern.lantern" -s "Developer ID Application: Innovate Labs LLC (4FYC28AXA2)" -v binaries/osx/cocoasudo
