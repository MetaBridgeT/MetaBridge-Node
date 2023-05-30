import requests 

res = requests.get("https://raw.githubusercontent.com/thetatoken/theta-metachain-guide/master/sdk/contracts/TFuelTokenBank.json")
print(res.json()["source"])

with open("ab.sol", "w") as f:
    f.write(res.json()["source"])
    