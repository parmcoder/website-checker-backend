module github.com/parmcoder/website-checker-backend

replace github.com/parmcoder/website-checker-backend/commands => ../commands

go 1.18

require (
	github.com/joho/godotenv v1.4.0
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.5.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
)
