/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/telebot.v3"

	"github.com/spf13/cobra"

	"github.com/hirosassa/zerodriver"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var (
	TeleToken   = os.Getenv("TELE_TOKEN")
	MetricsHost = os.Getenv("METRICS_HOST")

	logger = zerodriver.NewProductionLogger()
)

func initMetrics(ctx context.Context) {

	if len(os.Getenv("METRICS_HOST")) == 0 {
		logger.Info().Str("Version", appVersion).Msg("No METRICS_HOST var provided, metrics are disabled")
		return
	}

	// Create a new OTLP Metric gRPC exporter with the specified endpoint and options
	exporter, _ := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpoint(MetricsHost),
		otlpmetricgrpc.WithInsecure(),
	)

	// Define the resource with attributes that are common to all metrics.
	// labels/tags/resources that are common to all metrics.
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(fmt.Sprintf("kbot_%s", appVersion)),
	)

	// Create a new MeterProvider with the specified resource and reader
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(resource),
		sdkmetric.WithReader(
			// collects and exports metric data every 10 seconds.
			sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(10*time.Second)),
		),
	)

	// Set the global MeterProvider to the newly created MeterProvider
	otel.SetMeterProvider(mp)

}

func sendMetrics(ctx context.Context, payload string) {

	if len(MetricsHost) <= 0 {
		return
	}
	//  Get the global MeterProvider and create a new Meter
	meter := otel.GetMeterProvider().Meter("kbot_commands")

	// Get or create an Int64Counter instrument
	counter, _ := meter.Int64Counter(fmt.Sprintf("kbot_command_%s", payload))

	// Add a value of 1 to the Int64Counter
	counter.Add(ctx, 1)
}

// prometheusKbotCmd represents the prometheusKbot command
var prometheusKbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "starts the application",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {

		// fmt.Printf("kbot %s started", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			logger.Fatal().Str("Error", err.Error()).Msg("Please set correct TELE_TOKEN")
			return
		} else {
			logger.Info().Str("Version", appVersion).Msg("kbot started")
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Println(m.Message().Payload, m.Text())

			payload := m.Message().Payload
			sendMetrics(context.Background(), payload)

			switch payload {

			case "hello":
				err = m.Send(fmt.Sprintf("Hello I'm kbot %s!", appVersion))
				if err != nil {
					log.Fatalf("ERROR: can't sent message. %s", err)
					return err
				}
			case "bye":
				err = m.Send(fmt.Sprintf("Sadly, bye"))
				if err != nil {
					log.Fatalf("ERROR: can't sent message. %s", err)
					return err
				}
			case "version":
				err = m.Send(fmt.Sprintf("Hello I'm Kbot %s!", appVersion))

			default:
				err = m.Send("Usage: /s hello|bye")
			}

			return err

		})

		kbot.Start()
	},
}

func init() {

	ctx := context.Background()
	initMetrics(ctx)

	rootCmd.AddCommand(prometheusKbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prometheusKbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prometheusKbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
