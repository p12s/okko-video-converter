#  --ask-become-pass - если sudo запрашивает пароль. но лучше сделать беспарольное переключение на sudo

site:
	ansible-playbook -i hosts.yml site.yml -vvv

upgrade:
	ansible-playbook -i hosts.yml upgrade.yml -v

renew-certificates:
	ansible-playbook -i hosts.yml renew-certificates.yml -v

authorize:
	ansible-playbook -i hosts.yml authorize.yml -v

generate-deploy-key:
	ssh-keygen -q -t rsa -N '' -f files/deploy_rsa -v

authorize-deploy:
	ansible-playbook -i hosts.yml authorize-deploy.yml -v

docker-login:
	ansible-playbook -i hosts.yml docker-login.yml -v
