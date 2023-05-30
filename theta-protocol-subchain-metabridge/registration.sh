pwd

cd $GOPATH/src/github.com/thetatoken/theta-metachain-guide

cd sdk/js


#create governance token on mainchain
govTokenAddr=$(node deployGovToken.js privatenet "Subchain 360777 Gov" GOV360777 0x2E833968E5bB786Ae419c4d13189fB081Cc43bab 0x2E833968E5bB786Ae419c4d13189fB081Cc43bab ~/.thetacli/keys/encrypted/2E833968E5bB786Ae419c4d13189fB081Cc43bab qwertyuiop | grep Address | awk '{print $3}')
sed -i "s/\(govTokenContractAddr\s*:\s*\).*/\1\"$govTokenAddr\",/" configs.js


# mint mock wrapped theta
node mintMockWrappedTheta.js privatenet 0x2E833968E5bB786Ae419c4d13189fB081Cc43bab 50000000000000000000000 ~/.thetacli/keys/encrypted/2E833968E5bB786Ae419c4d13189fB081Cc43bab qwertyuiop
sleep 3

node registerSubchain.js privatenet 0x9fbd08fc250bdf051e5a031457ce8225a7d511bf6c31cbc521d1c8c2e37323dd ~/.thetacli/keys/encrypted/2E833968E5bB786Ae419c4d13189fB081Cc43bab qwertyuiop

# stake from validator
node depositStake.js privatenet 100000000000000000000000 0x2E833968E5bB786Ae419c4d13189fB081Cc43bab ~/.thetacli/keys/encrypted/2E833968E5bB786Ae419c4d13189fB081Cc43bab qwertyuiop
cd /app/src/github.com/thetatoken/theta-subchain