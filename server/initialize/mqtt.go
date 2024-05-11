package initialize

import (
	"fmt"
	"github.com/CYsiod/GreenHydrogen/server/global"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

//var greenHEnergyService = service.ServiceGroupApp.EquipmentServiceGroup.GreenHEnergyService

func Mqtt() mqtt.Client {
	mq := global.GVA_CONFIG.Mqtt
	mqttOption := &mqtt.ClientOptions{
		KeepAlive:      mq.KeepAlive,
		Username:       mq.UserName,
		Password:       mq.PassWord,
		ClientID:       mq.ClientID,
		ConnectRetry:   true,
		ConnectTimeout: 5 * time.Second,
		AutoReconnect:  true,
	}
	mqttOption.AddBroker(mq.Broker)
	fmt.Println(mqttOption)
	mqttOption.OnConnectionLost = OnConnectionLost
	mqttOption.OnConnect = OnConnect
	mqttOption.SetDefaultPublishHandler(DefaultPublishHandler)

	client := mqtt.NewClient(mqttOption)
	token := client.Connect()
	token.Wait()
	client.Subscribe(mq.Topic, 2, func(client mqtt.Client, message mqtt.Message) {
		global.GVA_LOG.Info(fmt.Sprintf("subscribe [%s] %s\n", message.Topic(), string(message.Payload())))

		//newEquipment := &equipment.GreenHEnergy{}
		//err := json.Unmarshal(message.Payload(), &newEquipment)
		//if err != nil {
		//	global.GVA_LOG.Info("Mqtt message Unmarshal failed")
		//	return
		//}
		//err = greenHEnergyService.CreateGreenHEnergy(newEquipment)
		//if err != nil {
		//	global.GVA_LOG.Info("Mqtt greenHEnergyService.CreateGreenHEnergy failed")
		//}
	})
	if !client.IsConnected() {
		global.GVA_LOG.Panic("Mqtt Is Not Connected")
	}
	return client
}

func OnConnectionLost(client mqtt.Client, err error) {
	fmt.Printf("Mqtt Connect lost: %v", err)
}

func OnConnect(client mqtt.Client) {
	fmt.Println("Mqtt Connected")
}

func DefaultPublishHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Mqtt Received message: %s from topic: %s\n", message.Payload(), message.Topic())
}
