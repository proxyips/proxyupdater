cd bin
rm proxyupdater.zip
7z a proxyupdater.zip data proxyupdater install.sh
scp proxyupdater.zip root@tlm.pw:/opt/proxyips
ssh -t root@tlm.pw "cd /opt/proxyips ; bash --login"
