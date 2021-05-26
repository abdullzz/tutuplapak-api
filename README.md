# tutuplapak-api
tutuplapak api
source tutorial : https://www.youtube.com/watch?v=FLnxgSZ0DG4
tutorial:
1. npm install -g json-server / kalo gabisa => / npm install -g json --registry https://registry.npmjs.org
2. json-server db.json
3. harusnya sampai sini sudah jalan di local

===============================================

untuk deploy di heroku :
4. npm init
5. npm i
6. npm i json-server / kalo gabisa => / npm i json-server --registry https://registry.npmjs.org
7. ubah package.json. tambahkan di dalam bagian key "scripts" dibawah key "test", yaitu "start": "node server.js"
8. install heroku (kalo di mac pake brew => liat aja di web heroku nya)
9. npm install -g heroku / kalo gabisa => / npm -g heroku --registry https://registry.npmjs.org
10. heroku create tutuplapak-api
11. git push heroku main
12. coba buka link heroku dari poin 10
