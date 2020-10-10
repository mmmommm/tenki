# tenki  
  
## About  
  
現在の天気と最高気温、最低気温などの気象情報  
明日の天気と最高気温、最低気温などの気象情報を返すCLIです。  
  
## Install  
  
`brew tap mmmommm/tenki`  
or  
`brew install tenki`  
  
## Usage  
  
現在の情報を返す  
  
`tenki -c <県名> or tenki --curent <県名>`  
> tenki -c tokyo | tenki --curent tokyo  
  
明日の情報を朝・昼・晩に分けて返す  
  
`tenki -n <県名> or tenki --next <県名>`  
> tenki -n osaka | tenki --next osaka  
  
[県名一覧](./prefecture.md)