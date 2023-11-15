First off, let's start with a basic concept.
What is a network?
Simply put, it's a **group of devices that are connected to each other.**

The internet is nothing more than a network of networks.

All the devices in my home are connected to my local network, and that network is connected to all the other networks in the world -- that grouping is called the internet.

## ISP

An ISP (or Internet Service Provider) is a company that provides access to the mainstream internet. They bridge the gap between your local network and all the other networks, most often through a device called a _router,_ which is nothing more than a computer that connects to all computers in your local network, and forwards all of their data to the ISP by connecting to their communication network, which is *then* connected to all of the other ISPs in a bigger network.

The ISP, anytime they want to send over data to your devices, sends the information to your router. The router then sends it over to the device that requested the information.

Each device has a specific address to identify it --- it's called the IP address, or _Internet Protocol address._ That is how the router is able to know where it should send the data. Additionally, each application in the device uses a _port number,_ so that we may also specify which application will receive the data. 

An example IP address would look like 192.168.0.1. An example IP along with the port number would look something like 192.168.0.1:8080.

Websites are nothing more than public IP addresses, hosted by computers (called web servers), that can be accessed by other computers freely. For example, Google's address could be 721.111.42.11. Upon entering that into a web browser, you would be redirected to Google.
However, IP numbers are often considered difficult to remember, and they may change with time. That's where the Domain Name System ([DNS](Software%20Engineering/Internet/DNS.md)) comes into play.

## [DNS](Software%20Engineering/Internet/DNS.md)

A [domain name](Software%20Engineering/Internet/domain%20name.md) is nothing more than a human-readable address which, when entered into a web browser, makes a request to a specific kind of server from your ISP that finds the equivalent IP address associated with that [domain name](Software%20Engineering/Internet/domain%20name.md). This is known as the Domain Name System, or [DNS](Software%20Engineering/Internet/DNS.md).

Thus, using [DNS](Software%20Engineering/Internet/DNS.md), people do not have to remember the IP addresses of websites they'd like to visit -- that is the [DNS](Software%20Engineering/Internet/DNS.md) server's job. 

If the [DNS](Software%20Engineering/Internet/DNS.md) server is having issues responding to your computer, technically you could insert the IP address of the website you're looking for, and you would be able to access it in that way.
(In reality, it's not that simple. But in theory, sure!)

## Open Systems Interconnection 

Also known as the Layers model. 
This defines the different layers in which networking functions operate.

![[osi-model.png]]

##### Relevant material:
[Cloudfare OSI documentation](https://www.cloudflare.com/learning/ddos/glossary/open-systems-interconnection-model-osi/)