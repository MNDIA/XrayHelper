English | [简体中文](README_zh_CN.md)

# XrayHelper
XrayHelper for Android, some scripts in [Xray4Magisk](https://github.com/Asterisk4Magisk/Xray4Magisk) rewritten with golang, provide arm64 and amd64 binary.

## Control Core Service
`xrayhelper service start`, start core service  
`xrayhelper service stop`, stop core service  
`xrayhelper service restart`, restart core service  
`xrayhelper service status`, show core status

## Control System Proxy
Support application package proxy list run with blacklist and whitelist, bypass specific network interface, and proxy ap interface, should configure **proxy**
```yaml
proxy:
    method: tproxy
    tproxyPort: 65535
    socksPort: 65534
    enableIPv6: false
    mode: whitelist
    pkgList:
        - com.kiwibrowser.browser
        - com.termux
    apList:
        - wlan2
        - rndis0
    ignoreList:
        - ignore
    intraList:
        - 192.168.123.0/24
```
`xrayhelper proxy enable`, enable system proxy  
`xrayhelper proxy disable`, disable system proxy    
`xrayhelper proxy refresh`, refresh system proxy  

## Update Components
- update core  
  `xrayhelper update core`, should configure **xrayHelper.coreType** first, support xray, sing-box  
- update tun2socks  
  `xrayhelper update tun2socks`, update tun2socks from [heiher/hev-socks5-tunnel](https://github.com/heiher/hev-socks5-tunnel)  
- update geodata  
  `xrayhelper update geodata`, update geodata from [Loyalsoldier/v2ray-rules-dat](https://github.com/Loyalsoldier/v2ray-rules-dat)  
- update subscribe nodes  
  `xrayhelper update subscribe`, update your subscribe, should configure **xrayHelper.subList** first, fully compatible with [v2rayNg](https://github.com/2dust/v2rayNG)'s subscription link standard  

## Switch Proxy Node  
- switch subscribe nodes  
  `xrayhelper switch`, should configure **xrayHelper.proxyTag** and update subscribe first, **warning: it will replace your outbounds configuration which has the same proxy tag**
- switch custom nodes  
  `xrayhelper switch custom`, put custom nodes share link into `${xrayHelper.dataDir}/custom.txt` file, then you can find them use this command

## Credits
- [@Loyalsoldier/v2ray-rules-dat](https://github.com/Loyalsoldier/v2ray-rules-dat)
- [@Asterisk4Magisk/Xray4Magisk](https://github.com/Asterisk4Magisk/Xray4Magisk)
- [@2dust/v2rayNG](https://github.com/2dust/v2rayNG)
- [@heiher/hev-socks5-tunnel](https://github.com/heiher/hev-socks5-tunnel)
