# Web Hosting

To _host_ something on the web means to publish your website files onto the [internet](Software%20Engineering/Internet/internet.md), so that anyone who has access to the [internet](Software%20Engineering/Internet/internet.md) can also access your website.

(This is an oversimplification, of course. But hey, basic concepts, basic steps.)

### How does it work?

A web server stores all the data required to run your website, and it allows for users (or _clients_) to make requests to it -- the request, in this case, being all the files required for the user to render your website on their computer. 

By entering the correct [domain name](Software%20Engineering/Internet/domain%20name.md) (or IP address) for your server, the user requests the server to send these files to them, and if the server responds as expected, voilá: someone just accessed your website.

These files are usually HTML, along with CSS styling and perhaps some JavaScript running alongside it. These are used by the client's [web browser](Software%20Engineering/Internet/web%20browser.md) to render the page for them.

For the website to be publicly accessible on the [internet](Software%20Engineering/Internet/internet.md), a process called _domain name registration,_ is required where the owner of the website (in this case, you) must pay a fee to a company that manages the registration of domain names, and they are given a [domain name](Software%20Engineering/Internet/domain%20name.md) that is unique to them.

#### Different kinds of web hosting

A website can be hosted in different ways. 
- a dedicated server, where you have a single web server hosting your website
- shared hosting, where a single web server hosts several websites, possibly including yours
- VPS hosting, where a single web server hosts several websites, possibly including yours, but unlike shared hosting, a virtualized "piece" of the server hosts your website, which means you are guaranteed a certain amount of resources (such as RAM and storage) from the server. 

### Sources

[What Is Web Hosting? Explained](https://www.youtube.com/watch?v=htbY9-yggB0)  
[Different Types of Web Hosting Explained](https://www.youtube.com/watch?v=AXVZYzw8geg)   
[Where to Host a Fullstack Project on a Budget](https://www.youtube.com/watch?v=Kx_1NYYJS7Q)  
[What is a web server?](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/Web_mechanics/What_is_a_web_server)