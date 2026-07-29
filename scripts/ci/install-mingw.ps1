# Install a native MinGW gcc -E toolchain for Windows CI.
# Usage: install-mingw.ps1 -Goarch amd64|arm64
param(
    [Parameter(Mandatory = $true)]
    [ValidateSet("amd64", "arm64", "386")]
    [string]$Goarch
)

$ErrorActionPreference = "Stop"

function Add-PathAndEnv([string]$Bin, [string]$Root, [string]$Gcc) {
    $Bin | Out-File -FilePath $env:GITHUB_PATH -Encoding utf8 -Append
    "MINGW_SYSROOT=$Root" | Out-File -FilePath $env:GITHUB_ENV -Encoding utf8 -Append
    "CC=$Gcc" | Out-File -FilePath $env:GITHUB_ENV -Encoding utf8 -Append
    Write-Host "Using MinGW gcc=$Gcc sysroot=$Root"
    & $Gcc --version
}

function Find-ExistingMingw([string]$Goarch) {
    $triples = @{
        "amd64" = "x86_64-w64-mingw32"
        "arm64" = "aarch64-w64-mingw32"
        "386"   = "i686-w64-mingw32"
    }
    $triple = $triples[$Goarch]
    $candidates = @()
    if ($Goarch -eq "amd64") {
        $candidates += @(
            "C:\ProgramData\mingw64\mingw64",
            "C:\mingw64",
            "C:\ProgramData\chocolatey\lib\mingw\tools\install\mingw64",
            "C:\msys64\mingw64"
        )
    } elseif ($Goarch -eq "arm64") {
        $candidates += @(
            "C:\llvm-mingw",
            "C:\msys64\clangarm64"
        )
    } else {
        $candidates += @(
            "C:\msys64\mingw32",
            "C:\mingw32"
        )
    }

    foreach ($root in $candidates) {
        foreach ($gccName in @("$triple-gcc.exe", "gcc.exe")) {
            $bin = Join-Path $root "bin"
            $gcc = Join-Path $bin $gccName
            if (Test-Path $gcc) {
                return @{ Root = $root; Bin = $bin; Gcc = $gcc }
            }
        }
    }
    return $null
}

$found = Find-ExistingMingw -Goarch $Goarch
if ($null -ne $found) {
    Add-PathAndEnv -Bin $found.Bin -Root $found.Root -Gcc $found.Gcc
    exit 0
}

if ($Goarch -eq "amd64") {
    choco install mingw -y --no-progress
    $found = Find-ExistingMingw -Goarch $Goarch
    if ($null -eq $found) { throw "MinGW gcc.exe not found after choco install" }
    Add-PathAndEnv -Bin $found.Bin -Root $found.Root -Gcc $found.Gcc
    exit 0
}

if ($Goarch -ne "arm64") {
    throw "No automated MinGW installer for goarch=$Goarch"
}

# llvm-mingw ships aarch64-w64-mingw32-gcc for windows/arm64 hosts.
# Pin to LLVM 21: modernc.org/cc fails parsing clang 22+ predefines
# (<predefined>: unexpected integer constant, expected direct declarator).
$version = "20251216"
$asset = "llvm-mingw-$version-ucrt-aarch64.zip"
$uri = "https://github.com/mstorsjo/llvm-mingw/releases/download/$version/$asset"
$zip = Join-Path $env:RUNNER_TEMP $asset
$dest = "C:\llvm-mingw"
Write-Host "Downloading $uri"
Invoke-WebRequest -Uri $uri -OutFile $zip
if (Test-Path $dest) { Remove-Item -Recurse -Force $dest }
Expand-Archive -Path $zip -DestinationPath $env:RUNNER_TEMP -Force
$extracted = Get-ChildItem -Path $env:RUNNER_TEMP -Directory |
    Where-Object { $_.Name -like "llvm-mingw-*-ucrt-aarch64" } |
    Select-Object -First 1
if ($null -eq $extracted) { throw "llvm-mingw extract dir not found" }
Move-Item -Path $extracted.FullName -Destination $dest

$found = Find-ExistingMingw -Goarch $Goarch
if ($null -eq $found) { throw "aarch64 MinGW gcc not found under $dest" }
Add-PathAndEnv -Bin $found.Bin -Root $found.Root -Gcc $found.Gcc
