'use strict';
const Kafka = require('kafkajs');

var BOOTSTRAP_SERVERS
var TOPIC_NAME
var producer
exports.initialize = async (context, callback) => {
    BOOTSTRAP_SERVERS = process.env.BOOTSTRAP_SERVERS
    TOPIC_NAME = process.env.TOPIC_NAME
    console.log("Servers: ", BOOTSTRAP_SERVERS);
    console.log("TopicName: ", TOPIC_NAME);
    var servers = BOOTSTRAP_SERVERS.split(",");
    const kafka = new Kafka.Kafka({
        clientId: 'testId',
        brokers: servers
    })
    producer = kafka.producer()
    await producer.connect()
    
    callback(null,"initialize");
}

exports.handler = async (event, context, callback) => {
  try {
    await producer.send({
      topic: 'TestTopic',
      messages: [
          { value: event },
      ],
    })
    callback(null,"Finish sending the message:" + event);
  } catch (e) {
    console.log(e)
    callback(e,"Send message fail!");
  } 
}

exports.preStop = async function(context, callback) {
    console.log('preStop hook start');
    if (producer != null) {
        await producer.disconnect();
    }
    console.log('preStop hook finish');
    callback(null, "");
}