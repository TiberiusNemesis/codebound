# Domain Name System

A [domain name](Software%20Engineering/Internet/domain%20name.md) system (or DNS) is responsible for converting a human-readable address (like google.com) into an IP address (such as 192.0.2.172), allowing web browsers to make requests and access the resources of a network location without the need to know (and type in) the exact IP address associated with it.

# How does it work?

A group of servers is involved in the [domain name](Software%20Engineering/Internet/domain%20name.md) resolution process. 

First off, the *DNS recursor* is the first step. 

A DNS recursor will receive queries from a [web browser](Software%20Engineering/Internet/web%20browser.md), and look up if it has the information for that domain in its cache. If not, it forwards the request to a *DNS root nameserver*, which will respond to the resolver with the address of a top level domain (or TLD) server -- each TLD server stores the information for a specific top-level domain (such as *.net*, *.com*, *.edu*), and it responds to a recursor's request with the address for the *domain nameserver* (e.g. google.com's nameserver). 

The resolver then makes another request to this nameserver, and finally, it will get the IP address of the domain in the response.

The [domain name](Software%20Engineering/Internet/domain%20name.md) information may be stored/cached in several different parts of this process.
For the DNS client, it may be stored in their browser or operating system.   
In the server-side, the recursor may keep IP addresses for common domains in their cache in order to shorten the process for any requesting clients. 

### Sources

[What is DNS?](https://www.cloudflare.com/en-gb/learning/dns/what-is-dns/)
[How DNS works (comic)](https://howdns.works/)
[Understanding Domain names](https://developer.mozilla.org/en-US/docs/Glossary/DNS/)