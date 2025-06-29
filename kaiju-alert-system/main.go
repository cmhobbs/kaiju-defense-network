package main

import (
	"context"
	"fmt"
	kaijuSightingsGenerator "kaiju-sightings-generator"
	"log"
	"os"
	"strings"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func shouldAlert(threatLevel string) bool {
	return threatLevel == "High" || threatLevel == "Critical"
}

func requiredAction(threatLevel string) string {
	switch threatLevel {
	case "Critical":
		return "EVACUATE IMMEDIATELY"
	case "High":
		return "PREPARE DEFENSES"
	case "Medium":
		return "SHELTER IN PLACE"
	default:
		return "MONITOR SITUATION"
	}
}

func formatAlert(sighting kaijuSightingsGenerator.Sighting) string {
	kaiju := sighting.Kaiju
	threatLevel := strings.ToUpper(kaiju.ThreatLevel)
	action := requiredAction(kaiju.ThreatLevel)

	return fmt.Sprintf("%s - WARNING: A %s LEVEL KAIJU HAS BEEN SPOTTED!  %s.", sighting.Timestamp.Format(time.DateTime), threatLevel, action)
}

func initTracer() (*sdktrace.TracerProvider, error) {
	// Create stdout exporter to be able to retrieve
	// the collected spans.
	exporter, err := otlptrace.New(context.Background(), otlptracehttp.NewClient())
	if err != nil {
		return nil, err
	}

	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, err
}

func processKaijuSighting(ctx context.Context, sighting kaijuSightingsGenerator.Sighting) {
	ctx, span := otel.Tracer("kaiju-alert-system").Start(ctx,
		"processKaijuSighting",
		trace.WithAttributes(
			attribute.String("kaiju.name", sighting.Kaiju.Name),
			attribute.String("kaiju.threat_level", sighting.Kaiju.ThreatLevel),
			attribute.String("kaiju.location", sighting.Kaiju.Location),
		),
	)
	defer span.End()

	kaijuSightingsGenerator.PrintSingleSighting(sighting)
	kaiju := sighting.Kaiju

	if shouldAlert(kaiju.ThreatLevel) {
		logAlert(ctx, sighting)
	} else {
		fmt.Println("Threshold too low, alert will not be logged...")
		fmt.Println()
	}
}

func logAlert(ctx context.Context, sighting kaijuSightingsGenerator.Sighting) {
	ctx, span := otel.Tracer("kaiju-alert-system").Start(ctx,
		"logAlert",
		trace.WithAttributes(
			attribute.String("alert.kaiju_name", sighting.Kaiju.Name),
			attribute.String("alert.threat_level", sighting.Kaiju.ThreatLevel),
		),
	)
	defer span.End()

	fmt.Println("Threshold hit, logging alert...")
	fmt.Println()

	alert := formatAlert(sighting)

	file, err := os.OpenFile("alert.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		span.RecordError(err)
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(alert + "\n")
	if err != nil {
		span.RecordError(err)
		log.Fatal(err)
	}

	span.SetAttributes(attribute.String("alert.message", alert))
}

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer tp.Shutdown(context.Background())

	ctx, span := otel.Tracer("kaiju-alert-system").Start(context.Background(), "main")
	defer span.End()

	sightings := kaijuSightingsGenerator.GenerateMultiple(5)
	span.SetAttributes(attribute.Int("sightings.count", len(sightings)))

	for i, sighting := range sightings {
		ctx, sightingSpan := otel.Tracer("kaiju-alert-system").Start(ctx,
			fmt.Sprintf("sighting_%d", i+1),
			trace.WithAttributes(attribute.Int("sighting.index", i+1)),
		)

		processKaijuSighting(ctx, sighting)
		sightingSpan.End()
	}
}
