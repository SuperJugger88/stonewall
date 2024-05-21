start:
	kubectl create ns demo \
	&& helm install demo charts/

stop:
	helm uninstall demo \
	&& kubectl delete ns/demo

restart:
	helm upgrade demo charts/