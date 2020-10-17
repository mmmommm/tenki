# tenki  
  
## About  
  
現在の天気と最高気温、最低気温などの気象情報  
明日の天気と最高気温、最低気温などの気象情報を返すCLIです。  
  
## Install  
  
`brew tap mmmommm/tenki`  
`brew install tenki`  
  
## Usage  
  
現在の情報を返す。  
  
`tenki -c <県名> or tenki --curent <県名>`  
> tenki -c tokyo | tenki --curent tokyo  
  
![-cコマンド](https://user-images.githubusercontent.com/51479834/96333964-2a96ad00-10a8-11eb-8219-ad6cca82aa46.png)
  
現在から9時間分の天気予報を返す。  
  
`tenki -f <県名> or tenki --forecast <県名>`  
> tenki -f fukuoka | tenki --forecast fukuoka  
  
![-fコマンド](https://user-images.githubusercontent.com/51479834/96333999-6467b380-10a8-11eb-814a-14b0af1586f6.png)
  
明日の9時、12時、18時、21時の天気予報を返す。  
  
`tenki -n <県名> or tenki --next <県名>`  
> tenki -n osaka | tenki --next osaka  
  
![-nコマンド](https://user-images.githubusercontent.com/51479834/96334005-6fbadf00-10a8-11eb-81aa-047497612185.png)
  
[県名一覧](./prefecture.md)