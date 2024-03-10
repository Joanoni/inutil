$versionType = $args[0]
$message = $args[0]

$lastVersion = Get-Content ./lastVersion.txt -Tail 1

$versionSplitted = $lastVersion.Split(".")

$versionNew = ""

switch ($versionType) {
    "patch" {
        "Generating new version: patch";
        $versionNew = $versionSplitted[0] + "." + $versionSplitted[1] + "." + ([int]$versionSplitted[2] + 1)
        $versionNew
        break
    }
    "minor" {
        "Generating new version: minor";
        $versionNew = $versionSplitted[0] + "." + ([int]$versionSplitted[1] + 1) + ".0"
        $versionNew
        break
    }
    "major" {
        "Generating new version: major";
        $versionNew = ($versionSplitted[0] + 1) + ".0.0"
        $versionNew
        break
    }
    default {
        "First parameter is the version type: patch||minor||major"; 
        break
    }
}

"New version: " + $versionNew

git add .
git commit -m $message  --quiet
git tag $versionNew
git push origin $versionNew  --quiet

"Pushed to git"

Add-Content -Path .\lastVersion.txt -Value $versionNew 