go build -o nexus-cli .
curl -u admin:admin123 --upload-file nexus-cli "https://nexus.szistech.com/repository/raw-szis-releases/sziscloud/nexus-cli/1.0.1/nexus-cli"
if [ $? eq 0]; then
  echo 'upload success.'
fi