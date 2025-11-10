# workflow-service

1. Start temporal server with docker-compose for local development.
```shell
git clone https://github.com/temporalio/docker-compose.git 
cd docker-compose
docker-compose up
```

2. Check temporal console at http://localhost:8080

3. Start http local http server
```shell
go run cmd/main.go
```

4. Try to run example workflow, and record the workflow_id from the response.
```shell
curl http://127.0.0.1:8765/v1/workflows/run -d'{"workflow_type":"test", "task_queue":"default-task-queue"}'
{"workflow_id":"8e30001d-cb6c-4961-8d1e-0d452450f3c2","run_id":"019a6d8b-0bee-7042-a98c-9538aa424dd5","started_at":"2025-11-10T19:33:40.473924+08:00"}
```

5. Check workflow status (running) in console at http://localhost:8080/namespaces/default/workflows/{workflow_id}
```
http://localhost:8080/namespaces/default/workflows/8e30001d-cb6c-4961-8d1e-0d452450f3c2
```

6. Send signal to the workflow http://127.0.0.1:8765/v1/workflows/{workflow_id}/signals
```shell
curl http://127.0.0.1:8765/v1/workflows/8e30001d-cb6c-4961-8d1e-0d452450f3c2/signals -d'{"signal_name":"my-signal"}'
```

7. Check workflow status (completed) in console at http://localhost:8080/namespaces/default/workflows/{workflow_id}
```
http://localhost:8080/namespaces/default/workflows/8e30001d-cb6c-4961-8d1e-0d452450f3c2
```