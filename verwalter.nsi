Name "verwalter - v0.1.0-alpha"
OutFile "verwalter_INSTALLER-v0.1.0-alpha.exe"
InstallDir $PROGRAMFILES\verwalter
Section "Version Disclaimer"
MessageBox "This is Version 0.1.-alpha. Expect Bugs."
SectionEnd
Page directory
Page instfiles 
SetOutPath $INSTDIR
Section "Copy bin to program files"
