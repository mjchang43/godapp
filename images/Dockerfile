FROM registry.redhat.io/rhscl/nginx-112-rhel7
LABEL author="NARUTO"

USER root
WORKDIR /opt/gitlab

RUN yum install -y curl policycoreutils-python postfix fipscheck-lib tcp_wrappers-libs

#RUN rpm -ivh openssh-7.4p1-16.el7.x86_64.rpm \
#    openssh-server-7.4p1-16.el7.x86_64.rpm \
#    gitlab-ee-11.10.4-ee.0.el7.x86_64.rpm

#RUN systemctl enable sshd.service && systemctl start sshd.service
#RUN firewall-cmd --permanent --add-service=http && systemctl reload firewalld
#RUN systemctl enable postfix && systemctl start postfix

CMD ["/bin/bash"]