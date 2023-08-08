Param(
    [string] $filter
)

. (Join-Path $PSScriptRoot .. common scripts common.ps1)
. (Join-Path $PSScriptRoot MgmtTestLib.ps1)

$env:TEMP = [System.IO.Path]::GetTempPath()
Write-Host "Path tmp: $env:TEMP"

$sdks = Get-AllPackageInfoFromRepo $filter

Write-Host "Prepare fake server"
if ($sdks.Count -eq 0)
{
    Write-Host "No package need to be test"
    exit 0
}

foreach ($sdk in $sdks)
{
    if ($sdk.SdkType -eq "mgmt")
    {
        try
        {
            Write-Host "Execute fake test for $($sdk.Name)"
            TestAndGenerateReport $sdk.DirectoryPath
        }
        catch
        {
            Write-Host "##[error]can not finish single test for $sdks :`n$_"
            exit 1
        }
    }
}