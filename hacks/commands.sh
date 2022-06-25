mariaguicactl kubernetes node create --name master-0 --private-ip 193.168.50.10  --type master  --lb-ip 193.168.50.5 --primary
mariaguicactl kubernetes node create --name master-2 --private-ip 193.168.50.12  --type master  --lb-ip 193.168.50.5
mariaguicactl kubernetes node create --name worker-1 --private-ip 193.168.50.20  --type worker 
mariaguicactl kubernetes node create --name worker-2 --private-ip 193.168.50.21  --type worker 
mariaguicactl kubernetes node create -f node-worker.yml
mariaguicactl kubernetes node create -f node-master.yml

mariaguicactl load-balancer haProxy create -f ha.yml
mariaguicactl active-directory windowsserver16 create -f active-directory.yml
mariaguicactl database postgres create -f postgres.yml
mariaguicactl stream kafka create -f postgres.yml
mariaguicactl api-gateway kong create -f node-master.yml
mariaguicactl machine ubuntu create -f node-master.yml

mariaguicactl api-gateway kong create -f node-master.yml
mariaguicactl api-gateway kong konga
mariaguicactl api-gateway kong service 
mariaguicactl api-gateway kong plugin 


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