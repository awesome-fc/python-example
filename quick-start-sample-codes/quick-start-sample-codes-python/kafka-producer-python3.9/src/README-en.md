# Python3.9 Kafka message queue producer example

This example shows you the Python [Message Queue Kafka](https://help.aliyun.com/document_detail/68151.html?spm=5176.167616.J_5253785160.5.2dfe6feexRPqMj) producer example. (See Kafka Trigger Example for Consumer Example)

This example uses the [Message Queue Kafka Version Official Python SDK](https://help.aliyun.com/document_detail/159700.html)



 ## ready to start

 - For an available Kafka message queue, please refer to the official document of message queue Kafka version [Quick Start of Message Queue](https://help.aliyun.com/document_detail/99949.html).

   - Create a VPC private network (VPC is recommended in the production environment), please refer to [VPC official document](https://help.aliyun.com/document_detail/65398.htm?spm=a2c4g.11186623.0.0.61be4c9d4aGfpg# task-1012575). VPC console [link](https://vpcnext.console.aliyun.com/). At this point, you can have a VPC and corresponding switches.

   > When deploying a Kafka instance, you will be prompted to create an available VPC private network

 - [Optional] Install and configure Serverless Devs tools. (https://help.aliyun.com/document_detail/195474.html)

 ## quick start

 ### Method 1. Use the console to create

#### 1. Install dependencies and deploy code packages

For Python, it is recommended to use `layer` to deploy dependencies.

In the advanced functions on the left menu bar of the Function Compute console, select `Layer Management`. Select Python3.6 for the compatible runtime, select `Online Build Dependency Layer` for the layer upload method, and copy the **requirements.txt** file into it. Finally, wait for the layer to be created.

![CreateLayer](assets/layer.png)

For the code, just package and compress it in the code directory:

```shell
zip code.zip -r ./*
````



 #### 2. Create service

It is recommended to create services in the same Region of the Kafka instance.

When creating a service, select AliyunFcDefaultRole in `Service Role` in `Advanced Options` (if not, create a corresponding role according to the prompts), and enable `Allow access to VPC`, and select `VPC' and `Switch when creating a Kafka instance. ` and the corresponding `security group (created automatically after the Kafka instance is deployed)`.

![CreateService.png](assets/CreateService.png)



#### 3. Create function

  After creating the service, click Create Function as shown

 - Select `Create from scratch with standard Runtime`
 - Fill in the function name
 - Select the code upload method `Upload code via zip package` to upload the corresponding code zip package
 - Select the operating environment Python3.9
 - Select function trigger method: trigger by event request
 - Use default for other settings

 > For the detailed function creation process, see the document: [Create a function using the console](https://help.aliyun.com/document_detail/51783.html)



#### 4. Configure environment variables, instance lifecycle callbacks and layers

Set environment variables in the `Function Configuration` module in the function details and configure the Initializer callback procedure in the instance lifecycle callback.

Where environment variables:

- Modify region, serviceName, functionName (set the same region as the Kafka instance)
- BOOTSTRAP_SERVERS is set to the `default access point` address corresponding to the `access point information` in the Kafka instance details.
- TOPIC_NAME is set to the topic to which the corresponding message is sent (need to be created in advance in the Kafka message queue version)

Initializer is set to index.initialize;

Layer editor adds previously built dependency layers.

![FunctionConfig.png](assets/FunctionConfig.png)



 #### 5. Test function

 The returned results are as follows:

 ```bash
finish sending message: b'{\n "key1": "value1",\n "key2": "value2",\n "key3": "value3"\n}'
 ````

The log output is as follows:

```bash
FunctionCompute python3 runtime initiated.
FC Initialize Start RequestId: fcfd1459-6c97-42b8-bd14-ab205b6d2f27
FC Initialize End RequestId: fcfd1459-6c97-42b8-bd14-ab205b6d2f27
FC Invoke Start RequestId: 6476cbb0-223b-4763-a13e-15cf2c2cf280
2022-07-31 14:42:20 6476cbb0-223b-4763-a13e-15cf2c2cf280 [INFO] Message delivered to HelloTopic [2]
FC Invoke End RequestId: 6476cbb0-223b-4763-a13e-15cf2c2cf280
````



 ### Method 2. Compile and deploy using Serverless Devs tools

 #### 1. Modify s.yaml configuration

- Modify vpcConfig and fill in the VPC ID, security group ID, and vSwitchID corresponding to the Kafka instance.
- Modify the environmentVariables configuration, fill in BOOTSTRAP_SERVERS and TOPIC_NAME

 #### 2. Install dependencies and deploy

 Install dependent libraries

 ```shell
# Use the s tool to install dependencies, you need to use docker
s build --use-docker
 ````

 deploy code

 ```bash
s deploy -y
 ````

 #### 3. Invoke the test

 ```shell
s invoke -e '{
    "Key": "test python3 serverless devs"
}'
 ````

 The response received when calling the function looks like this:

 ```bash
========= FC invoke Logs begin =========
FunctionCompute python3 runtime inited.
FC Initialize Start RequestId: 27a1f732-9375-432f-84b5-33b46210f712
FC Initialize End RequestId: 27a1f732-9375-432f-84b5-33b46210f712
FC Invoke Start RequestId: 27a1f732-9375-432f-84b5-33b46210f712
2022-07-31T06:46:24.956Z 27a1f732-9375-432f-84b5-33b46210f712 [INFO] Message delivered to HelloTopic [9]
FC Invoke End RequestId: 27a1f732-9375-432f-84b5-33b46210f712

Duration: 34.20 ms, Billed Duration: 35 ms, Memory Size: 128 MB, Max Memory Used: 29.25 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e62540-2b701a7e77e048eebb14

FC Invoke Result:
finish sending message: b'{\n    "Key": "test python3 serverless devs"\n}'


End of method: invoke
 ```

