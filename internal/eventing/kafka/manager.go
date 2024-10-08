package kafka

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"strings"

	confluent "github.com/cloudevents/sdk-go/protocol/kafka_confluent/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cecontext "github.com/cloudevents/sdk-go/v2/context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	authnapi "github.com/project-kessel/inventory-api/internal/authn/api"
	"github.com/project-kessel/inventory-api/internal/eventing/api"
)

type KafkaManager struct {
	Config   CompletedConfig
	Source   string
	Protocol *confluent.Protocol
	Client   cloudevents.Client
	Errors   <-chan error

	Logger *log.Helper
}

func New(config CompletedConfig, source string, logger *log.Helper) (*KafkaManager, error) {
	logger.Info("Using eventing: kafka")
	if sender, err := confluent.New(
		confluent.WithSenderTopic(config.DefaultTopic),
		confluent.WithConfigMap(config.KafkaConfig),
	); err != nil {
		return nil, err
	} else {
		client, err := cloudevents.NewClient(sender, cloudevents.WithUUIDs())
		if err != nil {
			return nil, err
		}

		errChan := make(chan error)

		go func() {
			eventChan, err := sender.Events()
			if err != nil {
				logger.Errorf("failed to get events channel for sender, %v", err)
				errChan <- err
			} else {
				for e := range eventChan {
					switch ev := e.(type) {
					case *kafka.Message:
						// The message delivery report, indicating success or permanent failure after retries have
						// been exhausted. Application level retries won't help since the client is already
						// configured to do that.
						m := ev
						if m.TopicPartition.Error != nil {
							logger.Errorf("Delivery failed: %v\n", m.TopicPartition.Error)
							errChan <- err
						} else {
							logger.Infof("Delivered message to topic %s [%d] at offset %v\n",
								*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
						}
					case kafka.Error:
						e := ev
						if e.IsFatal() {
							logger.Errorf("Error: %v\n", ev)
							errChan <- e
						} else {
							logger.Infof("Error: %v\n", ev)
						}
					default:
						logger.Infof("Ignored event: %v\n", ev)
					}
				}
			}
		}()

		return &KafkaManager{
			Config:   config,
			Source:   source,
			Protocol: sender,
			Client:   client,
			Errors:   errChan,

			Logger: logger,
		}, nil
	}
}

func (m *KafkaManager) Errs() <-chan error {
	return m.Errors
}

// Lookup figures out which topic should be used for the given identity and resource.
func (m *KafkaManager) Lookup(identity *authnapi.Identity, resource_type string, resource_id int64) (api.Producer, error) {

	// there is no complicated topic dispatch logic... for now.
	producer, err := NewProducer(m, m.Config.DefaultTopic, identity)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func (m *KafkaManager) Shutdown(ctx context.Context) error {
	return m.Protocol.Close(ctx)
}

type kafkaProducer struct {
	Manager  *KafkaManager
	Topic    string
	Identity *authnapi.Identity

	Logger        *log.Helper
	eventsCounter metric.Int64Counter
}

// NewProducerEventsCounter creates a meter for capturing event metrics
func NewProducerEventsCounter(meter metric.Meter, histogramName string) (metric.Int64Counter, error) {
	return meter.Int64Counter(histogramName, metric.WithUnit("{event_type}"))
}

// NewProducer produces a kafka producer that is bound to a particular topic.
func NewProducer(manager *KafkaManager, topic string, identity *authnapi.Identity) (*kafkaProducer, error) {
	meter := otel.Meter("github.com/project-kessel/inventory-api/blob/main/internal/server/otel")
	eventsCounter, err := NewProducerEventsCounter(meter, "kafka_event")
	if err != nil {
		return nil, err
	}
	return &kafkaProducer{
		Manager:  manager,
		Topic:    topic,
		Identity: identity,

		Logger:        manager.Logger,
		eventsCounter: eventsCounter,
	}, nil
}

// Produce creates the cloud event and sends it on the Kafka Topic
func (p *kafkaProducer) Produce(ctx context.Context, event *api.Event) error {
	e := cloudevents.NewEvent()

	eventId, err := uuid.NewUUID() // Todo: we need to have an stable id if we implement some re-trying logic
	if err != nil {
		return err
	}

	e.SetSpecVersion(cloudevents.VersionV1)
	e.SetType(makeCloudEventType(event.EventType, event.ResourceType, event.OperationType))
	e.SetSource(p.Manager.Source)
	e.SetID(eventId.String())
	e.SetTime(event.EventTime)

	err = e.SetData(cloudevents.ApplicationJSON, event.Resource)
	if err != nil {
		return err
	}

	e.SetDataContentType("application/json")
	e.SetSubject(makeCloudEventSubject(event.EventType, event.ResourceType, event.ResourceId))

	ret := p.Manager.Client.Send(confluent.WithMessageKey(cecontext.WithTopic(cloudevents.WithEncodingStructured(ctx), p.Topic), p.Manager.Source), e)
	if cloudevents.IsUndelivered(ret) {
		p.Logger.Infof("Failed to send %v", ret)
	} else {
		p.Logger.Infof("Kafka returned: %v", ret)
	}

	p.eventsCounter.Add(
		context.Background(),
		1,
		metric.WithAttributes(
			attribute.String("event_type", event.EventType),
			attribute.String("resource_type", event.ResourceType),
		),
	)
	return ret
}

func makeCloudEventType(eventType, resourceType, operation string) string {
	return fmt.Sprintf("redhat.inventory.%s.%s.%s", eventType, resourceType, operation)
}

func makeCloudEventSubject(eventType, resourceType, resourceId string) string {
	return "/" + strings.Join([]string{eventType, resourceType, resourceId}, "/")
}
