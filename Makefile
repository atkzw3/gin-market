# air パッケージ起動
# デバック時に使用する air 単体でも良いが、叩くコマンドはmakefileで一括管理
# air があれば毎回 go run main.goを叩く必要なく、発火させる必要もない
air:
	air

up:
	docker-compose up -d
