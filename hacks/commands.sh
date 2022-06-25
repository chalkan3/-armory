armory provider vagrant kubernetes node create --name master-0 --private-ip 193.168.50.10  --type master  --lb-ip 193.168.50.5 --primary
armory provider vagrant kubernetes node create --name master-2 --private-ip 193.168.50.12  --type master  --lb-ip 193.168.50.5
armory provider vagrant kubernetes node create --name worker-1 --private-ip 193.168.50.20  --type worker 
armory provider vagrant kubernetes node create --name worker-2 --private-ip 193.168.50.21  --type worker 
armory provider vagrant kubernetes node create -f node-worker.yml
armory provider vagrant kubernetes node create -f node-master.yml

armory pv vg k8s node create -n master-0 --private-ip 193.168.50.10  -t master  --lb-ip 193.168.50.5 --primary
armory pv vg k8s node create -n master-2 --private-ip 193.168.50.12  -t master  --lb-ip 193.168.50.5
armory pv vg k8s node create -n worker-1 --private-ip 193.168.50.20  -t worker 
armory pv vg k8s node create -n worker-2 --private-ip 193.168.50.21  -t worker 
armory pv vg k8s node create -f node-worker.yml
armory pv vg k8s node create -f node-master.yml



armory provider vagrant load-balancer haProxy create -f ha.yml
armory provider vagrant active-directory windowsserver16 create -f active-directory.yml
armory provider vagrant database postgres create -f postgres.yml
armory provider vagrant stream kafka create -f postgres.yml
armory provider vagrant api-gateway kong create -f node-master.yml
armory provider vagrant machine ubuntu create -f node-master.yml

armory provider vagrant api-gateway kong create -f node-master.yml
armory provider vagrant api-gateway kong konga
armory provider vagrant api-gateway kong service 
armory provider vagrant api-gateway kong plugin 


ssh vagrant@localhost -p 2229 -N -L  8001:127.0.1:8001


kubernetes node create --name worker-55 --private-ip 193.168.50.11  --type worker
mariaguicactl kubernetes node port-forward --name node-worker-10 --port 2200 --addr 4001:193.168.50.5:6443
mariaguicactl kubernetes node create --name worker-34 --private-ip 193.168.50.34  --type worker 

mariaguicactl kubernetes node delete --name node-worker-12


curl -i -X POST \
--url http://localhost:8001/services/testApi/routes \
--data 'hosts[]=localhost' \
--data 'paths[]=/api/v1/test1' \
--data 'strip_path=false' \
--data 'methods[]=GET'


curl -i -X POST \
--url http://localhost:8001/services/ \
--data 'name=testApi' \
--data 'url=http://localhost:5055'