FROM jenkins/jenkins:lts

# Installer des plugins Jenkins
COPY plugins.txt /usr/share/jenkins/ref/plugins.txt
RUN jenkins-plugin-cli --plugins < plugins.txt

# Configurer Jenkins à démarrage
COPY custom-config.xml /usr/share/jenkins/ref/jobs/
