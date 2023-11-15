# Web Browsers

Simply put, they are the most widely used software in the world.   
If we're not counting the operating system that the browsers are running on, of course.

Their main purpose is to allow users to access content within the web, provide simple, straightforward navigation between content, and visually render these the user.

# How do they work?

![A web browser's standard structure.](/Assets/web-browser-structure.png)

These are the main parts that make up a web browser. 

1. **The user interface**: this includes the address bar, the back/forward button, the bookmarks menu, etc. So, pretty much every part of the browser display, except the window where you see the requested page.
2. **The browser engine**: marshals actions between the UI and the rendering engine.
3. **The rendering engine**: displays requested content. If the requested content is HTML, the rendering engine parses HTML and CSS, and displays the parsed content on the screen. Other file types, such as JSON, MP4, or PDFs may also be displayed.
4. **Networking**: network calls, such as HTTP requests, using different implementations for different platform behind a platform-independent interface.
5. **UI backend**: draws basic widgets, such as combo boxes and windows. This backend exposes a generic interface that is not platform specific. Underneath, it uses operating system user interface methods.
6. **JavaScript interpreter**. Parses and executes JavaScript code.
7. **Data storage**. A persistence layer that stores data as needed. The browser may need to locally save or cache all sorts of data -- one such example are *cookies*. Browsers also support storage mechanisms such as *localStorage, IndexedDB, WebSQL* and *FileSystem.*

### Sources

[How Browsers Work](https://www.html5rocks.com/en/tutorials/internals/howbrowserswork/)  
[Role of Rendering Engine in Browsers](https://www.browserstack.com/guide/browser-rendering-engine)  
[Populating the Page: How Browsers Work](https://developer.mozilla.org/en-US/docs/Web/Performance/How_browsers_work) 