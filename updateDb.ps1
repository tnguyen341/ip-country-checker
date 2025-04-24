$LicenseKey = "YOUR_LICENSE_KEY"

# Set the database type
$DBType = "GeoLite2-Country"

# Download URL
$URL = "https://download.maxmind.com/app/geoip_download?edition_id=$DBType&license_key=$LicenseKey&suffix=tar.gz"

# Output paths
$DownloadPath = "GeoLite2.tar.gz"
$ExtractPath = "GeoLite2"
$FinalDBPath = "GeoLite2-Country.mmdb"

# Download the file
Invoke-WebRequest -Uri $URL -OutFile $DownloadPath

# Extract the tar.gz
tar -xzf $DownloadPath -C .

# Find and move the .mmdb file
$mmdbFile = Get-ChildItem -Recurse -Filter *.mmdb | Select-Object -First 1
Move-Item -Path $mmdbFile.FullName -Destination $FinalDBPath

# Cleanup
Remove-Item $DownloadPath

Remove-Item -Recurse -Force $ExtractPath

Write-Host "✅ Downloaded and extracted $FinalDBPath"
