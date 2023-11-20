# News_fetcher

A simple http server designed to allow users to make a request to a url (to fetch content from news websites, for example). This code is purposely vulnerable to demonstrate the effect of server side request forgery attacks.

The code is susceptible to Cross-Site Scripting (XSS), a critical web application security flaw. The lack of sanitation enables attackers to insert harmful scripts, such as JavaScript, within the URL field. Upon submission, these scripts execute within the context of the page, posing severe risks. Exploiting this vulnerability, attackers can compromise user sessions, steal sensitive data, or manipulate the application's behavior, potentially impacting the security and integrity of the entire system. Additionally, this vulnerability opens the door to server-side request forgery, potentially allowing unauthorized access to internal systems or resources. For example, if a user sends a request to http://localhost:8080/?url=http://localhost:8080/secret, the server will forward a request to the /secret endpoint on localhost:8080. This could enable an attacker to gain access to sensitive information or execute unauthorized actions within the internal network. (VulnPlanet)

To prevent such security breaches, it's crucial to enforce rigorous input validation, properly sanitize user input, and implement output encoding techniques to prevent malicious script injection attempts. The News-fetcher_safe.go provides code that implements these factors, making it more secure.

Reference: https://github.com/yevh/VulnPlanet/blob/main/web2/type/ssrf.md
