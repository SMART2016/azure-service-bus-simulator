# azure-service-bus-simulator
This is azure service buss simulator with go code
## Running Service bus emulator

- Download the git example Emulator project
    - https://github.com/alex-wolf-ps/service-bus-emulator/tree/main
- Change few configs in the below files
    - change the `.env` and `config.json`
        - `.env changes`
            - Set the `CONFIG_PATH` property to the absolute location of the config.json file
            - Set `ACCEPT_EULA` to `"Y"`
            - Set `MSSQL_SA_PASSWORD` to whatever you wish
        - `config.json` changes
            - Add the namespace you wish to have for your service bus and add the queues needed as below
                ```json
                {  
                        "UserConfig": {  
                         "Namespaces": [  
                           {         "Name": "sbemulatorns",  
                             "Queues": [  
                               {             "Name": "control-plane-notificatio",  
                                 "Properties": {  
                                   "DeadLetteringOnMessageExpiration": false,  
                                   "DefaultMessageTimeToLive": "PT1H",  
                                   "DuplicateDetectionHistoryTimeWindow": "PT20S",  
                                   "ForwardDeadLetteredMessagesTo": "",  
                                   "ForwardTo": "",  
                                   "LockDuration": "PT1M",  
                                   "MaxDeliveryCount": 10,  
                                   "RequiresDuplicateDetection": false,  
                                   "RequiresSession": false  
                                 }  
                               },{  
                                 "Name": "control-plane-orchestrations",  
                                 "Properties": {  
                                   "DeadLetteringOnMessageExpiration": false,  
                                   "DefaultMessageTimeToLive": "PT1H",  
                                   "DuplicateDetectionHistoryTimeWindow": "PT20S",  
                                   "ForwardDeadLetteredMessagesTo": "",  
                                   "ForwardTo": "",  
                                   "LockDuration": "PT1M",  
                                   "MaxDeliveryCount": 10,  
                                   "RequiresDuplicateDetection": false,  
                                   "RequiresSession": false  
                                 }  
                               }         ]       }     ],  
                         "Logging": {  
                           "Type": "File"  
                         }  
          }    }
          ```
- Run the below command to start the local service bus container
    - `docker compose -f <docker-compose-file-loc> up -d`
- Now use the below connection string to connect to the service bus- "Endpoint=sb://localhost;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=SAS_KEY_VALUE;UseDevelopmentEmulator=true;"
- Emulator repo
    - https://github.com/SMART2016/azure-service-bus-simulator
- For running the sender and reciever set below env variable
    - `export SERVICEBUS_QUEUE_NAME="control-plane-notifications"`
    - `export SERVICEBUS_CONNECTION_STRING="Endpoint=sb://localhost:5672;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=SAS_KEY_VALUE;UseDevelopmentEmulator=true;"`