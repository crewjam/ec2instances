package ec2instances

//go:generate sh -c "(echo 'package ec2instances'; echo ''; echo 'const data = `'; curl -sSL http://www.ec2instances.info/instances.json; echo '`') > data.go"
