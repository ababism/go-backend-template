package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"hw_3/cmd/app"
	"hw_3/internal/config"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.InitViperConfig()
	logrus.SetFormatter(new(logrus.TextFormatter))

	var cfgPath = "./config/config.yml"
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		//fmt.Println(fmt.Errorf("fatal: init config %w", err))
		//os.Exit(1)
		logrus.Fatalf("fatal: init config %v", err)
	}

	a := app.NewApp(cfg)
	// контекст для того, чтобы поймать сигнал о завершении приложения
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	defer func() {
		v := recover()

		if v != nil {
			ctx, _ := context.WithTimeout(ctx, 5*time.Second)
			a.Stop(ctx)
			//fmt.Println(v)
			//os.Exit(1)
			logrus.Fatalln(v)
		}
	}()

	a.Run()

	<-ctx.Done()
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	a.Stop(ctx)
}
