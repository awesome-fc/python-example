# kafka-producer-fc-event-nodejs14 Help Documentation

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=kafka-producer-fc-event-nodejs14&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=kafka-producer-fc-event-nodejs14" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=kafka-producer-fc-event-nodejs14&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=kafka-producer-fc-event-nodejs14" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=kafka-producer-fc-event-nodejs14&type=packageDownload">
  </a>
</p>



## Preliminary preparation

### Permission preparation

Using this item, verify that your operational account has the following product permissions/policies:


| Service/Business     | Functional Computing                                         |
| -------------------- | ------------------------------------------------------------ |
| Permissions/Policies | AliyunFCFullAccess<br/>AliyunKafkaFullAccess<br/>AliyunVPCReadOnlyAccess |


### Resource preparation

  * For an available Kafka message queue, please refer to the official document of message queue Kafka version [Quick Start of Message Queue](https://help.aliyun.com/document_detail/99949.html).

    - Create a VPC private network (VPC is recommended in the production environment), please refer to [VPC official document](https://help.aliyun.com/document_detail/65398.htm?spm=a2c4g.11186623.0.0.61be4c9d4aGfpg# task-1012575). VPC console [link](https://vpcnext.console.aliyun.com/). At this point, you can have a VPC and corresponding switches.

    > When deploying a Kafka instance, you will be prompted to create an available VPC private network

  * Create the Kafka Topic and Consumer Group to be used in the Kafka console.

# Code & Preview

- [ :smiley_cat:source code](https://github.com/devsapp/)
- In order to successfully deploy this sample code, you need to provide the following parameters during the deployment process:
  - Region: You need to configure the region where your Function Compute service needs to be deployed through this parameter. The default value is cn-hangzhou (Hangzhou).
    - The geographic options available to you are:
      - cn-beijing (Beijing)
      - cn-hangzhou (Hangzhou)
      - cn-shanghai (Shanghai)
      - cn-qingdao (Qingdao)
      - cn-zhangjiakou (Zhangjiakou)
      - cn-huhehaote (Hohhot)
      - cn-shenzhen (Shenzhen)
      - cn-chengdu (Chengdu)
      - cn-hongkong (Hong Kong)
      - ap-southeast-1 (Singapore)
      - ap-southeast-2 (Sydney)
      - ap-southeast-3 (Kuala Lumpur)
      - ap-southeast-5 (Jakarta)
      - ap-northeast-1 (Tokyo)
      - eu-central-1 (Frankfurt)
      - eu-west-1 (London)
      - us-west-1 (Silicon Valley)
      - us-east-1 (Virginia)
      - ap-south-1 (Mumbai)
  - Service name: You need to name your Function Compute service. The service name can only contain letters, numbers, underscores and dashes. Cannot start with a number or a dash. The length is between 1-128, the default value is kafka-producer-quick-start.
  - Function name: You need to name your function calculation function. The function name can only contain letters, numbers, underscores and dashes. Cannot start with a number or a dash. The length is between 1-64. The default is kafka-producer-event-function-nodejs14.
  - vpcId: We recommend that you use a VPC to access Kafka, and select the VPC used to create a Kafka instance. Note that az needs to be supported in Function Compute.
  - vswitchIds: Use the vswitch id in the vpc to access kafka on the intranet. Note that az needs to be supported by Function Compute.
  - securityGroupId: The security group id of the vpc where the kafka instance is located, which can be found in the `ECS` console `Network and Security` menu item.
  - Kafka access point (bootstrapServers): the default access point for the Kafka instance you purchased, which can be found in the instance details
  - topicName: The topic name of the Kafka instance. The data production of this topic will trigger the deployment function, which needs to be created in advance.

</codepre>

<deploy>

## Deployment & Experience

<appcenter>

- :fire: via [Serverless Application Center](https://fcnext.console.aliyun.com/applications/create?template=kafka-producer-fc-event-nodejs14) ,
   [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun. com/applications/create?template=kafka-producer-fc-event-nodejs14) the application.

</appcenter>

- Deploy via [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install):

  - [Install Serverless Devs Cli Developer Tools](https://www.serverless-devs.com/serverless-devs/install), and perform [Authorization Information Configuration](https://www.serverless-devs.com/ fc/config);
  - Initialize the project: `s init kafka-producer-fc-event-nodejs14 -d kafka-producer-fc-event-nodejs14`
  - Fill in the parameters described in the above modules
  - Go to the project directory, `cd kafka-producer-fc-event-nodejs14`
  - Do project deployment: `s deploy -y`
- local debugging
  - Enter the application project project and execute the following command: `s invoke -e '{"Key": "test nodejs14 kafka producer"}'`.
  - You can view the result after the simulated event triggers the function.

```bash
========= FC invoke Logs begin =========
c-62f075a3-58d34b2b8b444083bb972022-08-08 10:34:13FC Initialize Start RequestId: c47410a0-ada0-45d8-863f-a9343feaa47e
c-62f075a3-58d34b2b8b444083bb972022-08-08 10:34:13load code for handler:index.initialize
c-62f075a3-58d34b2b8b444083bb972022-08-08 10:34:132022-08-08 10:34:13 c47410a0-ada0-45d8-863f-a9343feaa47e [verbose] Servers:  alikafka-pre-cn-7mz2sr1xa00c-1-vpc.alikafka.aliyuncs.com:9092
c-62f075a3-58d34b2b8b444083bb972022-08-08 10:34:132022-08-08 10:34:13 c47410a0-ada0-45d8-863f-a9343feaa47e [verbose] TopicName:  HelloTopic
c-62f075a3-58d34b2b8b444083bb972022-08-08 10:34:13FC Initialize End RequestId: c47410a0-ada0-45d8-863f-a9343feaa47e
c-62f075a3-58d34b2b8b444083bb972022-08-08 10:34:13FC Invoke Start RequestId: c47410a0-ada0-45d8-863f-a9343feaa47e
c-62f075a3-58d34b2b8b444083bb972022-08-08 10:34:13load code for handler:index.handler
c-62f075a3-58d34b2b8b444083bb972022-08-08 10:34:43FC Invoke End RequestId: 1c233449-024d-4a67-8e7f-83fe3bab6bac

Duration: 10006.53 ms, Billed Duration: 10007 ms, Memory Size: 128 MB, Max Memory Used: 52.28 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62f075a3-58d34b2b8b444083bb97

FC Invoke Result:
Finish sending the message:{"Key": "test nodejs14 kafka producer"}


End of method: invoke
```

- End-to-end testing

  - Log in to the Function Compute console and click `Test Function`

  The returned results are as follows:

  ```bash
  Finish sending the message:{
      "key1": "value1",
      "key2": "value2",
      "key3": "value3"
  }
  ````

  The log is as follows:

  ```bash
  c-62f075a3-58d34b2b8b444083bb972022-08-08 10:32:13FC Initialize Start RequestId: c47410a0-ada0-45d8-863f-a9343feaa47e
  c-62f075a3-58d34b2b8b444083bb972022-08-08 10:32:13 load code for handler:index.initialize
  c-62f075a3-58d34b2b8b444083bb972022-08-08 10:32:132022-08-08 10:32:13 c47410a0-ada0-45d8-863f-a9343fepcaa47e [verbose] Servers: alikafka-pre-cn-1-7mz2srikafkafka0-ada0-45d8-863f-a9343fepcaa47e .aliyuncs.com:9092
  c-62f075a3-58d34b2b8b444083bb972022-08-08 10:32:132022-08-08 10:32:13 c47410a0-ada0-45d8-863f-a9343feaa47e [verbose] TopicName: HelloTopic
  c-62f075a3-58d34b2b8b444083bb972022-08-08 10:32:13FC Initialize End RequestId: c47410a0-ada0-45d8-863f-a9343feaa47e
  c-62f075a3-58d34b2b8b444083bb972022-08-08 10:32:13FC Invoke Start RequestId: c47410a0-ada0-45d8-863f-a9343feaa47e
  c-62f075a3-58d34b2b8b444083bb972022-08-08 10:32:13 load code for handler:index.handler
  c-62f075a3-58d34b2b8b444083bb972022-08-08 10:32:13FC Invoke End RequestId: c47410a0-ada0-45d8-863f-a9343feaa47e
  ````

  - Log in to the Kafka console, view the corresponding topic `message details` of the corresponding instance, find the corresponding partition, and then query the corresponding message.



</deploy>

<appdetail id="flushContent">

# application details



This application is only used for learning and reference. You can carry out secondary development and improvement based on this project to realize your own business logic.



</appdetail>

<devgroup>

## Developer Community

If you have feedback about bugs or future expectations, you can give feedback and exchange in [Serverless Devs repo Issues](https://github.com/serverless-devs/serverless-devs/issues). If you want to join our discussion group or keep up to date with the latest developments in FC components, you can do so through the following channels:

<p align="center">




| <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407298906_20211028074819117230.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407044136_20211028074404326599.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407252200_20211028074732517533.png" width="130px" > |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| <center>WeChat public account:`serverless`</center>          | <center>WeChat assistant: `xiaojiangwh`</center>             | <center>Dingding exchange group:`33947367`</center>          |

</p>

</devgroup>