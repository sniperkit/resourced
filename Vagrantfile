# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  # All Vagrant configuration is done here. The most common configuration
  # options are documented and commented below. For a complete reference,
  # please see the online documentation at vagrantup.com.

  config.vm.define "ubuntu" do |ubuntu|
    ubuntu.vm.box = "phusion/ubuntu-14.04-amd64"
    ubuntu.vm.provision :shell, path: "tests/data/provisioning/vagrant-ubuntu.sh"
    ubuntu.vm.network :private_network, type: :static, ip: "192.168.50.240"
  end

  config.vm.define "centos" do |centos|
    # centos.gui = true
    centos.vm.box = "metcalfc/centos70-docker"
    centos.vm.provision :shell, path: "tests/data/provisioning/vagrant-centos.sh"
    centos.vm.network :private_network, type: :static, ip: "192.168.50.241"
  end

  config.ssh.insert_key = 'true'

  config.ssh.forward_agent = true

  config.vm.synced_folder ENV['GOPATH'], "/go"
end
